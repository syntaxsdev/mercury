package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/syntaxsdev/mercury/internal/services"
	"github.com/syntaxsdev/mercury/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Get All Strategies
func GetAllStrategies(w http.ResponseWriter, r *http.Request, f *services.Factory) {
	var results []*interface{}
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		payload = nil
	}

	if payload != nil {
		var results map[string]interface{}

		err := f.MongoService.First("strategies", bson.M(payload), &results)
		if err == nil && results == nil {
			WriteHttp(w, http.StatusNotFound, "No results were returned with the supplied filter", nil)
			return
		} else if err == nil {
			WriteHttp(w, http.StatusOK, "Successfully fetched all strategies by filter", results)
			return
		}
	}
	err := f.MongoService.All("strategies", bson.M{}, &results)
	if err != nil {
		http.Error(w, "Could not retrieve strategies.", http.StatusInternalServerError)
		return
	}
	WriteHttp(w, http.StatusOK, "Successfully fetched all strategies", results)

}

// Create A Strategy
func CreateStrategy(w http.ResponseWriter, r *http.Request, f *services.Factory) {
	var newStrategy models.Strategy
	if err := json.NewDecoder(r.Body).Decode(&newStrategy); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	res, err := f.MongoService.Insert("strategies", newStrategy)
	if err != nil {
		http.Error(w, "Could not insert strategy", http.StatusInternalServerError)
		return
	}

	WriteHttp(w, http.StatusCreated, "Strategy successfully created", res)
}
