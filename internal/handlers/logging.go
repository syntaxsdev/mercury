package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/syntaxsdev/mercury/internal/services"
	"github.com/syntaxsdev/mercury/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Create a new log
func GetAllLogs(w http.ResponseWriter, r *http.Request, f *services.Factory) {
	var logs []*interface{}
	err := f.MongoService.All("logs", bson.M{}, &logs)
	if err != nil {
		WriteHttp(w, http.StatusInternalServerError, "Failed to retrieve logs.", err)
		return
	}
	WriteHttp(w, http.StatusOK, "Successfully fetched all logs", logs)
}

// Get log of a specific strategy
// func GetLog(w http.ResponseWriter, r *http.Request, f *services.Factory) {
// 	var logs []*interface{}
// 	var filterPayload map[string]interface{}

// 	var payload map[string]interface{}
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		payload = nil
// 	}

// 	err := f.MongoService.All("logs", bson.M(filterPayload), &logs)
// 	if err != nil {
// 		WriteHttp(w, http.StatusInternalServerError, "Failed to retrieve logs.", err)
// 		return
// 	}
// 	WriteHttp(w, http.StatusOK, "Successfully fetched all logs", logs)
// }

// Create a new log
func NewLog(w http.ResponseWriter, r *http.Request, f *services.Factory) {
	var newLog models.Log
	if err := json.NewDecoder(r.Body).Decode(&newLog); err != nil {
		WriteHttp(w, http.StatusBadRequest, "Invalid Log object in payload!", err)
		return
	}
	if _, err := f.MongoService.Insert("logs", newLog); err != nil {
		WriteHttp(w, http.StatusInternalServerError, "Could not add log to database.", err)
		return
	}
	WriteHttp(w, http.StatusCreated, "Success", nil)
}
