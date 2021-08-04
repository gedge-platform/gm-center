package model

import "time"

type POD struct {
	Workspace string    `json:"workspace"`
	Cluster   string    `json:"cluster"`
	Project   string    `json:"project"`
	Name      string    `json:"name"`
	Ready     string    `json:"ready"`
	Status    string    `json:"status"`
	PodIP     string    `json:"podIP"`
	Restart   string    `json:"restart"`
	CreatedAt time.Time `json:"created_at"`
	// Cronjob   CRONJOB        `json:"cronjob"`
}

type PODDETAIL struct {
	Workspace         string                 `json:"workspace"`
	Cluster           string                 `json:"cluster"`
	Project           string                 `json:"project"`
	Name              string                 `json:"name"`
	Namespace         string                 `json:"namespace"`
	Ready             string                 `json:"ready"`
	Status            string                 `json:"status"`
	PodIP             string                 `json:"podIP"`
	Restart           int                    `json:"restart"`
	CreatedAt         time.Time              `json:"created_at"`
	Lable             map[string]interface{} `json:"label"`
	VolumeMounts      VOLUMEMOUNTS           `json:"volumemounts"`
	ContainerStatuses CONTAINERSTATUSES      `json:"containerStatuses"`
	BackoffLimit      int                    `json:"backoffLimit"`
	Completions       int                    `json:"completions"`
	Parallelism       int                    `json:"parallelism"`
}

type VOLUMEMOUNTS struct {
	MountPath string `json:"mountpath"`
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readonly"`
}

type CONTAINERSTATUSES struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	Ready        bool   `json:"ready"`
	RestartCount int    `json:"restartcount"`
	State        string `json:"state"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ReadinessProbe struct {
	FailureThreshold int `json:"failureThreshold"`
	//    httpGet{
	//         Path     string
	//         Port     int
	//         Scheme   string
	//         }
	InitialDelaySeconds int `json:"initialDelaySeconds"`
	PeriodSeconds       int `json:"periodSeconds"`
	SuccessThreshold    int `json:"successThreshold"`
	TimeoutSeconds      int `json:"timeoutSeconds"`
}

type Ports struct {
	ContainerPort string `json:"containerPort"`
	Name          string `json:"name"`
	Protocol      string `json:"protocol"`
}
