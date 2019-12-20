package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/stevenlee87/go-gin-example/pkg/logging"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
