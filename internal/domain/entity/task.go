package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID        uuid.UUID       `json:"id" gorm:"type:char(36);primary_key"`
	Task      string          `json:"task"`
	Finished  bool            `json:"finished"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type TaskRepository interface {
	Create(task *Task) error
	Update(task *Task) error
	Get(id int) (*Task, error)
	GetAll() ([]*Task, error)
	Delete(id int) (bool, error)
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
