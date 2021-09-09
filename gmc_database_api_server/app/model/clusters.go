package model

import (
	"time"
)

type Cluster struct {
	Num        int       `gorm:"column:clusterNum; primary_key" json:"clusterNum"`
	Ip         string    `gorm:"column:ipAddr; not null; default:null" json:"ipAddr"`
	Name       string    `gorm:"column:clusterName; not null; default:null" json:"clusterName"`
	Type       string    `gorm:"column:clusterType; not null; default:null" json:"clusterType"`
	Endpoint   string    `gorm:"column:clusterEndpoint; not null; default:null" json:"clusterEndpoint"`
	Creator    string    `gorm:"column:clusterCreator; not null; default:null" json:"clusterCreator"`
	Created_at time.Time `gorm:"column:created_at" json:"created_at"`
	Token      string    `gorm:"column:token; not null; default:null" json:"token"`
	// Monitoring []MONITOR `json:"monitoring"`
}

type CLUSTER struct {
	Cluster
	Gpu           []map[string]interface{} `json:"gpu"`
	Version       string                   `json:"kubeVersion"`
	Status        string                   `json:"status"`
	Network       string                   `json:"network"`
	Os            string                   `json:"os"`
	Kernel        string                   `json:"kernel"`
	Label         interface{}              `json:"lables"`
	Annotation    interface{}              `json:"annotations"`
	CreateAt      time.Time                `json:"created_at"`
	ResourceUsage interface{}              `json:"resourceUsage"`
	Allocatable   interface{}              `json:"allocatable"`
	Capacity      interface{}              `json:"capacity"`
	Resource      PROJECT_RESOURCE         `json:"resource"`
	Events        []EVENT                  `json:"events"`
}
type GPU struct {
	Name string `json:"name"`
}

// type CLUSTER struct {
// 	// Cluster
// 	Name       string            `json:"name"`
// 	Status     string            `json:"status"`
// 	IP         string            `json:"ip"`
// 	Role       string            `json:"role"`
// 	Network    string            `json:"network"`
// 	Os         string            `json:"os"`
// 	Type       string            `json:"type"`
// 	Kernel     string            `json:"kernel"`
// 	kubernetes string            `json:"kubernetes"`
// 	Lable      map[string]string `json:"lables"`
// 	Annotation map[string]string `json:"annotations"`
// 	CreatedAt  time.Time         `json:"created_at"`
// 	// Pod        []Pod    `json:"pods"`
// 	// Metadata   Metadata `json:"metadata"`
// 	Events []EVENT `json:"events"`
// 	// Monitoring model.monitoring `json:"monitoring"`
// }

// Set Cluster table name to be `CLUSTER_INFO`
func (Cluster) TableName() string {
	return "CLUSTER_INFO"
}

type CLUSTERS []CLUSTER
