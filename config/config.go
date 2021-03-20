package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Database string `json:"database"`
}

var (
	DbConfig DBConfig
)

func ParserConfig() {
	config := viper.New()
	config.AddConfigPath("./config")
	config.SetConfigName("config")
	config.SetConfigType("toml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	DbConfig.Host = config.GetString("mysql.host")
	DbConfig.Port = config.GetInt64("mysql.port")
	DbConfig.User = config.GetString("mysql.user")
	DbConfig.Passwd = config.GetString("mysql.passwd")
	DbConfig.Database = config.GetString("mysql.database")
}

func ProjectName() string {
	return "golang-gin-api"
}

func ProjectLogFile() string {
	return fmt.Sprintf("./logs/%s-access.log", ProjectName())
}
