package config

import (
	"os"
	"strconv"
)

type Config struct {
	FrontendPort int
	BackendPort  int
	LogLevel     string
}

var globalConfig Config

func LoadConfig() error {
	frontendPort, _ := strconv.Atoi(getEnv("FRONTEND_PORT", "5173"))
	backendPort, _ := strconv.Atoi(getEnv("BACKEND_PORT", "8080"))

	globalConfig = Config{
		FrontendPort: frontendPort,
		BackendPort:  backendPort,
		LogLevel:     getEnv("LOG_LEVEL", "info"),
	}

	return nil
}

func GetConfig() Config {
	return globalConfig
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
