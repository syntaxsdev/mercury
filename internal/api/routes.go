package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/syntaxsdev/mercury/internal/services"
)

func InitRoutes(factory *services.Factory) http.Handler {
	r := chi.NewRouter()

	r.Route("/strategy", func(r chi.Router) {
		StrategyRoutes(r, factory)
	})

	r.Route("/logs", func(r chi.Router) {
		LogRoutes(r, factory)
	})
	return r
}
