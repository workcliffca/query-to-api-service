package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	AdminKey   string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "1433"),
		DBUser:     getEnv("DB_USER", "azuresql_admin"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "queryapi_db"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		AdminKey:   getEnv("ADMIN_API_KEY", "default-secret-key"),
	}
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
