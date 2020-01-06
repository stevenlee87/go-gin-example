package qrcode

import (
	"image/jpeg"

	"fmt"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"github.com/stevenlee87/go-gin-example/pkg/file"
	"github.com/stevenlee87/go-gin-example/pkg/setting"
	"github.com/stevenlee87/go-gin-example/pkg/util"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

// NewQrCode initialize instance
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

// GetQrCodePath get save path
func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath // qrcode/
}

// GetQrCodeFullPath get full save path
func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath // runtime/ + qrcode/
}

// GetQrCodeFullUrl get the full access path
// http://127.0.0.1:8000 + / + qrcode/ + name
func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}

// GetQrCodeFileName get qr file name
func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

// GetQrCodeExt get qr file ext
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// Encode generate QR code
/*
这里主要聚焦 func (q *QrCode) Encode 方法，做了如下事情：
获取二维码生成路径
创建二维码
缩放二维码到指定大小
新建存放二维码图片的文件
将图像（二维码）以 JPEG 4：2：0 基线格式写入文件
另外在 jpeg.Encode(f, code, nil) 中，第三个参数可设置其图像质量，默认值为 75
*/
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name // /runtime/qrcode/xxx.jpg
	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode) // QRCODE_URL M Auto
		fmt.Printf("code is %v#", code)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
