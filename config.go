package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func NewConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/chat-cli/")
	viper.AddConfigPath("/etc/chat-cli/")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("chatcli")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}
	viper.SetDefault(modelFlagName, "gpt-3.5-turbo")
}

func initConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("OpenAI token []: ")
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)
	viper.Set(tokenFlagName, token)

	fmt.Printf("model [%s]: ", viper.GetString(modelFlagName))
	model, _ := reader.ReadString('\n')
	model = strings.TrimSpace(model)
	if model != "" {
		viper.Set(modelFlagName, model)
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dirname = fmt.Sprintf("%s/.config/chat-cli/", dirname)

	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	configFile := fmt.Sprintf("%sconfig.yaml", dirname)
	err = viper.WriteConfigAs(configFile)
	if err != nil {
		log.Fatalf("Can't save config: %s", err.Error())
	}
	log.Printf("saved config %s", configFile)
	os.Exit(0)
}
