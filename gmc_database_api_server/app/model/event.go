package model

import (
	"time"
)

type EVENT struct {
	Kind      string    `json:"kind"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Message   string    `json:"message"`
	Reason    string    `json:"reason"`
	Type      string    `json:"type"`
	EventTime time.Time `json:"eventTime"`
	// Monitoring []MONITOR `json:"monitoring"`
}
