package monitoring

import (
	"net/http"

	"github.com/syntaxsdev/mercury/internal/services"
)

type Monitor struct {
	db            *services.MongoService
	client        *http.Client
	Strategy      *services.StrategyService
	HealthChecker *HealthChecker
}

// Create a new Monitoring Service
func NewMonitor(d *services.MongoService, client *http.Client) *Monitor {
	if client == nil {
		client = &http.Client{}
	}
	return &Monitor{
		db:            d,
		client:        client,
		Strategy:      services.NewStrategyService(d),
		HealthChecker: NewHealthChecker(client),
	}

}

// Start the Monitoring
func (m *Monitor) Start() {
	// Run HealthChecker in the background
	go m.HealthChecker.BackgroundProcess(m.Strategy)
}
