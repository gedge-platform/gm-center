package model

import (
	"time"
)

type EVENT struct {
	Metadata struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
	} `json:"metadata"`
	Regarding struct {
		Kind string `json:"kind"`
		Name string `json:"name"`
	} `json:"regarding"`
	Reason string `json:"reason"`
	Type   string `json:"type"`
	Note   string `json:"note"`
}
