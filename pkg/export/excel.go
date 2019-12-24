package export

import (
	"github.com/stevenlee87/go-gin-example/pkg/setting"
)

const EXT = ".xlsx"

// GetExcelFullUrl get the full access path of the Excel file
// http://127.0.0.1:8000 + / + export/ + name
func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath // export/
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath() // runtime/ + export/
}
