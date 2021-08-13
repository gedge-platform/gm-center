package model

import "time"

type CRONJOB struct {
	Workspace                  string       `json:"workspace"`
	Cluster                    string       `json:"cluster"`
	Project                    string       `json:"project"`
	Name                       string       `json:"name"`
	Schedule                   string       `json:"schedule"`
	ConcurrencyPolicy          string       `json:"concurrencyPolicy"`
	SuccessfulJobsHistoryLimit int          `json:"successfulJobsHistoryLimit"`
	FailedJobsHistoryLimit     int          `json:"failedJobsHistoryLimit"`
	LastScheduleTime           time.Time    `json:"lastScheduleTime"`
	CreationTimestamp          time.Time    `json:"creationTimestamp"`
	Containers                 []Containers `json:"containers"`
	Active                     []Active     `json:"active"`
	Lable                      interface{}  `json:"label"`
	Annotations                interface{}  `json:"annotations"`
	JOB                        JOB          `json:"job"`
	Events                     []EVENT      `json:"events"`
}

type REFERJOB struct {
	Name string `json:"name"`
}

type Active struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
}
