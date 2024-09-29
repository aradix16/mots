package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	envFile := fmt.Sprintf("%s.env", environment)

	if err := godotenv.Load(envFile); err != nil {
		return nil, fmt.Errorf("error loading %s file: %w", envFile, err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT must be set in %s", envFile)
	}

	config := &Config{
		Port: port,
	}

	return config, nil
}
