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
}

type REFERJOB struct {
	Name string `json:"name"`
}

type Active struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
}
type ReferCronJob struct {
	JOBList []JOBList `json:"jobs"`
	Event   []EVENT   `json:"events"`
}
type JOBList struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Status struct {
		Conditions     []Conditions `json:"conditions"`
		CompletionTime time.Time    `json:"completionTime"`
		StartTime      time.Time    `json:"startTime"`
		Succeeded      int          `json:"succeeded"`
	} `json:"status"`
}
