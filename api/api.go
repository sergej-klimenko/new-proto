package api

import (
	"net/http"
	"new-proto/api/handlers"
	"new-proto/api/repository"
	"new-proto/api/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *http.Server {
	r := chi.NewRouter()

	// global middlware
	r.Use(middleware.Recoverer)

	// repositories
	taskRepo := repository.NewTaskRepository()

	//services
	taskSvc := services.NewTaskService(taskRepo)

	// routes
	r.Route("/api/v1", func(v1 chi.Router) {
		v1.Mount("/tasks", handlers.NewTaskHandler(taskSvc))
		v1.Mount("/env", handlers.NewEnvHandler())
	})

	return &http.Server{
		Handler: r,
		Addr:    ":8888",
	}

}
