package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
	Paystack PaystackConfig
	Flutterwave FlutterwaveConfig
	Stripe StripeConfig
}

type PaystackConfig struct {
	BaseURL string
	SecretKey string
}

type StripeConfig struct {
	BaseURL string
	SecretKey string
}

type FlutterwaveConfig struct {
	BaseURL string
	SecretKey string
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		Paystack: PaystackConfig{
            BaseURL:   os.Getenv("PAYSTACK_BASE_URL"),
            SecretKey: os.Getenv("PAYSTACK_SECRET"),
        },
		Stripe: StripeConfig{
			BaseURL:   os.Getenv("STRIPE_BASE_URL"),
			SecretKey: os.Getenv("STRIPE_SECRET"),
		},
		Flutterwave: FlutterwaveConfig{
			BaseURL:   os.Getenv("FLUTTERWAVE_BASE_URL"),
			SecretKey: os.Getenv("FLUTTERWAVE_SECRET"),
		},
	}, nil
}