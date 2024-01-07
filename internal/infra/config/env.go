package config

import (
	"github.com/joho/godotenv"
)

func initEnv() {
	err := godotenv.Load()

	if err != nil {
		logger.Err("Error on load .env file")
	}
}
