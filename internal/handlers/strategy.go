package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/syntaxsdev/mercury/internal/services"
	"github.com/syntaxsdev/mercury/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Get All Strategies
func GetAllStrategies(w http.ResponseWriter, r *http.Request, f *services.Factory) {
	var results []interface{}
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		payload = nil
	}

	if payload != nil {
		err := f.MongoService.First("strategies", bson.M(payload), &results)
		if err == nil && results == nil {
			WriteHttp(w, http.StatusNotFound, "Nothing was found with the supplied filter", nil)
			return
		} else if err == nil {
			log.Println("Here")
			WriteHttp(w, http.StatusOK, "Successfully fetched all strategies by filter", results)
		}
		log.Println("Here", err)
	}
	err := f.MongoService.All("strategies", bson.M{}, &results)
	if err != nil {
		http.Error(w, "Could not retrieve strategies.", http.StatusInternalServerError)
		return
	}
	code := http.StatusOK
	WriteHttp(w, code, "Successfully fetched all strategies", results)

}

// // Get Strategy By ID
// func GetStrategyByID(w http.ResponseWriter, r *http.Request, f *services.Factory) {
// 	var result interface{}
// 	id := chi.URLParam(r, "id")
// 	port, err := strconv.Atoi(id)
// 	if err != nil {
// 		WriteHttp(w, http.StatusBadRequest, "Invalid", err)
// 	}
// 	// objectID, err := primitive.ObjectIDFromHex(id)
// 	// if err != nil {
// 	// 	WriteHttp(w, http.StatusBadRequest, "Invalid ID", err)
// 	// 	return
// 	// }

// 	err = f.MongoService.First("strategies", bson.M{"port": port}, &result)
// 	if err != nil {
// 		WriteHttp(w, http.StatusBadRequest, "Could not retrieve strategy by ID.", nil)
// 		log.Println(err)
// 		return
// 	}
// 	WriteHttp(w, http.StatusOK, "Successfully fetched strategy", result)

// }

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
	}

	WriteHttp(w, http.StatusCreated, "Strategy successfully created", res)
	// code := http.StatusCreated
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(models.Response{Message: "Strategy successfully created", Code: code, Data: res})
}
