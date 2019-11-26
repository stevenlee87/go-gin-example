module github.com/stevenlee87/go-gin-example

go 1.13

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20191007083155-a98c25f2172a // indirect
	github.com/astaxie/beego v1.12.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.51.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20191113165036-4c7a9d0fe056 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace (
	github.com/stevenlee87/go-gin-example/conf => ./go-gin-example/conf
	github.com/stevenlee87/go-gin-example/middleware => ./go-gin-example/middleware
	github.com/stevenlee87/go-gin-example/models => ./go-gin-example/models
	github.com/stevenlee87/go-gin-example/pkg/e => ./go-gin-example/pkg/e
	github.com/stevenlee87/go-gin-example/pkg/setting => ./go-gin-example/pkg/setting
	github.com/stevenlee87/go-gin-example/pkg/util => ./go-gin-example/pkg/util
	github.com/stevenlee87/go-gin-example/routers => ./go-gin-example/routers
	github.com/stevenlee87/go-gin-example/runtime => ./go-gin-example/runtime
)
