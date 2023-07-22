package main

import (
	"log"

	"github.com/spf13/viper"
)

func NewConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/chat-cli/")
	viper.AddConfigPath("/etc/chat-cli/")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("chat")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}
	viper.SetDefault("model", "gpt-3.5-turbo")
}
