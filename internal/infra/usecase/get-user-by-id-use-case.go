package usecase

import (
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
)

type GetUserByIdInputDto struct {
	ID string `json:"id"`
}

type GetUserByIdUseCase struct {
	UserRepository entity.UserRepository
}

func NewGetUserByIdUseCase(repository entity.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		UserRepository: repository,
	}
}

func (u *GetUserByIdUseCase) Execute(input GetUserByIdInputDto) (*entity.User, error) {

	user, err := u.UserRepository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
