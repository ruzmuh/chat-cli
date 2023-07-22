package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/sashabaranov/go-openai"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	client  *openai.Client
	Version = "dev"
)

func main() {
	NewConfig()
	prompt := flag.Arg(0)

	client = openai.NewClient(viper.GetString("token"))

	if versionFlag {
		fmt.Print(Version)
		os.Exit(0)
	}

	if listModelsFlag {
		showListOfModels()
	}

	if initFlag {
		initConfig()
	}

	chatCompletion(prompt)

}

func chatCompletion(prompt string) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: viper.GetString("model"),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		log.Fatalf("ChatCompletion error: %v\n", err)
	}

	answer := resp.Choices[0].Message.Content

	if viper.GetBool("codeonly") {
		answer, err = extractCode(answer)
		if err != nil {
			log.Fatalf("extracting code snippet error: %v\n", err)
		}
	}
	fmt.Println(answer)
	os.Exit(0)
}

func showListOfModels() {
	modelsList, err := client.ListModels(context.Background())
	if err != nil {
		log.Fatalf("ListModels error: %v\n", err)
	}

	sort.Slice(modelsList.Models, func(i, j int) bool {
		return modelsList.Models[i].ID < modelsList.Models[j].ID
	})

	for _, v := range modelsList.Models {
		fmt.Println(v.ID)
	}
	os.Exit(0)
}

func initConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("OpenAI token []: ")
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)
	viper.Set("token", token)

	fmt.Printf("model [%s]: ", viper.GetString("model"))
	model, _ := reader.ReadString('\n')
	model = strings.TrimSpace(model)
	if model != "" {
		viper.Set("model", model)
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
