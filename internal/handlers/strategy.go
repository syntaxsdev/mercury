package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/syntaxsdev/mercury/internal/services"
	"github.com/syntaxsdev/mercury/models"
)

// Get All Strategies
func GetAllStrategies(w http.ResponseWriter, r *http.Request, f *services.Factory) {

}

// Create A Strategy
func CreateStrategy(w http.ResponseWriter, r *http.Request, f *services.Factory) {
	var newStrategy models.Strategy
	if err := json.NewDecoder(r.Body).Decode(&newStrategy); err != nil {

	}
	f.MongoService.Insert("strategies", newStrategy)
}
