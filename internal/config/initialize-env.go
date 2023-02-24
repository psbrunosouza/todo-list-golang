package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeEnv() {
	dotEnvError := godotenv.Load()

	if dotEnvError != nil {
		log.Fatal("Error loading .env file")
	}
}
