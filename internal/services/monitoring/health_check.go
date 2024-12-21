package monitoring

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/syntaxsdev/mercury/internal/services"
	"github.com/syntaxsdev/mercury/models"
)

// HealthData tracks data of each Strategies health
type HealthData struct {
	mu           sync.Mutex
	Healthy      bool
	LastCheck    time.Time
	FailureCount int
}

// Update updates the health check result
func (h *HealthData) Update(isHealthy bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Healthy = isHealthy
	h.LastCheck = time.Now()
	if !isHealthy {
		h.FailureCount++
		return
	}
	h.FailureCount = 0
}

// HealthChecker
type HealthChecker struct {
	client     *http.Client
	mu         sync.Mutex
	healthData map[string]*HealthData
}

// Initialize a HealthChecker Manager
func NewHealthChecker(client *http.Client) *HealthChecker {
	if client == nil {
		client = &http.Client{}
	}
	return &HealthChecker{
		healthData: make(map[string]*HealthData),
		client:     client,
	}
}

func (h *HealthChecker) Check(data *HealthData, s *models.Strategy) {
	if s.Options.Active && s.HasHealthCheck() {
		url, _ := s.HealthCheckUrl()
		resp, err := h.client.Get(url)
		if err != nil || (resp.StatusCode < 200 || resp.StatusCode > 203) {
			data.Update(false)
		} else {
			data.Update(true)
		}
	}
	// Max 5 unhealthy checks
	if data.FailureCount == 5 {
		s.Options.Active = false
	}
}

func (h *HealthChecker) BackgroundProcess(s *services.StrategyService) {
	defaultWait := 10
	for {
		strats, err := s.GetAllStrategies()
		if err != nil {
			log.Println("ERROR: Cannot get strategies")
			continue
		}

		// Delay the wait if there is no strategies at the moment
		if len(strats) == 0 {
			defaultWait = 20
			continue
		}
		defaultWait = 10

		for _, strat := range strats {
			h.mu.Lock()
			if _, exists := h.healthData[strat.Name]; !exists {
				h.healthData[strat.Name] = &HealthData{}
			}
			h.mu.Unlock()
			val := h.healthData[strat.Name]
			go h.Check(val, &strat)
		}
		time.Sleep(time.Duration(defaultWait) * time.Second)
	}
}
