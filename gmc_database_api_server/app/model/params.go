package model

import "io"

type PARAMS struct {
	Kind      string        `json:"kind,omitempty"`
	Name      string        `json:"name,omitempty"`
	Cluster   string        `json:"cluster,omitempty"`
	Workspace string        `json:"workspace,omitempty"`
	Project   string        `json:"project,omitempty"`
	Uid       string        `json:"uid,omitempty"`
	Compare   string        `json:"compare,omitempty"`
	Method    string        `json:"reqMethod,omitempty"`
	Body      io.ReadCloser `json:"reqBody,omitempty"`
}

type PARAMSAll struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	// Cluster string        `json:"cluster,omitempty"`
	Project string        `json:"project,omitempty"`
	Uid     string        `json:"uid,omitempty"`
	Compare string        `json:"compare,omitempty"`
	Method  string        `json:"reqMethod,omitempty"`
	Body    io.ReadCloser `json:"reqBody,omitempty"`
}

type ClusterAll struct {
	Name string `json:"name,omitempty"`
}
