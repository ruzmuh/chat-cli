package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"sort"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/viper"
)

func extractCode(message string) (result string, err error) {
	re := regexp.MustCompile("(?s)```.*?\n(.*?)```") // regex to first code snippet

	matches := re.FindStringSubmatch(message)
	if len(matches) < 1 {
		err = fmt.Errorf("code snippet not found")
		return
	}
	result = matches[1]
	return
}

func chatCompletion(prompt string, codeonly bool) (result string, err error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: viper.GetString(modelFlagName),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return
	}
	result = resp.Choices[0].Message.Content

	if codeonly {
		result, err = extractCode(result)
		if err != nil {
			return
		}
	}
	return
}

func initAndRunProgressbar() *progressbar.ProgressBar {
	bar := progressbar.NewOptions(-1,
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("requestring"),
		progressbar.OptionSpinnerType(11),
		progressbar.OptionSetWriter(os.Stderr),
	)
	go func() {
		for {
			if bar.IsFinished() {
				return
			}
			bar.Add(1)
			time.Sleep(100 * time.Millisecond)
		}

	}()
	return bar
}

func showListOfModels() (result openai.ModelsList, err error) {
	result, err = client.ListModels(context.Background())
	if err != nil {
		return
	}
	sort.Slice(result.Models, func(i, j int) bool {
		return result.Models[i].ID < result.Models[j].ID
	})
	return
}
