package model

import "io"

type PARAMS struct {
	Kind      string        `json:"kind,omitempty"`
	Name      string        `json:"name,omitempty"`
	Cluster   string        `json:"cluster,omitempty"`
	Workspace string        `json:"workspace,omitempty"`
	Project   string        `json:"project,omitempty"`
	Uuid      string        `json:"uuid,omitempty"`
	Compare   string        `json:"compare,omitempty"`
	Method    string        `json:"reqMethod,omitempty"`
	Body      io.ReadCloser `json:"reqBody,omitempty"`
}