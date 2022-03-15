package model

import "time"

type CLUSTERROLEBINDING struct {
	Name        string      `json:"name"`
	Cluster     string      `json:"cluster"`
	Labels      interface{} `json:"labels,omitempty"`
	Subjects    interface{} `json:"subjects"`
	RoleRef     interface{} `json: "roleRef"`
	Annotations interface{} `json:"annotations, omitempty"`
	CreateAt    time.Time   `json:"createAt"`
}

type CLUSTERROLEBINDINGs []CLUSTERROLEBINDINGs

func (CLUSTERROLEBINDING) TableName() string {
	return "CLUSTERROLEBINDING_INFO"
}
