package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Constants struct {
	LogLevel string
	Port     string
}

type Config struct {
	Constants
}

var AppConfig Config

func New() *Config {
	initEnv()
	AppConfig = Config{}

	constants := Constants{
		os.Getenv("LOG_LEVEL"),
		os.Getenv("PORT"),
	}

	AppConfig.Constants = constants

	return &AppConfig
}

func GetAppConfig() *Config {
	return &AppConfig
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %q\n", err)
	}
}
