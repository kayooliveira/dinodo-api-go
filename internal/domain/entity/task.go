package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Task     string    `json:"task"`
	Finished bool      `json:"finished"`
	UserID   string    `json:"userId"`
	// User      User            `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type TaskRepository interface {
	Create(task *Task) error
	Update(task *Task) error
	Get(id string) (*Task, error)
	GetAll() ([]*Task, error)
	Delete(task *Task) error
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
