package config

import (
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
	config.SetConfigType("ini")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	DbConfig.Host = config.GetString("db.host")
	DbConfig.Port = config.GetInt64("db.port")
	DbConfig.User = config.GetString("db.user")
	DbConfig.Passwd = config.GetString("db.passwd")
	DbConfig.Database = config.GetString("db.database")
}
