package usecase

import "github.com/kayooliveira/dinodo-api-go/internal/domain/entity"

type DeleteTaskInputDto struct {
	ID string `json:"id"`
}

type DeleteTaskUseCase struct {
	TaskRepository entity.TaskRepository
}

func NewDeleteTaskUseCase(repository entity.TaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		TaskRepository: repository,
	}
}

func (u *DeleteTaskUseCase) Execute(input DeleteTaskInputDto) error {
	task, err := u.TaskRepository.Get(input.ID)

	if err != nil {
		return err
	}

	err = u.TaskRepository.Delete(task)

	if err != nil {
		return err
	}

	return err
}
