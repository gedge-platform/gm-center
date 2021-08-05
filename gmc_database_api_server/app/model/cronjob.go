package model

import "time"

type CRONJOB struct {
	Workspace         string    `json:"workspace"`
	Cluster           string    `json:"cluster"`
	Project           string    `json:"project"`
	Name              string    `json:"name"`
	Kind              string    `json:"kind"`
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
	JobRecords                 JOB               `json:"jobRecords"`
	JOBSTATUS                  string            `json:"jobstatus"`
	ACTIVE                     []ACTIVE          `json:"active"`
	UpdateAt                   time.Time         `json:"updateAt"`
	Events                     []EVENT           `json:"events"`
}

// type SPEC struct {
// 	ConcurrencyPolicy          string `json:"concurrencyPolicy"`
// 	FailedJobsHistoryLimit     int    `json:"failedJobsHistoryLimit"`
// 	Schedule                   string `json:"schedule"`
// 	SuccessfulJobsHistoryLimit int    `json:"successfulJobsHistoryLimit"`
// 	Suspend                    bool   `json:"suspend"`
// }

type ACTIVE struct {
	name      string `json:"name"`
	kind      string `json:"kind"`
	namespace string `json:"namespace"`
}
