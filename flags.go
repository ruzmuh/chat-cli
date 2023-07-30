package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	codeOnlyFlagName   = "codeonly"
	initFlagName       = "init"
	listModelsFlagName = "list-models"
	modelFlagName      = "model"
	tokenFlagName      = "token"
	versionFlagName    = "version"
)

var (
	initFlag       bool
	listModelsFlag bool
	versionFlag    bool
)

func init() {
	flag.Usage = func() {
		fmt.Printf("USAGE:\n  %s [flags] \"<propmt>\"\nFLAGS:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&listModelsFlag, listModelsFlagName, false, "list available models and exit")
	flag.BoolVar(&initFlag, initFlagName, false, "init config")
	flag.BoolVar(&versionFlag, versionFlagName, false, "print version and exit")

	flag.String(modelFlagName, "", "selected model")
	flag.String(tokenFlagName, "", "your openAI token")
	flag.Bool(codeOnlyFlagName, false, "print first code snipet from the answer")
	flag.Parse()

	viper.BindPFlag(modelFlagName, flag.Lookup(modelFlagName))
	viper.BindPFlag(tokenFlagName, flag.Lookup(tokenFlagName))
	viper.BindPFlag(codeOnlyFlagName, flag.Lookup(codeOnlyFlagName))
}
