package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	config = new(Config)
)

type Config struct {
	MySQL struct {
		Host     string `json:"host"`
		Port     int64  `json:"port"`
		User     string `json:"user"`
		Passwd   string `json:"passwd"`
		Database string `json:"database"`
	} `toml:"mysql"`

	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`
}

func init() {
	viperInit := viper.New()
	viperInit.AddConfigPath("./config")
	viperInit.SetConfigName("config")
	viperInit.SetConfigType("toml")

	if err := viperInit.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viperInit.Unmarshal(config); err != nil {
		fmt.Println(err)
	}
}

func GetConfig() Config {
	return *config
}

func ProjectName() string {
	return "golang-gin-api"
}

func ProjectLogFile() string {
	return fmt.Sprintf("./logs/%s-access.log", ProjectName())
}
