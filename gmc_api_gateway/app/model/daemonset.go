package model

import (
	"time"
)

type DAEMONSET struct {
	Name        string           `json:"name"`
	Namespace   string           `json:"namespace"`
	AccessMode      []string `json:"accessMode"`
	ReclaimPolicy         string     `json:"reclaimPolicy",`
	Status    interface{}      `json:"status"`
	Claim   interface{}        `json:"claim"`
	StorageClass       string           `json:"storageClass"`
	// Reason        []EVENT          `json:"events"`
	VolumeMode string  `json:"volumeMode"`
	Cluster string  `json:"cluster"`
	Workspace string `json:"workspace"`
	CreateAt time.Time          `json:"createAt"`
	Selector             interface{}        `json:"selector,omitempty"`
	Annotations       interface{}        `json:"annotations,omitempty"`
	Events  []EVENT          `json:"events"`
}


type DAEMONSETS []DAEMONSET

func (DAEMONSET) TableName() string {
	return "DAEMONSET_INFO"
}