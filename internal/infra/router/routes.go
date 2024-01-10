package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/controller"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/middleware"
)

func initRoutes(router *gin.RouterGroup) {
	controller.InitController()

	basePath := "/api/v1"
	v1 := router.Group(basePath)

	v1.POST("/tasks", controller.CreateTaskController) // Create a new task

	v1.GET("/tasks/:id", controller.GetOneTaskController) // Get a single task by ID

	v1.GET("/tasks", controller.GetAllTasksController) // List all tasks

	v1.PUT("/tasks/:id", controller.UpdateTaskController) // Update a task by ID

	v1.DELETE("/tasks/:id", controller.DeleteTaskController) // Delete a task

	v1.POST("/auth/signin", controller.CreateUserController) // Register a new User

	v1.GET("/users/:id", middleware.AuthMiddleware, controller.GetUserByIdController) // Register a new User
}
