package api

var messageError struct {
	Errors messageFormat `json:"errors"`
}

type messageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
}

var kubernetes_api="https://g-api.innogrid.cf/gmcapi/v1/"