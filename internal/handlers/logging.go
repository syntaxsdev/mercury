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
func GetLog(w http.ResponseWriter, r *http.Request, f *services.Factory, n string) {
	var logs []*interface{}
	filter := bson.M{"strategy": n}

	err := f.MongoService.All("logs", filter, &logs)
	if err != nil {
		WriteHttp(w, http.StatusInternalServerError, "Failed to retrieve logs.", err)
		return
	}
	WriteHttp(w, http.StatusOK, "Successfully fetched all logs", logs)
}

// Get log of a specific strategy
func DeleteAllLogs(w http.ResponseWriter, r *http.Request, f *services.Factory, n string) {
	filter := bson.M{"strategy": n}
	_, err := f.MongoService.DeleteAll("logs", filter)
	if err != nil {
		WriteHttp(w, http.StatusInternalServerError, "Failed to retrieve logs.", err)
		return
	}
	WriteHttp(w, http.StatusOK, "Successfully deleted all logs.", nil)
}

// Creates a new log
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
