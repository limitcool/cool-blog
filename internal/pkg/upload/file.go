package upload

import (
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

// GetFileName 获取文件名称，先是通过获取文件后缀并筛出原始文件名进行 MD5 加密，最后返回经过加密处理后的文件名。
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.Md5(fileName)
	return fileName + ext
}

// GetFileExt 获取文件后缀，主要是通过调用 path.Ext 方法进行循环查找”.“符号，最后通过切片索引返回对应的文化后缀名称。
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存地址，这里直接返回配置中的文件保存目录即可，也便于后续的调整。
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 检查保存目录是否存在，通过调用 os.Stat 方法获取文件的描述信息 FileInfo，并调用 os.IsNotExist 方法进行判断，其原理是利用 os.Stat 方法所返回的 error 值与系统中所定义的 oserror.ErrNotExist 进行判断，以此达到校验效果
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsExist(err)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// CheckMaxsize 检查文件大小是否超出最大大小限制。
func CheckMaxsize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return false
		}
	}
	return true
}

// CheckPermission 检查文件权限是否足够，与 CheckSavePath 方法原理一致，是利用 oserror.ErrPermission 进行判断。
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// CreateSavePath 创建在上传文件时所使用的保存目录，在方法内部调用的 os.MkdirAll 方法，该方法将会以传入的 os.FileMode 权限位去递归创建所需的所有目录结构，若涉及的目录均已存在，则不会进行任何操作，直接返回 nil。
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// SavaFile 保存所上传的文件，该方法主要是通过调用 os.Create 方法创建目标地址的文件，再通过 file.Open 方法打开源地址的文件，结合 io.Copy 方法实现两者之间的文件内容拷贝
func SavaFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}
