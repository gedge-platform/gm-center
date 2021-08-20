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
	Conditions     []Conditions     `json:"conditions"`
	StartTime      time.Time        `json:"startTime"`
	CompletionTime time.Time        `json:"completionTime"`
	// Cronjob   CRONJOB        `json:"cronjob"`
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
type ReferDataJob struct {
	ReferPodList []ReferPodList `json:"podList"`
	Event        []EVENT        `json:"event"`
}
type ReferPodList struct {
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Status struct {
		Phase  string `json:"phase"`
		HostIP string `json:"hostIP"`
		PodIP  string `json:"podIP"`
	} `json:"status"`
	Spec struct {
		NodeName string `json:"nodeName"`
	} `json:"spec"`
}
