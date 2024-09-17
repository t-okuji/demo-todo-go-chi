package repository

import (
	"errors"

	"github.com/t-okuji/demo-todo-go-chi/model"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTasks() (*[]model.Task, error)
	CreateTask(newTask model.Task) (*model.Task, error)
	UpdateTask(updateTask model.Task) (*model.Task, error)
	DeleteTask(taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetAllTasks() (*[]model.Task, error) {
	var tasks []model.Task
	result := r.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tasks, nil
}

func (r *taskRepository) CreateTask(newTask model.Task) (*model.Task, error) {
	result := r.db.Create(&newTask)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newTask, nil
}

func (r *taskRepository) UpdateTask(updateTask model.Task) (*model.Task, error) {
	result := r.db.Save(&updateTask)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateTask, nil
}

func (r *taskRepository) DeleteTask(taskId uint) error {
	task := model.Task{}
	result := r.db.First(&task, "id = ?", taskId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return errors.New("task not found")
		}
		return result.Error
	}

	result = r.db.Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
