package main

import (
	"fmt"
	"os"
)

type Config struct {
	AppHost string
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func loadConfig() (Config, error) {
	config := Config{
		AppHost: getEnv("APP_HOST", "127.0.0.1"),
		AppPort: getEnv("APP_PORT", "3000"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	if config.DBHost == "" {
		return Config{}, fmt.Errorf("DB_HOST is required")
	}

	if config.DBUser == "" {
		return Config{}, fmt.Errorf("DB_USER is required")
	}

	if config.DBPassword == "" {
		return Config{}, fmt.Errorf("DB_PASSWORD is required")
	}

	if config.DBName == "" {
		return Config{}, fmt.Errorf("DB_NAME is required")
	}

	return config, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}