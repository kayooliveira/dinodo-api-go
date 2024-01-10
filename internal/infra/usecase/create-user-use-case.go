package usecase

import (
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/auth"
)

type CreateUserInputDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type CreateUserOutputDto struct {
	User        *entity.User `json:"user"`
	AccessToken string       `json:"accessToken"`
}

type CreateUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewCreateUserUseCase(repository entity.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: repository,
	}
}

func (u *CreateUserUseCase) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {

	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
	}

	err := u.UserRepository.Create(&user)

	if err != nil {
		return nil, err
	}

	accessToken, err := auth.GenerateUserToken(user.ID.String(), user.Name, user.Email)

	if err != nil {
		accessToken = ""
	}

	createUserOutputDto := CreateUserOutputDto{
		User:        &user,
		AccessToken: accessToken,
	}

	return &createUserOutputDto, nil
}
