package service

import (
	"github.com/limitcool/blog/common/upload"
	"github.com/pkg/errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	// 获取文件所需的基本信息
	fileName := upload.GetFileName(fileHeader.Filename)
	if upload.IsMarkdownExist(fileName) {
		return nil, errors.New("文件已存在")
	}
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("文件格式不支持！")
	}
	if !upload.CheckMaxsize(fileType, file) {
		return nil, errors.New("文件超出上限！")
	}
	uploadSavePath := upload.GetSavePath(fileType)
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("创建文件目录失败！")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("文件权限不足！")
	}
	dst := uploadSavePath + "/" + fileName
	if err := upload.SavaFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := upload.GetUrlSavePath(fileType) + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
