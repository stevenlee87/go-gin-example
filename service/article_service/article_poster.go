package article_service

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/stevenlee87/go-gin-example/pkg/logging"

	"github.com/golang/freetype"

	"github.com/stevenlee87/go-gin-example/pkg/file"
	"github.com/stevenlee87/go-gin-example/pkg/qrcode"
	"github.com/stevenlee87/go-gin-example/pkg/setting"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePoster(posterName string, article *Article, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article:    article,
		Qr:         qr,
	}
}

func GetPosterFlag() string {
	return "poster"
}

func (a *ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckNotExist(path+a.PosterName) == true {
		return false
	}

	return true
}

func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
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

func (a *ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	fontSource := setting.AppSetting.RuntimeRootPath + setting.AppSetting.FontSavePath + fontName
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}

	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}

	fc := freetype.NewContext()
	fc.SetDPI(72)
	fc.SetFont(trueTypeFont)
	fc.SetFontSize(d.Size0)
	fc.SetClip(d.JPG.Bounds())
	fc.SetDst(d.JPG)
	fc.SetSrc(image.Black)

	pt := freetype.Pt(d.X0, d.Y0)
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

/*
func (a *ArticlePosterBg) Generate() 方法，做了如下事情：
获取二维码存储路径
生成二维码图像
检查合并后图像（指的是存放合并后的海报）是否存在
若不存在，则生成待合并的图像 mergedF
打开事先存放的背景图 bgF
打开生成的二维码图像 qrF
解码 bgF 和 qrF 返回 image.Image
创建一个新的 RGBA 图像
在 RGBA 图像上绘制 背景图（bgF）
在已绘制背景图的 RGBA 图像上，在指定 Point 上绘制二维码图像（qrF）
将绘制好的 RGBA 图像以 JPEG 4：2：0 基线格式写入合并后的图像文件（mergedF）
*/
//func (a *ArticlePosterBg) Generate() (string, string, error) {
//	fullPath := qrcode.GetQrCodeFullPath()
//	fileName, path, err := a.Qr.Encode(fullPath)
//	if err != nil {
//		return "", "", err
//	}
//
//	if !a.CheckMergedImage(path) {
//		mergedF, err := a.OpenMergedImage(path)
//		if err != nil {
//			return "", "", err
//		}
//		defer mergedF.Close()
//
//		bgF, err := file.MustOpen(a.Name, path)
//		if err != nil {
//			return "", "", err
//		}
//		defer bgF.Close()
//
//		qrF, err := file.MustOpen(fileName, path)
//		if err != nil {
//			return "", "", err
//		}
//		defer qrF.Close()
//
//		bgImage, err := jpeg.Decode(bgF)
//		if err != nil {
//			return "", "", err
//		}
//		qrImage, err := jpeg.Decode(qrF)
//		if err != nil {
//			return "", "", err
//		}
//
//		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))
//
//		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
//		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)
//
//		jpeg.Encode(mergedF, jpg, nil)
//	}
//
//	return fileName, path, nil
//}

/*
image.Rect     			 https://golang.org/pkg/image/#Rect
image.Pt       			 https://golang.org/pkg/image/#Pt
image.NewRGBA  			 https://golang.org/pkg/image/#NewRGBA
jpeg.Encode    			 https://golang.org/pkg/image/jpeg/#Encode
jpeg.Decode    			 https://golang.org/pkg/image/jpeg/#Decode
draw.Op        			 https://golang.org/pkg/image/draw/#Op
draw.Draw      			 https://golang.org/pkg/image/draw/#Draw
go-imagedraw-package     https://blog.golang.org/go-imagedraw-package
*/

func (a *ArticlePosterBg) Generate() (string, string, error) {
	fullPath := qrcode.GetQrCodeFullPath() // runtime/ + qrcode/
	fileName, path, err := a.Qr.Encode(fullPath)
	//fmt.Printf("fileName is:%s, path is %s", fileName, path)
	if err != nil {
		return "", "", err
	}

	if !a.CheckMergedImage(path) {
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()

		bgF, err := file.MustOpen(a.Name, path) // bg.jpg + path
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()

		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()

		bgImage, err := jpeg.Decode(bgF) // Decode reads a JPEG image from r and returns it as an image.Image.
		if err != nil {
			return "", "", err
		}
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}

		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1)) // 0 0 550 750

		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over) // 125 298

		err = a.DrawPoster(&DrawText{
			JPG:    jpg,
			Merged: mergedF,

			Title: "Golang Gin example",
			X0:    80,
			Y0:    160,
			Size0: 42,

			SubTitle: "---stevenlee",
			X1:       320,
			Y1:       220,
			Size1:    36,
		}, "msyhbd.ttc")

		if err != nil {
			logging.Warn(err)
			return "", "", err
		}
	}

	return fileName, path, nil
}
