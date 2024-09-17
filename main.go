package main

import (
	"net/http"

	"github.com/t-okuji/demo-todo-go-chi/controller"
	"github.com/t-okuji/demo-todo-go-chi/db"
	"github.com/t-okuji/demo-todo-go-chi/repository"
	"github.com/t-okuji/demo-todo-go-chi/router"
	"github.com/t-okuji/demo-todo-go-chi/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewTaskUsecase(userRepository)
	userController := controller.NewTaskController(userUsecase)
	r := router.NewRouter(userController)

	http.ListenAndServe(":3000", r)
}
