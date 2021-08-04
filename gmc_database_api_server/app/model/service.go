package model

import "time"

type SERVICE struct {
	Workspace []Workspace `json:"workspace"`
	Cluster   []Cluster   `json:"cluster"`
	Project   []Project   `json:"project"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Ports     []PORT      `json:"port"`
	ClusterIp string      `json:"clusterIp"`
	CreateAt  time.Time   `json:"createAt"`

	// jwt.StandardClaim
}

type SERVICE_DETAIL struct {
	Name           string       `json:"Name"`
	Clusters       []Cluster    `json:"cluster"`
	Project        []Project    `json:"project"`
	DeploymentInfo []DEPLOYMENT `json:"deploymentInfo"`
	// PodInfo         []POD                  `json:"podInfo"`
	ExternalIp      string                 `json:"externalIp"`
	Selector        map[string]interface{} `json:"selector"`
	Label           map[string]interface{} `json:"label"`
	Annotation      map[string]interface{} `json:"annotation"`
	SessionAffinity string                 `json:"sessionAffinity"`
	Ports           []PORT                 `json:"port"`
	Events          []EVENT                `json:"events"`
	CreateAt        time.Time              `json:"createAt"`
	UpdateAt        time.Time              `json:"updateAt"`

	// jwt.StandardClaim
}
type PORT struct {
	name       string `json:"name"`
	port       string `json:"port"`
	protocol   string `json:"protocol"`
	targetPort string `json:"targetPort"`
	// jwt.StandardClaim
}
