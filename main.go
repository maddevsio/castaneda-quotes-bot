package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	configFile string
	configType string
}

func (c *Config) Init() {
	viper.SetConfigName(c.configFile)
	viper.SetConfigType(c.configType)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal configuration error: %s \n", err)
		os.Exit(1)
	}
}

func (c *Config) Get(key string) interface{} {
	return viper.Get(key)
}

func main() {
	config := Config{"config", "yml"}
	config.Init()
	token := config.Get("telegram-token")
	log.Printf("%v\n", token)
}