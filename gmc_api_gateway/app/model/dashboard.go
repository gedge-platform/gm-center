package model

type TOTAL_DASHBOARD struct {
	ClusterCnt     int64       `json:"clusterCnt"`
	CoreClusterCnt int         `json:"coreClusterCnt"`
	EdgeClusterCnt int         `json:"edgeClusterCnt"`
	WorkspaceCnt   int64       `json:"workspaceCnt"`
	ProjectCnt     int64       `json:"projectCnt"`
	ClusterCpuTop5 interface{} `json:"clusterCpuTop5"`
	PodCpuTop5     interface{} `json:"podCpuTop5"`
	ClusterMemTop5 interface{} `json:"clusterMemTop5"`
	PodMemTop5     interface{} `json:"podMemTop5"`
	CredentialCnt  int     	   `json:"credentialCnt"`
	EdgeCloud      interface{} `json:"edgeInfo"`
}

type CoreCloud struct {
}
type EdgeCloud struct {
}
