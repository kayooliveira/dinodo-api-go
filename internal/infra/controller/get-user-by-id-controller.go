package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/usecase"
)

func GetUserByIdController(ctx *gin.Context) {
	repository := repository.NewUserRepositoryMysql(db)

	var input usecase.GetUserByIdInputDto

	usecase := usecase.NewGetUserByIdUseCase(repository)

	input.ID = ctx.Param("id")

	user, err := usecase.Execute(input)

	if err != nil {
		sendError(ctx, 400, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", user)

}
