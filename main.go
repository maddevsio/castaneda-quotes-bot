package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
		os.Exit(1)
	}
}