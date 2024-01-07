package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func initEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("[ERROR] Error on load .env")
	}
}
