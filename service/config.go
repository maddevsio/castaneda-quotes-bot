package service

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	configFile string
	configType string
}

func (c *Config) Get(key string) interface{} {
	return viper.Get(key)
}

func NewConfig(configFile string, configType string) Config {
	config := Config{configFile, configType}
	viper.SetConfigName(config.configFile)
	viper.SetConfigType(config.configType)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal configuration error: %s \n", err)
	}
	return config
}

