package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/stevenlee87/go-gin-example/pkg/file"
	"github.com/stevenlee87/go-gin-example/pkg/logging"
	"github.com/stevenlee87/go-gin-example/pkg/setting"
	"github.com/stevenlee87/go-gin-example/pkg/util"
)

func GetImageFullUrl(name string) string {
	// http://127.0.0.1:8000 + / + upload/images/ + name
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

// 获取文件后缀，获取文件名，然后把文件名进行加密，在加上后缀组合成新的文件名
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath // ImageSavePath = upload/images/
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath() // RuntimeRootPath = runtime/ + upload/images
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) { // ImageAllowExts = .jpg,.jpeg,.png
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize // ImageMaxSize = 5
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}

/*
在这里我们实现了 7 个方法，如下：

GetImageFullUrl：获取图片完整访问URL
GetImageName：获取图片名称
GetImagePath：获取图片路径
GetImageFullPath：获取图片完整路径
CheckImageExt：检查图片后缀
CheckImageSize：检查图片大小
CheckImage：检查图片
这里基本是对底层代码的二次封装，为了更灵活的处理一些图片特有的逻辑，并且方便修改，不直接对外暴露下层
*/
