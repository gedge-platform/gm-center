package model

import (
	"time"
)

type Cluster struct {
	Num        int       `gorm:"column:clusterNum; primary_key" json:"clusterNum,omitempty"`
	Ip         string    `gorm:"column:ipAddr; not null; default:null" json:"ipAddr,omitempty"`
	Name       string    `gorm:"column:clusterName; not null; default:null" json:"clusterName,omitempty"`
	Type       string    `gorm:"column:clusterType; not null; default:null" json:"clusterType,omitempty"`
	Endpoint   string    `gorm:"column:clusterEndpoint; not null; default:null" json:"clusterEndpoint,omitempty"`
	Creator    string    `gorm:"column:clusterCreator; not null; default:null" json:"clusterCreator,omitempty"`
	Created_at time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	Token      string    `gorm:"column:token; not null; default:null" json:"token,omitempty"`
}
type CLUSTER struct {
	Cluster
		// Status                  string                   `json:"status"`
	// Network                 string                   `json:"network"`
	Gpu                     []map[string]interface{} `json:"gpu"`
	ResourceUsage           interface{}              `json:"resourceUsage"`
	NodeCnt  int `json:"nodeCnt"`
}

type CLUSTER_DETAIL struct {
	Cluster
	Gpu                     []map[string]interface{} `json:"gpu"`
	Resource                PROJECT_RESOURCE         `json:"resource"`
	Nodes []NODE `json:"nodes"`
	Events                  []EVENT                  `json:"events"`
}


type NODE struct {
	Name  string                   `json:"name"`
	NodeType  string                   `json:"type"`
	IP string `json:"nodeIP"`
	Version                 string                   `json:"kubeVersion"`
	Os                      string                   `json:"os,omitempty"`
	Kernel                  string                   `json:"kernel,omitempty"`
	Label                   interface{}              `json:"labels,omitempty"`
	Annotation              interface{}              `json:"annotations,omitempty"`
	CreateAt                time.Time                `json:"created_at"`
	Allocatable             interface{}              `json:"allocatable,omitempty"`
	Capacity                interface{}              `json:"capacity,omitempty"`
	ContainerRuntimeVersion interface{}              `json:"containerRuntimeVersion"`
	
	Addresses               []ADDRESSES              `json:"addresses,omitempty"`
}
type GPU struct {
	Name string `json:"name"`
}
type ADDRESSES struct {
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
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
