package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/t-okuji/demo-todo-go-chi/dto"
	"github.com/t-okuji/demo-todo-go-chi/usecase"
)

type ITaskController interface {
	GetAllTasks(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := tc.tu.GetAllTasks()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, tasks)
}

func (tc *taskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateTaskInput{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Println(input)

	newTask, err := tc.tu.CreateTask(input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, newTask)
}

func (tc *taskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	input := dto.UpdateTaskInput{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	newTask, err := tc.tu.UpdateTask(input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, newTask)
}

func (tc *taskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = tc.tu.DeleteTask(uint(taskId))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, taskId)
}
