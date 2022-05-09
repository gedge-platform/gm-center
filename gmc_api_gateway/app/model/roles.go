package model

import (
	"time"
)

type ROLE struct {
	Name        string      `json:"name"`
	Namespace   string      `json:"namespace"`
	Lable       interface{} `json:"label,omitempty"`
	Annotations interface{} `json:"annotations,omitempty"`
	Rules       interface{} `json:"rules"`
	CreateAt    time.Time   `json:"createAt"`
}

type ROLEs []ROLEs

func (ROLE) TableName() string {
	return "ROLE_INFO"
}
