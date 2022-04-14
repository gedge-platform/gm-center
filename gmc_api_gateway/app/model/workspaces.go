package model

import (
	"time"
)

type Workspace struct {
	Num           int       `gorm:"column:workspaceNum; primary_key" json:"workspaceNum"`
	Name          string    `gorm:"column:workspaceName; not null" json:"workspaceName"`
	Description   string    `gorm:"column:workspaceDescription; not null" json:"workspaceDescription"`
	SelectCluster string    `gorm:"column:selectCluster; not null" json:"selectCluster"` // DB 에서 Cluster 목록만 출력
	Owner         string    `gorm:"column:workspaceOwner; not null" json:"workspaceOwner"`
	Creator       string    `gorm:"column:workspaceCreator; not null" json:"workspaceCreator"`
	Created_at    time.Time `gorm:"column:created_at" json:"created_at"`
}
type Workspace_detail struct {
	Workspace
	Project []workspace_project `json:"projects,omitempty"`

	Events []EVENT `json:"events"`
}
type workspace_project struct {
	Name          string           `json:"projectName"`
	SelectCluster string           `json:"selectCluster"`
	Resource      PROJECT_RESOURCE `json:"resource,omitempty"`
	ResourceUsage interface{}      `json:"resourceUsage,omitempty"`
}

// Set Cluster table name to be `CLUSTER_INFO`
func (Workspace) TableName() string {
	return "WORKSPACE_INFO"
}

func WorkspaceName(w Workspace) string {
	return w.Name
}
