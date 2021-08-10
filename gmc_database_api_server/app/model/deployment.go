package model

import "time"

type DEPLOYMENT struct {
	Name          string      `json:"name"`
	Namespace     string      `json:"project"`
	ClusterName   string      `json:"cluster"`
	WorkspaceName string      `json:"workspace"`
	Stauts        string      `json:"stauts"`
	Replica       REPLICA     `json:"replica"`
	Strategy      STRATEGY    `json:"strategy"`
	Containers    []CONTAINER `json:"containers"`
	// PodInfo     []model.Pod     `json:"pods"`
	// ServiceInfo []model.Service `json:"services"`
	Lable      map[string]string     `json:"lables"`
	Events     []EVENT               `json:"events"`
	Annotation map[string]string     `json:"annotations"`
	CreateAt   time.Time             `json:"createAt"`
	UpdateAt   time.Time             `json:"updateAt"`
	Resource   []DEPLOYMENT_RESOURCE `json:"resource"`
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
	Image    string `json:"image"`
	Name     string `json:"name"`
	Resource struct {
		Limit struct {
			Cpu    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"limits"`
		Request struct {
			Cpu    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"requests"`
	} `json:"resources"`
}

type DEPLOYMENT_RESOURCE struct {
}
