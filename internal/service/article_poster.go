package service

import (
	"github.com/goki/freetype"
	"github.com/limitcool/blog/internal/model"
	"github.com/limitcool/blog/pkg/file"
	"github.com/limitcool/blog/pkg/qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
)

type ArticlePoster struct {
	// 海报名称
	PosterName string
	// 文章信息
	*model.Articles
	// 二维码
	Qr *qrcode.QrCode
}

func NewArticlePoster(posterName string, article *model.Articles, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Articles:   article,
		Qr:         qr,
	}
}
func GetPosterFlag() string {
	return "poster"
}

// 检查是否有合并后的图片
func (a ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckNotExist(path+a.PosterName) == true {
		return false
	}
	return true
}
func (a ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}
type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

type Pt struct {
	X int
	Y int
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

func (a *ArticlePosterBg) Generate() (string, string, error) {
	fullpath := qrcode.GetQrCodePath()
	fileName, path, err := a.Qr.Encode(fullpath)
	if err != nil {
		return "", "", err
	}
	if !a.CheckMergedImage(path) {
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()
		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()
		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()
		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))
		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)
		jpeg.Encode(mergedF, jpg, nil)
	}
	return fileName, path, nil
}

type DrawText struct {
	JPG    draw.Image
	Merged *os.File

	Title string
	X0    int
	Y0    int
	Size0 float64

	SubTitle string
	X1       int
	Y1       int
	Size1    float64
}

func (a ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	fontSource := "storage/fonts/msyhbd.ttc"
	// 读取字体: 微软雅黑
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}
	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return nil
	}
	fc := freetype.NewContext()
	// fc.SetDPI: 设置每英寸的分辨率
	fc.SetDPI(72)
	// fc.SetFont 设置用于绘制文本的字体
	fc.SetFont(trueTypeFont)
	// fc.SetFontSize 设置字体大小
	fc.SetFontSize(d.Size0)
	// fc.SetClip 设置剪裁矩形以进行绘制
	fc.SetClip(d.JPG.Bounds())
	// fc.SetDst 设置目标图像
	fc.SetDst(d.JPG)
	// fc.SetSrc 设置绘制操作的源图像,通常为image.Uniform
	fc.SetSrc(image.Black)

	pt := freetype.Pt(d.X0, d.Y0)
	// fc.DrawString 根据Pt的坐标值绘制给定的文本内容
	_, err = fc.DrawString(d.Title, pt)
	if err != nil {
		return err
	}
	fc.SetFontSize(d.Size1)
	_, err = fc.DrawString(d.SubTitle, freetype.Pt(d.X1, d.Y1))
	if err != nil {
		return err
	}
	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		return err
	}
	return nil

}
