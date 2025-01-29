package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatastoreServiceAccountPath string
	Port                        string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{
		DatastoreServiceAccountPath: os.Getenv("DATASTORE_SERVICE_ACCOUNT_PATH"),
		Port:                        os.Getenv("PORT"),
	}

	return config, nil
}
