package model

import "time"

type JOB struct {
	Workspace      string           `json:"workspace"`
	Cluster        string           `json:"cluster"`
	Project        string           `json:"project"`
	Name           string           `json:"name"`
	Kind           string           `json:"kind"`
	Status         int              `json:"status"`
	CreationTime   time.Time        `json:"created_at"`
	Lable          interface{}      `json:"label"`
	Annotations    interface{}      `json:"annotations"`
	Containers     []Containers     `json:"containers"`
	BackoffLimit   int              `json:"backoffLimit"`
	Completions    int              `json:"completions"`
	Parallelism    int              `json:"parallelism"`
	OwnerReference []OwnerReference `json:"ownerReferences"`
	// POD            POD              `json:"referpod"`
	Conditions     []Conditions `json:"conditions"`
	StartTime      time.Time    `json:"startTime"`
	CompletionTime time.Time    `json:"completionTime"`
	EVENT          []EVENT      `json:"events"`
	// Cronjob   CRONJOB        `json:"cronjob"`
}

type JOBDETAIL struct {
	// JOB
	// Lable          map[string]string `json:"label"`
	// Annotations    map[string]string `json:"annotations"`
	// CONTAINERS     []CONTAINERS      `json:"containers"`
	// BackoffLimit   int               `json:"backoffLimit"`
	// Completions    int               `json:"completions"`
	// Parallelism    int               `json:"parallelism"`
	// OwnerReference []OwnerReference  `json:"ownerReferences"`
	// POD            POD               `json:"referpod"`
	// CONDITIONS     []CONDITIONS      `json:"conditions"`
	// StartTime      time.Time         `json:"startTime"`
	// CompletionTime time.Time         `json:"completionTime"`
	// EVENT          []EVENT           `json:"events"`
}

// type JOBSTATUS struct {
// 	CompletionTime time.Time `json:"completionTime"`
// 	StartTime      time.Time `json:"startTime"`
// }
type Containers struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
type Conditions struct {
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	LastProbeTime time.Time `json:"lastProbeTime"`
}
type JOBEvent struct {
	Reason  string `json:"reson"`
	Message string `json:"type"`
}
type PodListA struct {
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	Uid               string    `json:"uid"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	PodIP             string    `json:"podIP"`
	NodeName          string    `json:"node_name"`
	HostIP            string    `json:"hostIP"`
}
