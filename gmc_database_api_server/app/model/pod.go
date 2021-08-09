package model

type POD struct {
	Workspace string `json:"workspace"`
	Cluster   string `json:"cluster"`
	Project   string `json:"project"`
	Name      string `json:"name"`
	// Ready     int    `json:"Ready"`
	Namespace string `json:"namespace"`
	Status    string `json:"status"`
	PodIP     string `json:"podIP"`
	CreatedAt string `json:"creationTimestamp"`
	NodeName  string `json:"node_name"`
}

type PODDETAIL struct {
	POD
	Lable             map[string]string   `json:"label"`
	Annotations       map[string]string   `json:"annotations"`
	OwnerReference    []OwnerReference    `json:"ownerReferences"`
	Podcontainers     []PODCONTAINERS     `json:"Podcontainers"`
	QosClass          string              `json:"qosClass"`
	ContainerStatuses []ContainerStatuses `json:"containerStatuses"`
}

type VOLUMEMOUNTS struct {
	MountPath string `json:"mountpath"`
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readonly"`
}

type PODCONTAINERS struct {
	Name           string         `json:"name"`
	Image          string         `json:"image"`
	ReadinessProbe ReadinessProbe `json:"readinessProbe"`
	LivenessProbe  LivenessProbe  `json:"livenessProbe"`
	PortsPOD       []PORTPOD      `json:"ports"`
	Env            []ENV          `json:"env"`
	VolumeMounts   []VOLUMEMOUNTS `json:"volumemounts"`
}

type ENV struct {
	Name      string    `json:"name"`
	ValueFrom ValueFrom `json:"valueFrom"`
}

type ReadinessProbe struct {
	FailureThreshold    int     `json:"failureThreshold"`
	HTTPGET             HTTPGET `json:"httpGET"`
	InitialDelaySeconds int     `json:"initialDelaySeconds"`
	PeriodSeconds       int     `json:"periodSeconds"`
	SuccessThreshold    int     `json:"successThreshold"`
	TimeoutSeconds      int     `json:"timeoutSeconds"`
}
type LivenessProbe struct {
	FailureThreshold    int     `json:"failureThreshold"`
	HTTPGET             HTTPGET `json:"httpGET"`
	InitialDelaySeconds int     `json:"initialDelaySeconds"`
	PeriodSeconds       int     `json:"periodSeconds"`
	SuccessThreshold    int     `json:"successThreshold"`
	TimeoutSeconds      int     `json:"timeoutSeconds"`
}
type PORTPOD struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

type HTTPGET struct {
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}
type ContainerStatuses struct {
	Name         string `json:"name"`
	Ready        bool   `json:"ready"`
	RestartCount int    `json:"restartCount"`
	Image        string `json:"image"`
	ContainerID  string `json:"containerID"`
	Started      bool   `json:"started"`
}
type ValueFrom struct {
	FieldRef        FieldRef        `json:"fieldRef"`
	ConfigMapKeyRef ConfigMapKeyRef `json:"configMapKeyRef"`
}
type FieldRef struct {
	ApiVersion string `json:"apiVersion"`
	FieldPath  string `json:"fieldPath"`
}
type ConfigMapKeyRef struct {
	Name string `json:"name"`
	Key  int    `json:"key"`
}
