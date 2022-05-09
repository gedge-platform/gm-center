package model

import (
	"time"
)

type STORAGECLASS struct {
	Name                 string      `json:"name"`
	Cluster string      `json:"cluster"`
	ReclaimPolicy        string      `json:"reclaimPolicy"`
	Provisioner          string      `json:"provisioner"`
	VolumeBindingMode    string      `json:"volumeBindingMode"`
	AllowVolumeExpansion string      `json:"allowVolumeExpansion"`
	CreateAt             time.Time   `json:"createAt"`
	Parameters interface{}       `json:"parameters,omitempty"`
	Labels interface{} `json:"labels,omitempty"`
	Annotations          interface{} `json:"annotations,omitempty"`
	//Age                  string      `json:"age"`
}

type STORAGECLASSES []STORAGECLASSES

func (STORAGECLASS) TableName() string {
	return "STORAGECLASS_INFO"
}
