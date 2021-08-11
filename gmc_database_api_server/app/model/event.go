package model

import (
	"time"
)

type EVENT struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Message   string `json:"message"`
	Regarding struct {
		Kind string `json:"kind"`
		Name string `json:"name"`
		// Monitoring []MONITOR `json:"monitoring"`
	}
	Reason    string    `json:"reason"`
	Type      string    `json:"type"`
	EventTime time.Time `json:"eventTime"`
	Note      string    `json:"note"`
	// Monitoring []MONITOR `json:"monitoring"`
}
