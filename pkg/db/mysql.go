package db

import (
	"fmt"
	"golang-gin-api/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQLCon() (err error) {
	// 可以在api包里设置成init函数
	config.ParserConfig()
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbConfig.User, config.DbConfig.Passwd, config.DbConfig.Host, config.DbConfig.Port, config.DbConfig.Database)
	fmt.Println(connStr)
	DB, err = gorm.Open("mysql", connStr)

	if err != nil {
		return err
	}

	return DB.DB().Ping()
}
