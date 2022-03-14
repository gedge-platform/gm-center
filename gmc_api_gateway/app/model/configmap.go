package model

type CONFIGMAP struct {
	Name        string      `json:"name"`
	NameSpace   string      `json:"namespace"`
	Cluster     string      `json:"cluster"`
	Data        interface{} `json:"data ,omitempty"`
	Annotations interface{} `json:"annotations,omitempty"`
	DataCnt     int         `json:"dataCnt"`
	CreateAt    string      `json:"createAt"`
}

type CONFIGMAPs []CONFIGMAPs

func (CONFIGMAP) TableName() string {
	return "CONFIGMAP_INFO"
}
