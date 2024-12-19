package models

type Strategy struct {
	Name    string          `json:"name"`
	Host    string          `json:"host"`
	Port    int             `json:"port"`
	Options StrategyOptions `json:"options"`
}

type StrategyOptions struct {
	Active bool        `json:"active"`
	Meta   interface{} `json:"meta,omitempty"` // Any dictionary data allowed here
}
