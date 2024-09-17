package main

import (
	"fmt"

	"github.com/t-okuji/demo-todo-go-chi/db"
	"github.com/t-okuji/demo-todo-go-chi/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Task{})
}
