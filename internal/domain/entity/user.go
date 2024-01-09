package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID       `json:"id" gorm:"type:char(36);primary_key"`
	Username  string          `json:"username"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Tasks     []Task          `json:"tasks"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Get(id string) (*User, error)
	GetAll() ([]*User, error)
	Delete(user *User) error
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	u.ID = uuid.New()
	return
}
