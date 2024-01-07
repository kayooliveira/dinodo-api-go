package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitRouter() {

	router := gin.Default()

	initRoutes(&router.RouterGroup)

	if err := router.Run("127.0.0.1:4000"); err != nil {
		log.Fatal(err)
	}
}
