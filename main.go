package main

import (
	"fmt"

	"github.com/mikaelm1/pirate/cmd"
	"github.com/spf13/viper"
)

var (
	// VERSION is set during build
	VERSION = "1.0.0"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file")
	}

	cmd.Execute(VERSION)
}
