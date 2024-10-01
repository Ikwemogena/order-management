package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}, nil
}