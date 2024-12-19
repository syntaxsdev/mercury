package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/syntaxsdev/mercury/models"
)

func WriteHttp(w http.ResponseWriter, code int, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(models.Response{Message: message, Code: code, Data: data})
	if err != nil {
		return err
	}
	return nil
}
