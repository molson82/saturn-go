package config

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Constants struct {
	LogLevel           string
	Port               string
	NewRelicAppName    string
	NewRelicLicense    string
	ElegantCMSToken    string
	ElegantCMSUrl      string
	TwitchClientId     string
	TwitchClientSecret string
	RedisURL           string
}

type Config struct {
	Constants
	NewRelicApp *newrelic.Application
	Redis       *redis.Client
}

var AppConfig Config

func New() *Config {
	initEnv()
	AppConfig = Config{}

	constants := Constants{
		os.Getenv("LOG_LEVEL"),
		os.Getenv("PORT"),
		os.Getenv("NEW_RELIC_APP_NAME"),
		os.Getenv("NEW_RELIC_LICENSE_KEY"),
		os.Getenv("ELEGANT_CMS_TOKEN"),
		os.Getenv("ELEGANT_CMS_URL"),
		os.Getenv("TWITCH_CLIENT_ID"),
		os.Getenv("TWITCH_CLIENT_SECRET"),
		os.Getenv("REDIS_URL"),
	}

	AppConfig.Constants = constants

	// Initialize New Relic
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(constants.NewRelicAppName),
		newrelic.ConfigLicense(constants.NewRelicLicense),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		log.Printf("Error loading New Relic Config: %v\n", err)
	}

	AppConfig.NewRelicApp = app

	// initialize Redis connection
	opt, err := redis.ParseURL(constants.RedisURL)
	if err != nil {
		log.Printf("Error connecting to Redis: %v\n", err)
	}
	redisClient := redis.NewClient(opt)

	AppConfig.Redis = redisClient

	log.Printf("Redis connected\n")

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
