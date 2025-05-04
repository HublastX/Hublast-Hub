package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
	DB_URL string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Port:   getEnv("PORT", "8080"),
		DB_URL: getEnv("DB_URL", "localhost"),
		DbHost: getEnv("DB_HOST", "localhost"),
		DbPort: getEnv("DB_PORT", "5432"),
		DbUser: getEnv("DB_USER", "user"),
		DbPass: getEnv("DB_PASS", "password"),
		DbName: getEnv("DB_NAME", "project_management"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
