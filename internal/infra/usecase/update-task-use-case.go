package usecase

import (
	"fmt"

	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
)

type UpdateTaskInputDto struct {
	Task     *string `json:"task"`
	Finished *bool   `json:"finished"`
}

type UpdateTaskUseCase struct {
	TaskRepository entity.TaskRepository
}

func NewUpdateTaskUseCase(repository entity.TaskRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		TaskRepository: repository,
	}
}

func (u *UpdateTaskUseCase) Execute(id string, input UpdateTaskInputDto) (*entity.Task, error) {

	if len(id) <= 0 {
		return nil, fmt.Errorf("Please provide a valid id")
	}

	task, err := u.TaskRepository.Get(id)

	if (UpdateTaskInputDto{}) == input {
		return nil, fmt.Errorf("Please provide the data to update")
	}

	if err != nil {
		return nil, err
	}

	if input.Task != nil {
		task.Task = *input.Task
	}

	if input.Finished != nil {
		task.Finished = *input.Finished
	}

	err = u.TaskRepository.Update(id, task)

	if err != nil {
		return nil, err
	}

	return task, nil
}
