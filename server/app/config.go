package app

import (
	"github.com/spf13/viper"
)

var Config config

type config struct {
	DB struct {
		Name     string
		Host     string
		Port     int
		User     string
		Password string
	}
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
