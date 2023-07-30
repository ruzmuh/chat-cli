package main

import (
	"fmt"
	"log"
	"os"

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

	client = openai.NewClient(viper.GetString(tokenFlagName))

	if versionFlag {
		fmt.Print(Version)
		os.Exit(0)
	}
	if initFlag {
		initConfig()
	}

	bar := initAndRunProgressbar()
	defer bar.Finish()
	if listModelsFlag {
		modelsList, err := showListOfModels()
		bar.Finish()
		if err != nil {
			log.Fatalf("ModelList error: %v\n", err)
		}
		for _, v := range modelsList.Models {
			fmt.Println(v.ID)
		}
		os.Exit(0)
	}

	answer, err := chatCompletion(prompt, viper.GetBool(codeOnlyFlagName))
	bar.Finish()
	if err != nil {
		log.Fatalf("ChatCompletion error: %v\n", err)
	}
	fmt.Println(answer)
}
