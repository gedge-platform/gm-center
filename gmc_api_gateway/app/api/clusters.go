package api

import (
	"fmt"
	"net/http"
	"strings"

	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllClusters(c echo.Context) (err error) {
	db := db.DbManager()
	models := []model.Cluster{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
		return
	}
	fmt.Printf("[3##]models : %+v\n", models)
	return c.JSON(http.StatusOK, echo.Map{"data": models})
}
func GetAllDBClusters(params model.PARAMS) []model.Cluster {
	db := db.DbManager()
	models := []model.Cluster{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		// common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
		return nil
	}
	fmt.Printf("[3##]models : %+v\n", models)
	return models
}

// func GetCluster(c echo.Context) (err error) {
// 	db := db.DbManager()
// 	search_val := c.Param("name")
// 	models := FindClusterDB(db, "Name", search_val)

// 	if models == nil {
// 		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
// 		return
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{"data": models})
// }

func GetDBCluster(params model.PARAMS) *model.Cluster {
	search_val := params.Name
	db := db.DbManager()
	// search_val := c.Param("name")
	models := FindClusterDB(db, "Name", search_val)

	if models == nil {
		var model model.Cluster
		model.Name = params.Name
		return &model
	}

	return models
}

func CreateCluster(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.Cluster)

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
	}
	if err = c.Validate(models); err != nil {
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return nil
	}

	if err != nil {
		panic(err)
	}

	if err := db.Create(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return nil
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": models})
}

func UpdateCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := model.Cluster{}

	if err := c.Bind(&models); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := FindClusterDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	} else {
		models.Name = search_val
	}

	models2 := FindClusterDB(db, "Name", search_val)

	if models.Ip != "" {
		models2.Ip = models.Ip
	}
	// if models.extIp != "" { models2.extIp = models.extIp }
	// if models.Role != "" {
	// 	models2.Role = models.Role
	// }
	if models.Type != "" {
		models2.Type = models.Type
	}
	if models.Endpoint != "" {
		models2.Endpoint = models.Endpoint
	}
	if models.Creator != "" {
		models2.Creator = models.Creator
	}
	// if models.Version != "" {
	// 	models2.Version = models.Version
	// }

	if err := db.Save(&models2).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models2})
}

func DeleteCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")

	if err := FindClusterDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, common.ErrNotFound)
		return nil
	}

	models := FindClusterDB(db, "Name", search_val)

	if err := db.Delete(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func FindClusterDB(db *gorm.DB, select_val string, search_val string) *model.Cluster {
	models := model.Cluster{}

	if check := strings.Compare(search_val, "") == 0; check {
		return nil
	}

	if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.Cluster{Name: search_val}).Error; err != nil {
			return nil
		}
	}
	return &models
}
func ClusterResourceCnt(params model.PARAMS, kind string) int {
	fmt.Printf("[##]params : %+v\n", params)
	params.Kind = kind
	params.Project = ""
	params.Name = ""
	deployments, _ := common.DataRequest(params)
	deployment := common.FindingArray(common.Finding(deployments, "items"))
	// for i, _ := range deployment {
	// 	fmt.Printf("[##]names : %s\n", (gjson.Get(deployment[i].String(), "metadata.name")).String())
	// 	fmt.Printf("[##]index : %d\n", i)
	// }
	deployment_cnt := common.FindingLen2(deployment)
	// fmt.Printf("deployment_cnt : %d\n", deployment_cnt)
	return deployment_cnt
}
func GetCluster(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	params.Cluster = params.Name
	params.Name = ""

	//
	// var WorkerList []model.CLUSTER
	// for m, _ := range Master {
	params.Name = params.Cluster
	cluster := GetDBCluster(params)
	if cluster == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}
	var tsCluster model.Cluster
	var clusterModel model.CLUSTER_DETAIL
	common.Transcode(cluster, &tsCluster)
	common.Transcode(tsCluster, &clusterModel)
	fmt.Printf("[##]cluster : %+v\n", clusterModel)
	params.Name = ""
	getData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	Nodes, _ := common.FindDataLabelKey(getData, "items", "labels", "node-role.kubernetes.io/master")
	var NodeList []model.NODE
	fmt.Printf("[##]NODES : %+v\n", Nodes)
	for n, _ := range Nodes {
		Node := model.NODE{
			Name:     common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
			NodeType: common.InterfaceToString(common.FindData(Nodes[n], "nodeType", "")),

			CreateAt:    common.InterfaceToTime(common.FindData(Nodes[n], "metadata", "creationTimestamp")),
			Version:     common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "kubeletVersion")),
			Label:       common.FindData(Nodes[n], "metadata", "labels"),
			Annotation:  common.FindData(Nodes[n], "metadata", "annotations"),
			Allocatable: common.FindData(Nodes[n], "status", "allocatable"),
			IP:          common.InterfaceToString(common.FindData(Nodes[n], "status", "addresses.0.address")),
			// Status                   : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
			// Network                 : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
			Os:                      common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "operatingSystem")) + " / " + common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "osImage")),
			Kernel:                  common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "kernelVersion")),
			ContainerRuntimeVersion: common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "containerRuntimeVersion")),
			Capacity:                common.FindData(Nodes[n], "status", "capacity"),
			// Addresses  : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
		}
		NodeList = append(NodeList, Node)
	}
	ResourceCnt := model.PROJECT_RESOURCE{
		DeploymentCount:  ClusterResourceCnt(params, "deployments"),
		DaemonsetCount:   ClusterResourceCnt(params, "daemonsets"),
		StatefulsetCount: ClusterResourceCnt(params, "statefulsets"),
		PodCount:         ClusterResourceCnt(params, "pods"),
		ServiceCount:     ClusterResourceCnt(params, "services"),
		CronjobCount:     ClusterResourceCnt(params, "cronjobs"),
		JobCount:         ClusterResourceCnt(params, "jobs"),
		VolumeCount:      ClusterResourceCnt(params, "persistentvolumes"),
	}
	params.Name = params.Cluster
	gpuList, check := GpuCheck(params.Name)
	if check != false {
		clusterModel.Gpu = gpuList
	} else {
		clusterModel.Gpu = nil
	}
	clusterModel.Gpu = gpuList
	clusterModel.Resource = ResourceCnt
	clusterModel.Events = getCallEvent(params)
	clusterModel.Nodes = NodeList

	return c.JSON(http.StatusOK, echo.Map{
		"data": clusterModel,
	})
}

