package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/syntaxsdev/mercury/models"
)

func TestWriteHttp(t *testing.T) {
	recorder := httptest.NewRecorder()
	data := map[string]interface{}{"data": map[string]string{"name": "testing"}}

	err := WriteHttp(recorder, http.StatusOK, "success", data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Check response body
	var response models.Response

	if err = json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Error("Could not convert to Response type")
	}
	if dataMap, ok := response.Data.(map[string]interface{}); ok {
		if name, exists := dataMap["name"]; exists && name != "testing" {
			t.Errorf("Expected %v, got %v", response.Data, data)
		}
	}
}
