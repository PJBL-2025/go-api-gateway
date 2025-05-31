package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Port          string
	GoBackendURL  string
	NodeBackendURL string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	config := Config{
		Port:          viper.GetString("PORT"),
		GoBackendURL:  viper.GetString("GO_HOST"),
		NodeBackendURL: viper.GetString("NODE_HOST"),
	}
	if config.Port == "" {
		config.Port = "8000"
	}
	if config.GoBackendURL == "" {
		config.GoBackendURL = "http://app:3000"
	}
	if config.NodeBackendURL == "" {
		config.NodeBackendURL = "http://node-api:5500"
	}

	return config
}
