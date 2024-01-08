package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/usecase"
)

func GetAllTasksController(ctx *gin.Context) {
	repository := repository.NewTaskRepositoryMysql(db)
	usecase := usecase.NewGetAllTasksUseCase(repository)

	tasks, err := usecase.Execute()

	if err != nil {
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, "get-all-tasks", tasks)
}
