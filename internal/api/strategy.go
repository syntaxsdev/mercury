package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/syntaxsdev/mercury/internal/handlers"
	"github.com/syntaxsdev/mercury/internal/services"
)

func StrategyRoutes(r chi.Router, factory *services.Factory) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllStrategies(w, r, factory)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateStrategy(w, r, factory)
	})
}
