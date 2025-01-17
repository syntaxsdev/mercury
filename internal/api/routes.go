package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/syntaxsdev/mercury/internal/handlers"
	"github.com/syntaxsdev/mercury/internal/services"
	version "github.com/syntaxsdev/mercury/shared"
)

func InitRoutes(factory *services.Factory) http.Handler {
	r := chi.NewRouter()

	// Base endpoint / healthy status
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.WriteHttp(w, 200, "Mercury is up and running.", map[string]interface{}{
			"status":  "healthy",
			"version": version.Version,
		})
	})

	r.Route("/strategy", func(r chi.Router) {
		StrategyRoutes(r, factory)
	})

	r.Route("/logs", func(r chi.Router) {
		LogRoutes(r, factory)
	})
	return r
}
