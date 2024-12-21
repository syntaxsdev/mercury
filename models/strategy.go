package models

import (
	"errors"
	"fmt"
)

type Strategy struct {
	Name        string             `json:"name"`
	Host        string             `json:"host"`
	Port        int                `json:"port"`
	Options     StrategyOptions    `json:"options"`
	HealthCheck HealthCheckOptions `json:"healthcheck,omitempty"`
}

type HealthCheckOptions struct {
	Path        string      `json:"path"`
	Frequency   int         `json:"frequency"`
	MaxFailures int         `json:"max_failures"`
	Meta        interface{} `json:"meta,omitempty"` // Any dictionary data allowed here
}

type StrategyOptions struct {
	Active bool        `json:"active"`
	Meta   interface{} `json:"meta,omitempty"` // Any dictionary data allowed here
}

// Set defaults
func (s *Strategy) SetDefaults() {
	// Set Healthcheck Path
	if s.HealthCheck.Path == "" {
		s.HealthCheck.Path = "/health"
	}
	// Set Frequency
	if s.HealthCheck.Frequency == 0 {
		s.HealthCheck.Frequency = 5
	}
	// Max Failures
	if s.HealthCheck.MaxFailures == 0 {
		s.HealthCheck.MaxFailures = 5
	}
}

func (s *Strategy) ToUrl() string {
	host := s.Host
	// Remove trailing "/"
	if host[len(host)-1] == '/' {
		host = host[:len(host)-1]
	}
	return fmt.Sprintf("%s:%d", host, s.Port)
}

func (s *Strategy) HasHealthCheck() bool {
	// Check if `HealthCheck` was omitted
	return s.HealthCheck != (HealthCheckOptions{})
}
func (s *Strategy) HealthCheckUrl() (string, error) {
	// Check if `HealthCheck` was omitted
	if !s.HasHealthCheck() {
		return "", errors.New("this strategy has omitted healthcheck data")
	}
	url := s.ToUrl()
	return fmt.Sprintf("%s%s", url, s.HealthCheck.Path), nil
}
