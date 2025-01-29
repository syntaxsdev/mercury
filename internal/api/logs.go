package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/syntaxsdev/mercury/internal/handlers"
	"github.com/syntaxsdev/mercury/internal/services"
)

func LogRoutes(r chi.Router, factory *services.Factory) {
	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		handlers.GetLog(w, r, factory, name)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllLogs(w, r, factory)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.NewLog(w, r, factory)
	})

	r.Delete("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		handlers.DeleteAllLogs(w, r, factory, name)
	})
}
