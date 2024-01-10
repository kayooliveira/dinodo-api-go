package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
)

type CreateTaskInputDto struct {
	UserID   string `json:"userId"`
	Task     string `json:"task"`
	Finished bool   `json:"finished"`
}

type CreateTaskOutputDto struct {
	ID        uuid.UUID `json:"id"`
	Task      string    `json:"task"`
	UserID    string    `json:"userId"`
	Finished  bool      `json:"finished"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type CreateTaskUseCase struct {
	TaskRepository entity.TaskRepository
}

func NewCreateTaskUseCase(repository entity.TaskRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		TaskRepository: repository,
	}
}

func (u *CreateTaskUseCase) Execute(input CreateTaskInputDto) (*entity.Task, error) {

	if len(input.Task) <= 0 {
		return nil, fmt.Errorf("Task cannot be empty")
	}

	task := entity.Task{
		Task:     input.Task,
		Finished: input.Finished || false,
		UserID:   input.UserID,
	}

	err := u.TaskRepository.Create(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}
