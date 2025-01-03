package models

import "time"

type Log struct {
	Strategy  string                 `json:"strategy" bson:"strategy"`
	Timestamp time.Time              `json:"timestamp" bson:"timestamp"`
	Message   string                 `json:"message" bson:"message"`
	Meta      map[string]interface{} `json:"meta" bson:"meta"`
}
