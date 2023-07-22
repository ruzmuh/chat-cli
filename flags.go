package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	listModelsFlag bool
	initFlag       bool
	versionFlag    bool
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n  %s [flags] \"<propmt>\"\n  FLAGS:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&listModelsFlag, "list-models", false, "list available models and exit")
	flag.BoolVar(&initFlag, "init", false, "init config")
	flag.BoolVar(&versionFlag, "version", false, "print version and exit")

	flag.String("model", "", "selected model")
	flag.String("token", "", "your openAI token")
	flag.Bool("codeonly", false, "print first code snipet from the answer")
	flag.Parse()

	viper.BindPFlag("model", flag.Lookup("model"))
	viper.BindPFlag("token", flag.Lookup("token"))
	viper.BindPFlag("codeonly", flag.Lookup("codeonly"))
}
