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
	ProjectList   []Workspace_project `json:"projectList,omitempty"`
	Resource      PROJECT_RESOURCE    `json:"resource,omitempty"`
	ResourceUsage Workspace_Usage     `json:"resourceUsage,omitempty"`
	Events        interface{}         `json:"events"`
}

type Workspace_Usage struct {
	Namespace_cpu    float64 `json:"cpu_usage"`
	Namespace_memory float64 `json:"memory_usage"`
}

type Workspace_project struct {
	Name          string    `json:"projectName"`
	SelectCluster string    `json:"selectCluster"`
	CreateAt      time.Time `json:"created_at"`
	Creator       string    `json:"projectCreator"`
	// Resource      PROJECT_RESOURCE `json:"resource,omitempty"`
	// ResourceUsage interface{}      `json:"resourceUsage,omitempty"`
}

// Set Cluster table name to be `CLUSTER_INFO`
func (Workspace) TableName() string {
	return "WORKSPACE_INFO"
}

func WorkspaceName(w Workspace) string {
	return w.Name
}
