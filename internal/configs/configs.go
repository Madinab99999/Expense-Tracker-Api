package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiHost string
	ApiPort string

	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string

	TokenSecret string
}

func LoadConfig() *Config {
	godotenv.Load()

	return &Config{
		ApiHost:     getEnv("API_HOST", "localhost"),
		ApiPort:     getEnv("API_PORT", "8080"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "postgres"),
		DBHost:      getEnv("DB_HOST", "127.0.0.1"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBName:      getEnv("DB_NAME", "expense_tracker"),
		TokenSecret: getEnv("TOKEN_SECRET", "ExpenseTrackerSecret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
