package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/usecase"
)

func UpdateTaskController(ctx *gin.Context) {
	id := ctx.Param("id")

	repository := repository.NewTaskRepositoryMysql(db)

	var input usecase.UpdateTaskInputDto

	usecase := usecase.NewUpdateTaskUseCase(repository)

	err := ctx.BindJSON(&input)
	if err != nil {
		sendError(ctx, 400, "BAD REQUEST")
		return
	}

	task, err := usecase.Execute(id, input)

	if err != nil {
		sendError(ctx, 400, err.Error())
		return
	}

	sendSuccess(ctx, "update-task", task)

}
