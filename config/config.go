package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Optional, for loading .env files
)

var AppConfig *Config

type Config struct {
	Broker  string
	ApiKey  string
	AppMode string
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConfig() {
	// Optional: load .env file in development
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	// Read environment variables
	AppConfig = &Config{
		Broker:  getEnv("KAFKA_BROKER", "localhost:9092"),
		ApiKey:  getEnv("API_KEY", "blockhouse-key"),
		AppMode: getEnv("APP_MODE", "development"),
	}
}

// Helper function to read an environment variable or fallback to a default
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
