package model

import "time"

type CRONJOB struct {
	Workspace string `json:"workspace"`
	Cluster   string `json:"cluster"`
	Project   string `json:"project"`
	Name      string `json:"name"`
	// Kind              string    `json:"kind"`
	Schedule          string    `json:"schedule"`
	CreationTimestamp time.Time `json:"created_at"`
	LastScheduleTime  time.Time `json:"lastScheduleTime"`
}

type CRONJOBDETAIL struct {
	CRONJOB
	Lable                      map[string]string `json:"label"`
	Annotations                map[string]string `json:"annotations"`
	CONTAINERS                 []CONTAINERS      `json:"containers"`
	ConcurrencyPolicy          string            `json:"concurrencyPolicy"`
	FailedJobsHistoryLimit     int               `json:"failedJobsHistoryLimit"`
	SuccessfulJobsHistoryLimit int               `json:"successfulJobsHistoryLimit"`
	Status                     string            `json:"status"`
	LastScheduleTime           time.Time         `json:"lastScheduleTime"`
	// REFERJOB                   REFERJOB          `json:"job"`
	JOB JOB `json:"job"`
	// JOBSTATUS                  string            `json:"jobstatus"`
	ACTIVE   []ACTIVE  `json:"active"`
	UpdateAt time.Time `json:"updateAt"`
	Events   []EVENT   `json:"events"`
}

type REFERJOB struct {
	Name string `json:"name"`
}

type ACTIVE struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
}
