package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/controller"
)

func initRoutes(router *gin.RouterGroup) {
	controller.InitController()

	basePath := "/api/v1"
	v1 := router.Group(basePath)

	v1.POST("/tasks", controller.CreateTaskController)
	v1.GET("/tasks/:id", controller.GetOneTaskController)
	v1.GET("/tasks", controller.GetAllTasksController)
	v1.PUT("/tasks/:id", controller.UpdateTaskController)
}
