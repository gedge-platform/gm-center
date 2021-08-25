package model

import "time"

type POD struct {
	Workspace string `json:"workspace"`
	Cluster   string `json:"cluster"`
	Project   string `json:"project"`
	Name      string `json:"name"`
	// Kind              string    `json:"kind"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	// Ready     int    `json:"Ready"`
	Namespace string   `json:"namespace"`
	Status    string   `json:"status"`
	HostIP    string   `json:"hostIP"`
	PodIP     string   `json:"podIP"`
	PodIPs    []PodIPs `json:"podIPs"`
	// CreatedAt         time.Time           `json:"created_at"`
	NodeName       string           `json:"node_name"`
	Lable          interface{}      `json:"label"`
	Annotations    interface{}      `json:"annotations"`
	OwnerReference []OwnerReference `json:"ownerReferences"`
	Podcontainers  []PODCONTAINERS  `json:"Podcontainers"`
	QosClass       string           `json:"qosClass"`
	// VolumeMounts      []VolumeMounts      `json:"volumemounts"`
	ContainerStatuses []ContainerStatuses `json:"containerStatuses"`
}

type PodList1 struct {
	Workspace string `json:"workspace"`
	Cluster   string `json:"cluster"`
	Project   string `json:"project"`
	Name      string `json:"name"`
}
type VolumeMounts struct {
	MountPath string `json:"mountpath"`
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readonly"`
}

type PODCONTAINERS struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	// ReadinessProbe ReadinessProbe `json:"readinessProbe",omitempty`
	// LivenessProbe  LivenessProbe  `json:"livenessProbe",omitempty`
	Ports        []Ports        `json:"ports"`
	Env          []ENV          `json:"env",omitempty`
	VolumeMounts []VolumeMounts `json:"volumemounts"`
}

type ENV struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	ValueFrom ValueFrom `json:"valueFrom"`
}

type ReadinessProbe struct {
	FailureThreshold int `json:"failureThreshold"`
	HTTPGET          struct {
		Path   string `json:"path"`
		Port   int    `json:"port"`
		Scheme string `json:"scheme"`
	}
	TcpSocket struct {
		Port string `json:"port"`
	}
	InitialDelaySeconds int `json:"initialDelaySeconds"`
	PeriodSeconds       int `json:"periodSeconds"`
	SuccessThreshold    int `json:"successThreshold"`
	TimeoutSeconds      int `json:"timeoutSeconds"`
}
type LivenessProbe struct {
	FailureThreshold int `json:"failureThreshold"`
	HTTPGET          struct {
		Path   string `json:"path"`
		Port   int    `json:"port"`
		Scheme string `json:"scheme"`
	}
	TcpSocket struct {
		Port string `json:"port"`
	}
	InitialDelaySeconds int `json:"initialDelaySeconds"`
	PeriodSeconds       int `json:"periodSeconds"`
	SuccessThreshold    int `json:"successThreshold"`
	TimeoutSeconds      int `json:"timeoutSeconds"`
}
type Ports struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

type ContainerStatuses struct {
	ContainerID  string `json:"containerID"`
	Name         string `json:"name"`
	Ready        bool   `json:"ready"`
	RestartCount int    `json:"restartCount"`
	Image        string `json:"image"`
	Started      bool   `json:"started"`
}
type ValueFrom struct {
	FieldRef        FieldRef        `json:"fieldRef" `
	ConfigMapKeyRef ConfigMapKeyRef `json:"configMapKeyRef" `
}
type FieldRef struct {
	ApiVersion string `json:"apiVersion"`
	FieldPath  string `json:"fieldPath"`
}
type ConfigMapKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// type PodReferDeploy struct {
// 	Name       string     `json:"name"`
// 	Namespace  string     `json:"namespace"`
// 	Conditions Conditions `json:"status"`
// }

type PodIPs struct {
	Ip string `json:"ip"`
}
type DeployInfo struct {
	Metadata struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
	} `json:"metadata"`
	Status struct {
		ReadyReplicas   int `json:"readyReplicas"`
		Replicas        int `json:"replicas"`
		UpdatedReplicas int `json:"updatedReplicas"`
	} `json:"status"`
}
type ReferDataDeploy struct {
	DeployInfo []DeployInfo `json:"deployList"`
	Event      []EVENT      `json:"event"`
}
