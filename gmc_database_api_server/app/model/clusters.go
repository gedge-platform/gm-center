package model

import (
	"time"
)

  type Cluster struct {
	Num				int			`gorm:"column:clusterNum; primary_key" json:"clusterNum"`
	Ip				string			`gorm:"column:ipAddr; not null; default:null" json:"ipAddr"`
	extIp				string			`gorm:"column:extIpAddr" json:"extIpAddr"`
	Name 				string 			`gorm:"column:clusterName; not null; default:null" json:"clusterName"`
	Role 				string 			`gorm:"column:clusterRole; not null; default:null" json:"clusterRole"`
	Type				string			`gorm:"column:clusterType; not null; default:null" json:"clusterType"`
	Endpoint				string			`gorm:"column:clusterEndpoint; not null; default:null" json:"clusterEndpoint"`
	Creator				string			`gorm:"column:clusterCreator; not null; default:null" json:"clusterCreator"`
	State				string			`gorm:"column:clusterState; not null; default:null; DEFAULT:'pending'" json:"clusterState"`
	Version				string			`gorm:"column:kubeVersion" json:"kubeVersion"`
	Created_at 			time.Time 		`gorm:"column:created_at" json:"created_at"`
  }

// Set Cluster table name to be `CLUSTER_INFO`
func (Cluster) TableName() string {
	return "CLUSTER_INFO"
  }
