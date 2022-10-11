package tests

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	UserName string `config:"username"`
	Password string `config:"password"`
	Url      string `config:"url"`
}

func InitDb() *gorm.DB {
	dsnPath := fmt.Sprintf("%s:%s@%s", "root", "P@ssW0rd", "tcp(47.100.213.41)/sgr_pass?charset=utf8mb4&loc=Local&parseTime=True")
	db1, _ := gorm.Open(mysql.Open(dsnPath), &gorm.Config{})
	return db1
}
