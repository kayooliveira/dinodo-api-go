package main

import (
	"github.com/kayooliveira/dinodo-api-go/internal/infra/config"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.InitConfig()

	if err != nil {
		panic("Error on load configs")
	}

	router.InitRouter()
}
