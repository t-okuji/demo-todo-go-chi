package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/t-okuji/demo-todo-go-chi/controller"
)

func NewRouter(uc controller.ITaskController) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", uc.GetAllTasks)
		r.Post("/", uc.CreateTask)
		r.Put("/", uc.UpdateTask)
		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", uc.DeleteTask)
		})

	})

	return r
}
