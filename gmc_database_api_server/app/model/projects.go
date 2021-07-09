package model

import (
	"time"
)

  type Project struct {
	Num				int			`gorm:"column:projectNum; primary_key" json:"projectNum"`
	Name 				string 			`gorm:"column:projectName; not null; default:null" json:"projectName"`
	Postfix 				string 			`gorm:"column:projectPostfix; not null; default:null" json:"projectPostfix"`
	Description 				string 			`gorm:"column:projectDescription; not null; default:null" json:"projectDescription"`
	Type 				string 			`gorm:"column:projectType; not null; default:null" json:"projectType"`
	Owner				string			`gorm:"column:projectOwner; not null; default:null" json:"projectOwner"`
	Creator				string			`gorm:"column:projectCreator; not null; default:null" json:"projectCreator"`
	Created_at 			time.Time 		`gorm:"column:created_at" json:"created_at"`
	WorkspaceName				string			`gorm:"column:workspaceName; not null; default:null" json:"workspaceName"`
  }

// Set Cluster table name to be `CLUSTER_INFO`
func (Project) TableName() string {
	return "PROJECT_INFO"
  }