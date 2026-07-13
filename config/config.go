package config

import (
	"os"
	"strconv"
)

type CorsConfig struct {
	Origin string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	DebugMode bool
	Cors      CorsConfig
	Database  DatabaseConfig
}

func New() *Config {
	return &Config{
		DebugMode: getEnvAsBool("DEBUG_MODE", true),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Database: getEnv("DB_NAME", "postgres"),
		},
		Cors: CorsConfig{
			Origin: getEnv("ALLOW_ORIGIN", "http://localhost"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
