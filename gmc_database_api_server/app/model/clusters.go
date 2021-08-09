package model

import (
	"time"
)

// type Cluster struct {
// 	Num        int       `gorm:"column:clusterNum; primary_key" json:"clusterNum"`
// 	Ip         string    `gorm:"column:ipAddr; not null; default:null" json:"ipAddr"`
// 	ExtIp      string    `gorm:"column:extIpAddr" json:"extIpAddr"`
// 	Name       string    `gorm:"column:clusterName; not null; default:null" json:"clusterName"`
// 	Role       string    `gorm:"column:clusterRole; not null; default:null" json:"clusterRole"`
// 	Type       string    `gorm:"column:clusterType; not null; default:null" json:"clusterType"`
// 	Gpu        string    `json:"gpu"`
// 	Endpoint   string    `gorm:"column:clusterEndpoint; not null; default:null" json:"clusterEndpoint"`
// 	Creator    string    `gorm:"column:clusterCreator; not null; default:null" json:"clusterCreator"`
// 	State      string    `gorm:"column:clusterState; not null; default:null; DEFAULT:'pending'" json:"clusterState"`
// 	Version    string    `gorm:"column:kubeVersion" json:"kubeVersion"`
// 	Created_at time.Time `gorm:"column:created_at" json:"created_at"`
// 	Token      string    `gorm:"column:token; not null; default:null" json:"token"`
// 	// Monitoring []MONITOR `json:"monitoring"`
// }

type Cluster struct {
	Num      int    `gorm:"column:clusterNum; primary_key" json:"clusterNum"`
	Name     string `gorm:"column:clusterName; not null" json:"clusterName"`
	Ip       string `gorm:"column:ipAddr; not null" json:"ipAddr"`
	Role     string `json:"clusterRole"`
	Type     string `gorm:"column:clusterType; not null" json:"clusterType"`
	Gpu      string `json:"gpu"`
	Endpoint string `gorm:"column:clusterEndpoint; not null" json:"clusterEndpoint"`
	Creator  string `gorm:"column:clusterCreator; not null" json:"clusterCreator"`
	Version  string `json:"kubeVersion"`
	// Token    string `gorm:"column:token; not null" json:"token"`
	// Name       string            `json:"name"`
	Status string `json:"status"`
	// IP         string            `json:"ip"`
	Network    string            `json:"network"`
	Os         string            `json:"os"`
	Kernel     string            `json:"kernel"`
	Lable      map[string]string `json:"lables"`
	Annotation map[string]string `json:"annotations"`
	CreateAt   time.Time         `json:"created_at"`
	// Pod        []Pod    `json:"pods"`
	// Metadata   Metadata `json:"metadata"`
	Events []EVENT `json:"events"`
	// Monitoring []MONITOR `json:"monitoring"`
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

// type Clusters []Cluster
