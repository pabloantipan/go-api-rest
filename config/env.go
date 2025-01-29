package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file at package initialization
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}
}
