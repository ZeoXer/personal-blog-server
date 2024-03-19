package main

import (
	"fmt"

	"go-server/global"

	"github.com/spf13/viper"
)

func InitViper() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Start using Viper here
	v.Unmarshal(&global.CONFIG)
}
