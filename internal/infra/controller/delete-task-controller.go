package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/usecase"
)

func DeleteTaskController(ctx *gin.Context) {
	repository := repository.NewTaskRepositoryMysql(db)

	var input usecase.DeleteTaskInputDto

	usecase := usecase.NewDeleteTaskUseCase(repository)

	input.ID = ctx.Param("id")

	err := usecase.Execute(input)

	if err != nil {
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, "delete-task", "Task deleted")
}
