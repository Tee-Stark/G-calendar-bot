package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
