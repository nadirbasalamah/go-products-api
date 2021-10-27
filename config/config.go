package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config return value from .env file
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}
	return os.Getenv(key)
}
