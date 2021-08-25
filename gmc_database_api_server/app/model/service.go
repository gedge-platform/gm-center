package model

import (
	"time"
)

type SERVICE struct {
	Name       string       `json:"name"`
	Workspace  string       `json:"workspace"`
	Cluster    string       `json:"cluster"`
	Project    string       `json:"project"`
	Deployment []DEPLOYMENT `json:"deploymentInfo"`
	// PodInfo         []POD                  `json:"podInfo"`
	Type            string      `json:"type"`
	Ports           interface{} `json:"port"`
	ClusterIp       string      `json:"clusterIp"`
	ExternalIp      string      `json:"externalIp"`
	Selector        interface{} `json:"selector"`
	Label           interface{} `json:"label"`
	Annotation      interface{} `json:"annotation"`
	SessionAffinity string      `json:"sessionAffinity"`
	Events          []EVENT     `json:"events"`
	CreateAt        time.Time   `json:"createAt"`
	UpdateAt        time.Time   `json:"updateAt"`
}

type SERVICELISTS struct {
	Pods        interface{}       `json:"pods"`
	Deployments SERVICEDEPLOYMENT `json:"deployments"`
}

type SERVICEDEPLOYMENT struct {
	Name     string    `json:"name"`
	UpdateAt time.Time `json:"updateAt"`
}
