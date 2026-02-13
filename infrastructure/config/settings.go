package config

import (
	"os"
)

type Config struct {
	Port         string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecretKey string
}

func LoadConfig() *Config {
	LoadEnv()

	return &Config{
		Port:         getEnv("PORT", "8080"),
		DBHost:       getEnv("DATABASE_HOST", "localhost"),
		DBPort:       getEnv("DATABASE_PORT", "5432"),
		DBUser:       getEnv("DATABASE_USER", "postgres"),
		DBPassword:   getEnv("DATABASE_PASSWORD", ""),
		DBName:       getEnv("DATABASE_NAME", "book_khoone"),
		JWTSecretKey: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}
