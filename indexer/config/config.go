package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	fmt.Println("Config OK")
}
