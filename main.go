package main

import (
	"github.com/kayooliveira/dinodo-api-go/internal/infra/config"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/router"
)

func main() {
	err := config.InitConfig()

	if err != nil {
		panic("Error on load configs")
	}

	router.InitRouter()
}
