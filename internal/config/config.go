package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnection   string
	SendGridKey    string
	GoogleBooksKey string
	Port           string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		DBConnection:   os.Getenv("DATABASE_URL"),
		SendGridKey:    os.Getenv("SENDGRID_API_KEY"),
		GoogleBooksKey: os.Getenv("GOOGLE_BOOKS_API_KEY"),
		Port:           getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
