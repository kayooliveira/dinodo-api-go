package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/controller"
)

func initRoutes(router *gin.RouterGroup) {
	controller.InitController()

	basePath := "/api/v1"
	v1 := router.Group(basePath)

	v1.GET("/tasks", controller.GetAllTasksController)
	v1.POST("/tasks", controller.CreateTaskController)
}
