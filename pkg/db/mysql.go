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

// Initialize mysql connection
func InitMySQLCon() (err error) {
	mysql := config.GetConfig().MySQL
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysql.User, mysql.Passwd, mysql.Host, mysql.Port, mysql.Database)
	fmt.Println(connStr)
	DB, err = gorm.Open("mysql", connStr)

	if err != nil {
		return err
	}

	return DB.DB().Ping()
}
