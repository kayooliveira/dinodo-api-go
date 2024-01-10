package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/usecase"
)

func CreateUserController(ctx *gin.Context) {
	repository := repository.NewUserRepositoryMysql(db)

	var input usecase.CreateUserInputDto

	usecase := usecase.NewCreateUserUseCase(repository)

	err := ctx.BindJSON(&input)

	if err != nil {
		sendError(ctx, 400, err.Error())
		return
	}

	result, err := usecase.Execute(input)

	if err != nil {
		sendError(ctx, 400, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", gin.H{
		"user":        result.User,
		"accessToken": result.AccessToken,
	})

}
