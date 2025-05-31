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
		GoBackendURL:  viper.GetString("GO_BACKEND_URL"),
		NodeBackendURL: viper.GetString("NODE_BACKEND_URL"),
	}
	if config.Port == "" {
		config.Port = "8000"
	}
	if config.GoBackendURL == "" {
		config.GoBackendURL = "http://localhost:3000"
	}
	if config.NodeBackendURL == "" {
		config.NodeBackendURL = "http://localhost:5500"
	}

	return config
}
