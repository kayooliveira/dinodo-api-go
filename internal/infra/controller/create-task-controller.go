package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/repository"
)

type TaskRequestBody struct {
	Task     string
	Finished bool
}

func CreateTaskController(ctx *gin.Context) {

	var taskRequestBody TaskRequestBody

	err := ctx.BindJSON(&taskRequestBody)

	repository := repository.NewTaskRepositoryMysql(db)

	task := entity.Task{
		Task:     taskRequestBody.Task,
		Finished: taskRequestBody.Finished,
	}

	repository.Create(&task)

	if err != nil {
		sendError(ctx, 400, "BAD REQUEST")
		return
	}

	sendSuccess(ctx, "create-task", task)

}
