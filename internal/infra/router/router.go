package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	router := gin.Default()

	initRoutes(&router.RouterGroup)
	port := os.Getenv("PORT")
	appUrl := fmt.Sprintf("127.0.0.1:%s", port)

	if err := router.Run(appUrl); err != nil {
		log.Fatal(err)
	}
}
