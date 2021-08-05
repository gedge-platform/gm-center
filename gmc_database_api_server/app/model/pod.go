package model

import "time"

type POD struct {
	Workspace string    `json:"workspace"`
	Cluster   string    `json:"cluster"`
	Project   string    `json:"project"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Ready     string    `json:"ready"`
	Status    string    `json:"status"`
	PodIP     string    `json:"podIP"`
	Restart   string    `json:"restart"`
	CreatedAt time.Time `json:"creationTimestamp"`
	NodeName  string    `json:"node_name"`
	// Cronjob   CRONJOB        `json:"cronjob"`
}

type PODDETAIL struct {
	POD
	Lable             map[string]interface{} `json:"label"`
	OwnerReference    []OwnerReference       `json:"ownerReferences"`
	ContainerStatuses []CONTAINERSTATUSES    `json:"containerStatuses"`
	VolumeMounts      []VOLUMEMOUNTS         `json:"volumemounts"`

	// Env          []ENV   `json:"env"`
	Ports        []PORTS `json:"port"`
	BackoffLimit int     `json:"backoffLimit"`
	Completions  int     `json:"completions"`
	Parallelism  int     `json:"parallelism"`
	Test         string  `json:"test"`
}

type VOLUMEMOUNTS struct {
	MountPath string `json:"mountpath"`
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readonly"`
}

type CONTAINERSTATUSES struct {
	Name string `json:"name"`
	// Image string `json:"image"`
	// Ready bool   `json:"ready"`
	ENV struct {
		Name      string `json:"name"`
		valueFrom string `json:"value"`
	}
	// RestartCount int    `json:"restartcount"`
	// State        string `json:"state"`
}

// type ENV struct {
// 	Name      string `json:"name"`
// 	valueFrom string `json:"value"`
// }

type ReadinessProbe struct {
	FailureThreshold int `json:"failureThreshold"`
	httpGet          struct {
		Path   string
		Port   int
		Scheme string
	}
	InitialDelaySeconds int `json:"initialDelaySeconds"`
	PeriodSeconds       int `json:"periodSeconds"`
	SuccessThreshold    int `json:"successThreshold"`
	TimeoutSeconds      int `json:"timeoutSeconds"`
}

type PORTS struct {
	ContainerPort string `json:"containerPort"`
	Name          string `json:"name"`
	Protocol      string `json:"protocol"`
}
