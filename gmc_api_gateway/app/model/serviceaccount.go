package model

type SERVICEACCOUNT struct {
	Name      string `json:"name"`
	NameSpace string `json:"namespace"`
	Cluster   string `json:"cluster"`
	// Secrets     string      `json:"secrets"`
	Secrets     interface{} `json:"secrets, omitempty"`
	SecretCnt   int         `json:"secretCnt"`
	Label       interface{} `json:"label,omitempty"`
	Annotations interface{} `json:"annotations, omitempty"`
	CreateAt    string      `json:"createAt"`
}

type SERVICEs []SERVICEs

func (SERVICEACCOUNT) TableName() string {
	return "SERVICEACCOUNT_INFO"
}
