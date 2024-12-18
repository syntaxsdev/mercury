package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

// Convert Response data to JSON and send to http writer
func (r *Response) SendJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}
