package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
)

func GetAllTasksController(ctx *gin.Context) {
	repository := repository.NewTaskRepositoryMysql(db)

	tasks, err := repository.GetAll()

	if err != nil {
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, "get-all-tasks", tasks)
}
