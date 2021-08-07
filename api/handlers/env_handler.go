package handlers

import (
	"net/http"
	"new-proto/api/config"
	"new-proto/api/utils"

	"github.com/go-chi/chi/v5"
)

func NewEnvHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		env := config.Get("ENVIRONMENT")
		utils.WriteResponse(w, env, 200)
	})
	return r
}
