package model

import (
	"time"
)

type DAEMONSET struct {
	Name          string      `json:"name"`
	Namespace     string      `json:"project"`
	ClusterName   string      `json:"cluster"`
	WorkspaceName string      `json:"workspace,omitempty"`
	Stauts        interface{} `json:"status"`
	Strategy      interface{} `json:"strategy,omitempty"`
	Containers    interface{} `json:"containers,omitempty"`
	// Workspace     string      `json:"workspace,omitempty"`
	// PodInfo     []model.Pod     `json:"pods"`
	// ServiceInfo []model.Service `json:"services"`
	Labels     interface{} `json:"label,omitempty"`
	Events     []EVENT     `json:"events"`
	Annotation interface{} `json:"annotations,omitempty"`
	CreateAt   time.Time   `json:"createAt,omitempty"`
	// UpdateAt   time.Time   `json:"updateAt"`
	// Resource   []DEPLOYMENT_RESOURCE `json:"resource"`
	// jwt.StandardClaim
}

type DAEMONSETS []DAEMONSET

func (DAEMONSET) TableName() string {
	return "DAEMONSET_INFO"
}
