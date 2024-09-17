package usecase

import (
	"github.com/t-okuji/demo-todo-go-chi/dto"
	"github.com/t-okuji/demo-todo-go-chi/model"
	"github.com/t-okuji/demo-todo-go-chi/repository"
)

type ITaskUsecase interface {
	GetAllTasks() (*[]model.Task, error)
	CreateTask(createTaskInput dto.CreateTaskInput) (*model.Task, error)
	UpdateTask(updateTaskInput dto.UpdateTaskInput) (*model.Task, error)
	DeleteTask(taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAllTasks() (*[]model.Task, error) {
	return tu.tr.GetAllTasks()
}

func (tu *taskUsecase) CreateTask(createTaskInput dto.CreateTaskInput) (*model.Task, error) {
	newTask := model.Task{
		Task:   createTaskInput.Task,
		Status: createTaskInput.Status,
	}
	return tu.tr.CreateTask(newTask)
}

func (tu *taskUsecase) UpdateTask(updateTaskInput dto.UpdateTaskInput) (*model.Task, error) {
	newTask := model.Task{
		Id:     updateTaskInput.Id,
		Task:   updateTaskInput.Task,
		Status: updateTaskInput.Status,
	}
	return tu.tr.UpdateTask(newTask)
}

func (tu *taskUsecase) DeleteTask(taskId uint) error {
	return tu.tr.DeleteTask(taskId)
}
