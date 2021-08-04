package model

import (
	"time"
)

type Project struct {
	Num           int                    `gorm:"column:projectNum; primary_key" json:"projectNum"`
	Name          string                 `gorm:"column:projectName; not null; default:null" json:"projectName"`
	Postfix       string                 `gorm:"column:projectPostfix; not null; default:null" json:"projectPostfix"`
	Description   string                 `gorm:"column:projectDescription; not null; default:null" json:"projectDescription"`
	Type          string                 `gorm:"column:projectType; not null; default:null" json:"projectType"`
	Owner         string                 `gorm:"column:projectOwner; not null; default:null" json:"projectOwner"`
	Creator       string                 `gorm:"column:projectCreator; not null; default:null" json:"projectCreator"`
	Created_at    time.Time              `gorm:"column:created_at" json:"created_at"`
	WorkspaceName string                 `gorm:"column:workspaceName; not null; default:null" json:"workspaceName"`
	Status        string                 `json:"status"`
	Label         map[string]interface{} `json:"labels"`
	CreatedAt     string                 `json:"createAt"`
	Resource      PROJECT_RESOURCE       `json:"resource"`
	Lable         map[string]interface{} `json:"lables"`
	Annotation    map[string]interface{} `json:"annotations"`
	Events        []EVENT                `json:"events"`
}

// type PROJECT_DETAIL struct {
// 	Name        string                 `json:"name"`
// 	Status      string                 `json:"status"`
// 	Label       map[string]interface{} `json:"labels"`
// 	Description string                 `json:"description"`
// 	Creator     string                 `json:"creator"`
// 	Owner       string                 `json:"owner"`
// 	CreatedAt   string                 `json:"createAt"`
// 	Resource    PROJECT_RESOURCE       `json:"resource"`
// 	Lable       map[string]interface{} `json:"lables"`
// 	Annotation  map[string]interface{} `json:"annotations"`
// 	Events      []EVENT                `json:"events"`
// }

type PROJECT_RESOURCE struct {
	DeploymentCount int `json:"deployment_count"`
	PodCount        int `json:"pod_count"`
	ServiceCount    int `json:"service_count"`
	CronjobCount    int `json:"cronjob_count"`
	VolumeCount     int `json:"volume_count"`
}

func (Project) TableName() string {
	return "PROJECT_INFO"
}
