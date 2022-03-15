package model

import (
	"time"
)

type STORAGECLASS struct {
	Name                 string      `json:"name"`
	ResourceVersion      string      `json:"resourceVersion "`
	ReclaimPolicy        string      `json:"reclaimPolicy",`
	Provisioner          string      `json:"provisioner",`
	VolumeBindingMode    string      `json:"volumeBindingMode",`
	AllowVolumeExpansion string      `json:"allowVolumeExpansion",`
	CreateAt             time.Time   `json:"createAt"`
	Annotations          interface{} `json:"annotations,omitempty"`
	//Age                  string      `json:"age"`
}

type STORAGECLASSES []STORAGECLASSES

func (STORAGECLASS) TableName() string {
	return "STORAGECLASS_INFO"
}
