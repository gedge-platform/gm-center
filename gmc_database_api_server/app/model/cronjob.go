package model

import "time"

type CRONJOB struct {
	Workspace         string    `json:"workspace"`
	Cluster           string    `json:"cluster"`
	Project           string    `json:"project"`
	Name              string    `json:"name"`
	Schedule          string    `json:"schedule"`
	LastScheduleTime  time.Time `json:"lastScheduleTime"`
	CreationTimestamp time.Time `json:"created_at"`
}

type CRONJOBDETAIL struct {
	CRONJOB
	Name                       string            `json:"name"`
	Lable                      map[string]string `json:"label"`
	Annotations                map[string]string `json:"annotations"`
	CONTAINERS                 []CONTAINERS      `json:"containers"`
	LastScheduleTime           time.Time         `json:"lastScheduleTime"`
	SuccessfulJobsHistoryLimit int               `json:"successfulJobsHistoryLimit"`
	FailedJobsHistoryLimit     int               `json:"failedJobsHistoryLimit"`
	Status                     string            `json:"status"`
	JobRecords                 JOB               `json:"jobRecords"`
	CreateAt                   time.Time         `json:"createAt"`
	UpdateAt                   time.Time         `json:"updateAt"`
}
