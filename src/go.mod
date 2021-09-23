module sgr

go 1.16

require (
	github.com/jinzhu/copier v0.3.2
	github.com/yoyofx/yoyogo v1.7.9
	github.com/yoyofxteam/dependencyinjection v1.0.1
	gorm.io/gorm v1.21.11
	github.com/stretchr/testify v1.7.0
	gorm.io/driver/mysql v1.1.1
)

replace gorm.io/gorm v1.21.11 => github.com/go-gorm/gorm v1.21.11
