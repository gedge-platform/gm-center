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
	CoreCloud      CoreCloud
	EdgeCloud      interface{} `json:"edgeInfo"`
}

type CoreCloud struct {
}
type EdgeCloud struct {
}

type SERVICE_DASHBOARD struct {
	WorkspaceCnt   int64       `json:"workspaceCnt"`
	ProjectCnt     int64       `json:"projectCnt"`
	Resource       interface{} `json:"resource"`
	PodCpuTop5     interface{} `json:"podCpuTop5"`
	PodMemTop5     interface{} `json:"podMemTop5"`
	ProjectCpuTop5 interface{} `json:"projectCpuTop5"`
	ProjectMemTop5 interface{} `json:"projectMemTop5"`
}
