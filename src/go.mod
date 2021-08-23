module sgr

go 1.16

require (
 	github.com/yoyofx/yoyogo v1.7.7
	github.com/fasthttp/websocket v1.4.3
	github.com/go-sql-driver/mysql v1.6.0
	github.com/yoyofxteam/dependencyinjection v1.0.1
	github.com/yoyofxteam/reflectx v0.2.3
	gorm.io/gorm v1.21.13
	github.com/jinzhu/copier v0.3.2
	github.com/yoyofxteam/nacos-viper-remote v0.4.0
)

replace gorm.io/gorm v1.21.13 => github.com/go-gorm/gorm v1.21.13
