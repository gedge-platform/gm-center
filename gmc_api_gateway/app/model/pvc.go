package model

import (
	"time"
)

type PVC struct {
	Name        string           `json:"name"`
	Capacity   string           `json:"capacity"`
	AccessMode      []string `json:"accessMode"`
	Status    interface{}      `json:"status"`
	Volume   interface{}        `json:"volume"`
	StorageClass       string           `json:"storageClass"`
	Cluster string  `json:"clusterName"`
	// Reason        []EVENT          `json:"events"`
	CreateAt time.Time          `json:"createAt"`
	Events  []EVENT          `json:"events"`
}


type PVCs []PVC

func (PVC) TableName() string {
	return "PVC_INFO"
}