package controller

import (
	"net/http"

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

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if len(input.ID) <= 0 {
		sendError(ctx, http.StatusNotFound, "Task Not Found")
		return
	}

	userId, exists := ctx.Get("userId")

	if !exists {
		sendUnauthorized(ctx, "userId not found on token payload")
		return
	}

	if userId != task.UserID {
		sendForbidden(ctx, "You do not have permission to access the requested resource.")
		return
	}

	sendSuccess(ctx, "get-one-task", task)
}
