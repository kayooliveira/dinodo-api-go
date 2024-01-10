package repository

import (
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryMysql struct {
	DB *gorm.DB
}

func NewUserRepositoryMysql(db *gorm.DB) *UserRepositoryMysql {
	return &UserRepositoryMysql{
		DB: db,
	}
}

func (r *UserRepositoryMysql) Create(user *entity.User) error {
	err := r.DB.Create(&user)

	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (r *UserRepositoryMysql) GetById(id string) (*entity.User, error) {
	user := entity.User{}
	result := r.DB.Preload("Tasks").First(&user, "id = ? ", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
