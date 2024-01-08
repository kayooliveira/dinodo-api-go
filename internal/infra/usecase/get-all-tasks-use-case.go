package usecase

import (
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
)

type GetAllTasksUseCase struct {
	TaskRepository entity.TaskRepository
}

func NewGetAllTasksUseCase(repository entity.TaskRepository) *GetAllTasksUseCase {
	return &GetAllTasksUseCase{
		TaskRepository: repository,
	}
}

func (u *GetAllTasksUseCase) Execute() ([]*entity.Task, error) {
	tasks, err := u.TaskRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
