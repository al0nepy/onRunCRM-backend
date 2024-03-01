package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(name string) string {
	env := os.Getenv(name)
	return env
}
