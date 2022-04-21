package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/pkg/file"
	"github.com/limitcool/blog/util"
	"image/jpeg"
	"log"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const EXT_JPG = ".jpg"

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Ext:    EXT_JPG,
		Level:  level,
		Mode:   mode,
	}
}

func GetQrCodePath() string {
	return global.AppSetting.QrCodeSavePath + "/"
}

func GetQrCodeFullUrl(name string) string {
	return global.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	// 检查文件是否存在,如果不存在,返回false
	if file.CheckNotExist(src) {
		return false
	}
	return true
}

// 生成二维码
func (q QrCode) Encode(path string) (string, string, error) {
	log.Println(path)
	// name 保存生成后的MD5文件名称+文件后缀
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	// src 生成后的文件路径为 path + name
	src := path + name
	// 判断生成后的路径是否不存在,如果不存在就开始生成二维码
	if file.CheckNotExist(src) == true {
		// 进行二维码加密
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}
		// 设置二维码输出后的尺寸
		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}
		// 创建文件
		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()
		// 将图片写入文件 第三个参数可设置其图像质量，默认值为 75
		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}

	}
	return name, path, nil
}
