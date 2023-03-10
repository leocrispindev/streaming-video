package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
