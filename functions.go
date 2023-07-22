package main

import (
	"fmt"
	"regexp"
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
