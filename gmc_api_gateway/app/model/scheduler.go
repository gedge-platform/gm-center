package model

import (

)

type Scheduler struct {
	CallbackUrl         string           `json:"callbackUrl"`
	requestId  string  `json:"requestId"`
	Yaml  string  `json:"yaml "`
}

func (Scheduler) TableName() string {
	return "Scheduler_INFO"
}