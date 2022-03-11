package model

type DAEMONSET struct {
	Workspace   string `json:"workspace,omitempty"`
	Cluster     string `json:"cluster"`
	Name        string `json:"name"`
	Namespace   string `json:"project"`
	Desired     string
	Age         string      `json:"creationTimestamp"`
	Lable       interface{} `json:"label,omitempty"`
	Annotations interface{} `json:"annotations,omitempty"`
	Events      []EVENT     `json:"events"`
}

type DAEMONSETS []DAEMONSET
