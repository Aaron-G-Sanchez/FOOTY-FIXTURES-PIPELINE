package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIToken string
}

type Database struct {
	StagingURI string
}

var (
	AppConfig      Config
	DatabaseConfig Database
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	AppConfig.APIToken = os.Getenv("API_TOKEN")
	DatabaseConfig.StagingURI = os.Getenv("STAGING_DB_URI")
}
