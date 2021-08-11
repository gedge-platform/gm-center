package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get_Cluster(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.Param("name"),
		Workspace: c.Param("name"),
		Project:   c.Param("name"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
	var Cluster model.Node
	common.Transcode(getData0, &Cluster)
	clusterModel := GetCluster2(c)
	clusterModel.Label = common.FindData(getData, "metadata", "labels")
	clusterModel.Annotation = common.FindData(getData, "metadata", "annotations")
	clusterModel.CreateAt = common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp"))
	clusterModel.Version = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kubeletVersion"))
	clusterModel.Os = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "operatingSystem"))
	clusterModel.Kernel = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kernelVersion"))
	// common.Transcode(getData0, &clusterModel)
	return c.JSON(http.StatusOK, echo.Map{
		"cluster": clusterModel,
		// "getData":  getData98,
	})
	// return nil
}

func Get_Project(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.Param("name"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	if common.FindData(getData, "status", "") == "Failure" {
		fmt.Printf("error code : %s\n", getData)
		err_data := common.InterfaceToString(getData)
		// errReturn := KubernetesNS.Array()
		errJson := make(map[string]string)
		err_ns := json.Unmarshal([]byte(err_data), &errJson)
		if err_ns != nil {
			fmt.Printf("err_ns : %s\n", err_ns)
		}
		// 	common.ErrorMsg(c, http.StatusNotFound, err)
		// return c.JSON(http.StatusNotFound, errJson)
		return c.JSON(http.StatusNotFound, errJson)
	}
	getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
	var Namespace model.Namespace
	common.Transcode(getData0, &Namespace)
	projectModel := GetProject2(c)
	projectModel.Status = common.InterfaceToString(common.FindData(getData, "status", "phase"))
	projectModel.Label = common.FindData(getData, "metadata", "labels")
	projectModel.Annotation = common.FindData(getData, "metadata", "annotations")
	projectModel.ClusterName = c.QueryParam("cluster")

	return c.JSON(http.StatusOK, echo.Map{
		"project": projectModel,
	})
}

// func ResourceCnt(c echo.Context) (err error) {
// 	// func Get_Deployment(c echo.Context) (err error) {
// 	// namespaceName := c.QueryParam("namespace")
// 	// clusterName := c.QueryParam("cluster")
// 	// workspaceName := c.QueryParam("workspace")
// 	// // KubernetesDEPLOY, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/"+namespaceName+"/deployments/"+search_val, false)
// 	// KubernetesDEPLOY, err := common.GetModel(c, "deployments")
// 	// var DeploymentModel model.DEPLOYMENT
// 	// var STRATEGY model.STRATEGY
// 	// var REPLICA model.REPLICA
// 	// // var CONTAINER model.CONTAINER
// 	// // label := make(map[string]string)
// 	// // annotation := make(map[string]string)
// 	// // fmt.Printf("[#]KubernetesDEPLOY : %+v\n", KubernetesDEPLOY)
// 	// DeploymentModel.Lable = common.Filter(KubernetesDEPLOY, "metadata.labels")
// 	// // DeploymentModel.Lable = labels
// 	// DeploymentModel.Annotation = common.Filter(KubernetesDEPLOY, "metadata.labels")
// 	// //  = annotaitions
// 	// createTime := common.FilterStr(KubernetesDEPLOY, "metadata.creationTimestamp")

// 	// timer, err := time.Parse(time.RFC3339, createTime)
// 	// DeploymentModel.CreateAt = timer
// 	// Name := common.FilterStr(KubernetesDEPLOY, "metadata.name")
// 	// DeploymentModel.Name = Name
// 	// DeploymentModel.Namespace = namespaceName
// 	// DeploymentModel.ClusterName = clusterName
// 	// DeploymentModel.WorkspaceName = workspaceName
// 	// Replicas := common.FilterStr(KubernetesDEPLOY, "status.replicas")
// 	// REPLICA.Replicas = StringToInt(Replicas)
// 	// ReadyReplicas := common.FilterStr(KubernetesDEPLOY, "status.readyReplicas")
// 	// REPLICA.ReadyReplicas = StringToInt(ReadyReplicas)
// 	// UpdatedReplicas := common.FilterStr(KubernetesDEPLOY, "status.updatedReplicas")
// 	// REPLICA.UpdatedReplicas = StringToInt(UpdatedReplicas)
// 	// AvailableReplicas := common.FilterStr(KubernetesDEPLOY, "status.availableReplicas")
// 	// REPLICA.AvailableReplicas = StringToInt(AvailableReplicas)
// 	// UnavailableReplicas := common.FilterStr(KubernetesDEPLOY, "status.unavailableReplicas")
// 	// REPLICA.UnavailableReplicas = StringToInt(UnavailableReplicas)
// 	// DeploymentModel.Replica = REPLICA
// 	// strategyType := common.Finding(KubernetesDEPLOY, "spec.strategy", "type")
// 	// strategyMaxUnavailable := common.Finding(KubernetesDEPLOY, "spec.strategy", "MaxUnavailable")
// 	// strategyMaxSurge := common.Finding(KubernetesDEPLOY, "spec.strategy", "MaxSurge")
// 	// STRATEGY.Type = strategyType
// 	// STRATEGY.MaxUnavailable = strategyMaxUnavailable
// 	// STRATEGY.MaxSurge = strategyMaxSurge
// 	// DeploymentModel.Strategy = STRATEGY
// 	// containers := common.FilterStr(KubernetesDEPLOY, "spec.template.spec.containers")
// 	// containersData := []model.CONTAINER{}
// 	// err4 := json.Unmarshal([]byte(containers), &containersData)
// 	// if err4 != nil {
// 	// 	panic(err4)
// 	// }

// 	// DeploymentModel.Containers = containersData
// 	// updateTime := common.Finding(KubernetesDEPLOY, "status.conditions", "lastUpdateTime")
// 	// timer2, err := time.Parse(time.RFC3339, updateTime)
// 	// DeploymentModel.UpdateAt = timer2
// 	// fmt.Printf("[#]updateTime : %+v\n", updateTime)

// 	// return c.JSON(http.StatusOK, echo.Map{"data": DeploymentModel})
// 	// 	return nil
// }
func Get_Clusters(c echo.Context) (err error) {
	return nil
}
