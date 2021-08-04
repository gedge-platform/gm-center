package model

import "time"

type DEPLOYMENT struct {
	Name     string    `json:"name"`
	Stauts   string    `json:"stauts"`
	Replica  REPLICA   `json:"replica"`
	UpdateAt time.Time `json:"updateAt"`
	// jwt.StandardClaim
}

type DEPLOYMENT_DETAIL struct {
	Name       string      `json:"name"`
	Stauts     string      `json:"stauts"`
	Strategy   STRATEGY    `json:"strategy"`
	Replica    REPLICA     `json:"replica"`
	Containers []CONTAINER `json:"containers"`
	// PodInfo     []model.Pod     `json:"pods"`
	// ServiceInfo []model.Service `json:"services"`
	Lable      map[string]interface{} `json:"lables"`
	Events     []EVENT                `json:"events"`
	Annotation map[string]interface{} `json:"annotations"`
	CreateAt   time.Time              `json:"createAt"`
	UpdateAt   time.Time              `json:"updateAt"`
	// jwt.StandardClaim
}
type STRATEGY struct {
	Type           string `json:"type"`
	MaxSurge       string `json:"maxSurge"`
	MaxUnavailable string `json:"maxUnavailable"`
	// jwt.StandardClaim
}

type REPLICA struct {
	Replicas            int `json:"replicas"`
	ReadyReplicas       int `json:"readyReplicas"`
	UpdatedReplicas     int `json:"updatedReplicas"`
	AvailableReplicas   int `json:"availableReplicas"`
	UnavailableReplicas int `json:"unavailableReplicas"`
	// jwt.StandardClaim
}
type CONTAINER struct {
	Image    string              `json:"image"`
	Name     string              `json:"name"`
	Resource DEPLOYMENT_RESOURCE `json:"resource"`
}

type DEPLOYMENT_RESOURCE struct {
	Limit   map[string]interface{} `json:"limit"`
	Request map[string]interface{} `json:"request"`
}
