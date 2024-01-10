package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID       `json:"id" gorm:"type:char(36);primary_key"`
	Username  string          `json:"username" gorm:"unique"`
	Name      string          `json:"name"`
	Email     string          `json:"email" gorm:"unique"`
	Password  string          `json:"-"`
	Tasks     []Task          `json:"tasks"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type UserRepository interface {
	Create(user *User) error
	GetById(id string) (*User, error)
	// Update(user *User) error
	// Get(id string) (*User, error)
	// GetAll() ([]*User, error)
	// Delete(user *User) error
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	u.ID = uuid.New()

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}

	u.Password = string(password)

	return
}
