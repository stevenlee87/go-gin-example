package main

import (
	"fmt"
	"net/http"

	"github.com/stevenlee87/go-gin-example/pkg/setting"
	"github.com/stevenlee87/go-gin-example/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

/*
POST：http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1
GET：http://127.0.0.1:8000/api/v1/articles
GET：http://127.0.0.1:8000/api/v1/articles/1
PUT：http://127.0.0.1:8000/api/v1/articles/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0
DELETE：http://127.0.0.1:8000/api/v1/articles/1
*/
