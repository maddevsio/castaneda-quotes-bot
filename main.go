package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal configuration error: %s \n", err)
		os.Exit(1)
	}
}