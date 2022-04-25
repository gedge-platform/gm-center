package model

type PARAMS struct {
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	Cluster   string `json:"cluster,omitempty"`
	Workspace string `json:"workspace,omitempty"`
	Project   string `json:"project,omitempty"`
	Uid       string `json:"uid,omitempty"`
	Compare   string `json:"compare,omitempty"`
	Method    string `json:"reqMethod,omitempty"`
	Body      string `json:"reqBody,omitempty"`
}

type PARAM struct {
	CredentialName   string `json:"CredentialName,omitempty"`
	DomainName       string `json:"DomainName,omitempty"`
	IdentityEndPoint string `json:"IdentityEndPoint,omitempty"`
	Password         string `json:"Password,omitempty"`
	ProjectID        string `json:"ProjectID,omitempty"`
	Username         string `json:"Username,omitempty"`
	KeyValueInfoList string `json:"KeyValueInfoList"`
	Method           string `json:"reqMethod,omitempty"`
	// Body             string `json:"reqBody,omitempty"`
}

type VMPARAM struct {
	ConnectionName string `json:"ConnectionName,omitempty"`
	Method         string `json:"reqMethod,omitempty"`
}

type CPARAM struct {
	Name         string `json:"name,omitempty"`
	Type         string `json:"type,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Region       string `json:"region,omitempty"`
	Url          string `json:"url,omitempty"`
	Tenant       string `json:"tenant,omitempty"`
	Access_id    string `json:"access_id,omitempty"`
	Access_token string `json:"access_token,omitempty"`
	Project      string `json:"project,omitempty"`
	// Method       string `json:"reqMethod,omitempty"`
}
