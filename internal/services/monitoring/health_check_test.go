package monitoring

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/syntaxsdev/mercury/models"
)

func TestHealthCheck(t *testing.T) {
	// Mock HTTP Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate healthy check
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Mock Client
	mockClient := &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(server.URL)
			},
		},
	}
	healthChecker := NewHealthChecker(mockClient)

	s := models.Strategy{
		Name:        "Test",
		Host:        "http://localhost",
		Port:        8445,
		Options:     models.StrategyOptions{Active: true},
		HealthCheck: models.HealthCheckOptions{},
	}
	s.SetDefaults()

	t.Run("CheckStrategyHealth", func(t *testing.T) {
		data := HealthData{}
		healthChecker.Check(&data, &s)

		// Asserts
		if data.FailureCount != 0 {
			t.Errorf("Expected 0, got %v", data.FailureCount)
		}
		if !data.Healthy {
			t.Errorf("Expected true got %v", data.Healthy)
		}
	})

	t.Run("CheckUnhealthyStrategy", func(t *testing.T) {
		server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})
		data := HealthData{}
		healthChecker.Check(&data, &s)

		// Asserts
		if data.FailureCount != 1 {
			t.Errorf("Expected 1, got %v", data.FailureCount)
		}
		if data.Healthy {
			t.Errorf("Expected false got %v", data.Healthy)
		}
	})

	t.Run("CheckClosedServer", func(t *testing.T) {
		// Close the server to simulate unavailable resource
		server.Close()

		data := HealthData{}
		healthChecker.Check(&data, &s)

		// Asserts
		if data.FailureCount != 1 {
			t.Errorf("Expected FailureCount 1, got %v", data.FailureCount)
		}
		if data.Healthy {
			t.Errorf("Expected Healthy false got %v", data.Healthy)
		}
	})
}
