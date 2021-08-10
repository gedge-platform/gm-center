package model

import "time"

type JOB struct {
	Workspace    string    `json:"workspace"`
	Cluster      string    `json:"cluster"`
	Project      string    `json:"project"`
	Name         string    `json:"name"`
	Kind         string    `json:"kind"`
	Status       string    `json:"status"`
	CreationTime time.Time `json:"created_at"`
	// Cronjob   CRONJOB        `json:"cronjob"`
}

type JOBDETAIL struct {
	JOB
	Lable          map[string]string `json:"label"`
	Annotations    map[string]string `json:"annotations"`
	CONTAINERS     []CONTAINERS      `json:"containers"`
	BackoffLimit   int               `json:"backoffLimit"`
	Completions    int               `json:"completions"`
	Parallelism    int               `json:"parallelism"`
	OwnerReference []OwnerReference  `json:"ownerReferences"`
	// Time              string            `json:"time"`
	CONDITIONS     []CONDITIONS `json:"conditions"`
	StartTime      time.Time    `json:"startTime"`
	CompletionTime time.Time    `json:"completionTime"`
}

// type JOBSTATUS struct {
// 	CompletionTime time.Time `json:"completionTime"`
// 	StartTime      time.Time `json:"startTime"`
// }
type CONTAINERS struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
type CONDITIONS struct {
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	LastProbeTime time.Time `json:"lastProbeTime"`
}
type JOBEvent struct {
	Reason  string `json:"reson"`
	Message string `json:"type"`
}
