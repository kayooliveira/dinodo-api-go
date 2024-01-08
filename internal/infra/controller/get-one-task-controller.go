package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/usecase"
)

func GetOneTaskController(ctx *gin.Context) {
	var input usecase.GetOneTaskInputDto
	repository := repository.NewTaskRepositoryMysql(db)
	usecase := usecase.NewGetOneTaskUseCase(repository)

	input.ID = ctx.Param("id")

	task, err := usecase.Execute(input)

	if len(input.ID) <= 0 {
		sendError(ctx, 404, "Task Not Found")
		return
	}

	if err != nil {
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, "get-one-task", task)
}
