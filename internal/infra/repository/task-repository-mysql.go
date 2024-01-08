package repository

import (
	"github.com/kayooliveira/dinodo-api-go/internal/domain/entity"
	"gorm.io/gorm"
)

type TaskRepositoryMysql struct {
	DB *gorm.DB
}

func NewTaskRepositoryMysql(db *gorm.DB) *TaskRepositoryMysql {
	return &TaskRepositoryMysql{DB: db}
}

func (r *TaskRepositoryMysql) Create(task *entity.Task) error {
	result := r.DB.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepositoryMysql) Update(task *entity.Task) error {

	result := r.DB.Model(&task).Updates(map[string]interface{}{
		"Task":     task.Task,
		"Finished": task.Finished,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TaskRepositoryMysql) Get(id string) (*entity.Task, error) {
	task := entity.Task{}
	result := r.DB.First(&task, "id = ? ", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (r *TaskRepositoryMysql) GetAll() ([]*entity.Task, error) {
	tasks := []*entity.Task{}

	result := r.DB.Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *TaskRepositoryMysql) Delete(id string) error {
	task := entity.Task{}

	result := r.DB.Find(&task, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	r.DB.Delete(task)

	return nil
}