func GetClusters(c echo.Context) (err error) {
	var clusterList []model.CLUSTER
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	if c.QueryParam("workspace") == "" {
		clusters := GetAllDBClusters(params)
		if clusters == nil {
			common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
			return nil
		}
		for k, _ := range clusters {
			fmt.Printf("value : %+v\n", clusters[k].Name)
			params.Cluster = clusters[k].Name
			getData, err := common.DataRequest(params)
			if err != nil {
				common.ErrorMsg(c, http.StatusNotFound, err)
				return nil
			}
			fmt.Printf("######getData : %+v\n", getData)
			var clusterModel model.CLUSTER
			common.Transcode(clusters[k], &clusterModel)
			gpuList, check := GpuCheck(params.Name)
			if check != false {
				clusterModel.Gpu = gpuList
			} else {
				clusterModel.Gpu = nil
			}

			clusterModel.NodeCnt = common.InterfaceOfLen(common.FindData(getData, "items", ""))
			tempMetric := []string{"cpu_usage", "memory_usage", "pod_running"}
			tempresult := NowMonit("cluster", params.Cluster, params.Name, tempMetric)
			clusterModel.ResourceUsage = tempresult
			clusterList = append(clusterList, clusterModel)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": clusterList,
		})
	} else {
		params.Workspace = c.QueryParam("workspace")
		workspace := GetDBWorkspace(params)
		if workspace == nil {
			common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
			return
		}
		selectCluster := workspace.SelectCluster
		slice := strings.Split(selectCluster, ",")
		for i, _ := range slice {
			params.Name = slice[i]
			params.Cluster = slice[i]
			// params.Project = slice[i]
			cluster := GetDBCluster(params)
			if cluster == nil {
				common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
				return nil
			}
			var tsCluster model.Cluster
			var clusterModel model.CLUSTER
			common.Transcode(cluster, &tsCluster)
			common.Transcode(tsCluster, &clusterModel)
			params.Name = ""
			getData, err := common.DataRequest(params)
			if err != nil {
				common.ErrorMsg(c, http.StatusNotFound, err)
				return nil
			}
			fmt.Printf("[###data] : %s\n", getData)
			gpuList, check := GpuCheck(params.Name)
			if check != false {
				clusterModel.Gpu = gpuList
			} else {
				clusterModel.Gpu = nil
			}

			clusterModel.NodeCnt = common.InterfaceOfLen(common.FindData(getData, "items", ""))
			tempMetric := []string{"cpu_usage", "memory_usage", "pod_running"}
			tempresult := NowMonit("cluster", params.Cluster, params.Name, tempMetric)
			clusterModel.ResourceUsage = tempresult
			clusterList = append(clusterList, clusterModel)

		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": clusterList,
		})
	}

}
