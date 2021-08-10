package model

import (
	"time"

	"k8s.io/apimachinery/pkg/util/intstr"
)

type SERVICE struct {
	Name       string     `json:"name"`
	Workspace  string     `json:"workspace"`
	Cluster    string     `json:"cluster"`
	Project    string     `json:"project"`
	Deployment Deployment `json:"deploymentInfo"`
	// PodInfo         []POD                  `json:"podInfo"`
	Type            string                 `json:"type"`
	Ports           []PORT                 `json:"port"`
	ClusterIp       string                 `json:"clusterIp"`
	ExternalIp      string                 `json:"externalIp"`
	Selector        map[string]interface{} `json:"selector"`
	Label           map[string]interface{} `json:"label"`
	Annotation      map[string]interface{} `json:"annotation"`
	SessionAffinity string                 `json:"sessionAffinity"`
	Events          []EVENT                `json:"events"`
	CreateAt        time.Time              `json:"createAt"`
	UpdateAt        time.Time              `json:"updateAt"`
}

type PORT struct {
	name       string             `json:"name"`
	port       int32              `json:"port"`
	protocol   Protocol           `json:"protocol"`
	targetPort intstr.IntOrString `json:"targetPort"`
}
