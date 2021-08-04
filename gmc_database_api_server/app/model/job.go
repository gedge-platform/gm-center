package model

import "time"

type JOB struct {
	Workspace string    `json:"workspace"`
	Cluster   string    `json:"cluster"`
	Project   string    `json:"project"`
	Name      string    `json:"name"`
	Kind      string    `json:"kind"`
	Status    string    `json:"status"`
	UpdateAt  time.Time `json:"updated_at"`
	// Cronjob   CRONJOB        `json:"cronjob"`
}

type JOBDETAIL struct {
	JOB
	Lable             map[string]string `json:"label"`
	Annotations       map[string]string `json:"annotations"`
	CONTAINERS        []CONTAINERS      `json:"containers"`
	BackoffLimit      int               `json:"backoffLimit"`
	Completions       int               `json:"completions"`
	Parallelism       int               `json:"parallelism"`
	Parent            []OwnerReference  `json:"ownerReferences"`
	CreationTimestamp time.Time         `json:"create_at"`
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
	Status string `json:"status"`
	Type   string `json:"type"`
}
type Evnet struct {
	Reason  string `json:"reson"`
	Message string `json:"type"`
}
