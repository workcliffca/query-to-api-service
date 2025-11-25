package config

import (
	"os"
)

type Config struct {
	DBType     string // "postgres" or "mssql"
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	AdminKey   string
}

func LoadConfig() *Config {
	dbType := getEnv("DB_TYPE", "postgres")
	defaultPort := "5432"
	defaultUser := "queryapi_user"
	defaultPassword := "queryapi_password"

	// Adjust defaults based on database type
	if dbType == "mssql" {
		defaultPort = "1433"
		defaultUser = "azuresql_admin"
		defaultPassword = ""
	}

	return &Config{
		DBType:     dbType,
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", defaultPort),
		DBUser:     getEnv("DB_USER", defaultUser),
		DBPassword: getEnv("DB_PASSWORD", defaultPassword),
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
