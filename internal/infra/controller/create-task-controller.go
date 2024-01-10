package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/usecase"
)

func CreateTaskController(ctx *gin.Context) {

	var input usecase.CreateTaskInputDto

	repository := repository.NewTaskRepositoryMysql(db)
	usecase := usecase.NewCreateTaskUseCase(repository)

	err := ctx.BindJSON(&input)

	if err != nil {
		sendError(ctx, 400, "BAD REQUEST")
		return
	}

	userId, exists := ctx.Get("userId")

	userIdString, ok := userId.(string)

	if !exists || !ok {
		sendError(ctx, 500, "Internal server error")
		return
	}

	input.UserID = userIdString

	task, err := usecase.Execute(input)

	if err != nil {
		sendError(ctx, 400, err.Error())
		return
	}

	if task != nil {
		sendCreated(ctx, "create-task", task)
		return
	}

	sendError(ctx, 500, "INTERNAL SERVER ERROR")
}
