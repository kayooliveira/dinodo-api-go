package usecase

import (
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
)

type GetOneTaskInputDto struct {
	ID string `json:"id"`
}
type GetOneTaskUseCase struct {
	TaskRepository entity.TaskRepository
}

func NewGetOneTaskUseCase(repository entity.TaskRepository) *GetOneTaskUseCase {
	return &GetOneTaskUseCase{
		TaskRepository: repository,
	}
}

func (u *GetOneTaskUseCase) Execute(input GetOneTaskInputDto) (*entity.Task, error) {

	task, err := u.TaskRepository.Get(input.ID)

	if err != nil {
		return nil, err
	}

	return task, nil

}
