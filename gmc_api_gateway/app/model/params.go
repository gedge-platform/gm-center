package model

type PARAMS struct {
	Kind   string `json:"kind,omitempty"`
	Name   string `json:"name,omitempty"`
	Method string `json:"reqMethod,omitempty"`
	Body   string `json:"reqBody,omitempty"`
}
