module github.com/stevenlee87/go-gin-example

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.3.1-0.20180527032555-9e463b461434
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/boombuler/barcode v1.0.1-0.20180315051053-3c06908149f7
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.0
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/spec v0.19.4 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gomodule/redigo v2.0.1-0.20180401191855-9352ab68be13+incompatible
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.3
	github.com/tealeg/xlsx v1.0.4-0.20180419195153-f36fa3be8893
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.2 // indirect
	golang.org/x/image v0.0.0-20191214001246-9130b4cfad52 // indirect
	golang.org/x/net v0.0.0-20191204025024-5ee1b9f4859a // indirect
	golang.org/x/sys v0.0.0-20191204072324-ce4227a45e2e // indirect
	golang.org/x/tools v0.0.0-20191204011308-9611592c72f6
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.2 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
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
