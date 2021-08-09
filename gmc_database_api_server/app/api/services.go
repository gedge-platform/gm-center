package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetService(c echo.Context) error {
	// var ServicePorts []model.PORT
	params := model.PARAMS{
		Kind:      "services",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}

	getData, _ := WorkspaceData(c, "services")
	if getData == "" {
		return nil
	}

	services := model.SERVICE{
		Name:      common.Finding(getData, "metadata", "name"),
		Workspace: params.Workspace,
		Cluster:   params.Cluster,
		Project:   params.Project,
	}

	getData2 := common.Filter(getData, "spec.ports")

	fmt.Printf("getData2 is %+v\n", getData2["Protocol"])
	// err4 := json.Unmarshal([]byte(getData2), &ServicePortKube)
	// if err4 != nil {
	// 	fmt.Printf("Error : %s\n", err4)
	// }

	// fmt.Printf("ServicePortKube is %+v\n", ServicePortKube)
	// Service.Ports = ServicePortKube
	// Service.Cluster =

	// fmt.Printf("data type is %s\n", common.Typeof(data))

	// data1, _ := common.FilterStr(data, "metadata.ownerReferences")
	// fmt.Printf("data1 type is %s\n", common.Typeof(data1))

	// log.Println("[#2] data is", data2)
	// log.Println("[#2-1] data is", data2["name"])
	// data3, _ := common.Finding(data, "metadata.ownerReferences", "blockOwnerDeletion")
	// fmt.Printf("data3 type is %s\n", common.Typeof(data3))

	// log.Println("[#3] data is", data3)
	// data4, _ := common.Finding(data, "spec.volumes", "name")
	// fmt.Printf("data4 type is %s\n", common.Typeof(data4))

	// log.Println("[#4] data is", data4)

	return c.JSON(http.StatusOK, echo.Map{
		"services":  services,
		"workspace": getData,
	})
}

func GetServices(c echo.Context) (err error) {

	return nil
}

func GetDeployment(c echo.Context) model.DEPLOYMENT {
	namespaceName := c.QueryParam("namespace")
	clusterName := c.QueryParam("cluster")
	workspaceName := c.QueryParam("workspace")
	// KubernetesDEPLOY, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/"+namespaceName+"/deployments/"+search_val, false)
	KubernetesDEPLOY, err := common.GetModel(c, "deployments")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
	}
	var DeploymentModel model.DEPLOYMENT
	var STRATEGY model.STRATEGY
	var REPLICA model.REPLICA
	// var CONTAINER model.CONTAINER
	// label := make(map[string]string)
	// annotation := make(map[string]string)
	// fmt.Printf("[#]KubernetesDEPLOY : %+v\n", KubernetesDEPLOY)
	DeploymentModel.Lable = common.Filter(KubernetesDEPLOY, "metadata.labels")
	// DeploymentModel.Lable = labels
	DeploymentModel.Annotation = common.Filter(KubernetesDEPLOY, "metadata.labels")
	//  = annotaitions
	createTime := common.FilterStr(KubernetesDEPLOY, "metadata.creationTimestamp")

	timer, err := time.Parse(time.RFC3339, createTime)
	DeploymentModel.CreateAt = timer
	Name := common.FilterStr(KubernetesDEPLOY, "metadata.name")
	DeploymentModel.Name = Name
	DeploymentModel.Namespace = namespaceName
	DeploymentModel.ClusterName = clusterName
	DeploymentModel.WorkspaceName = workspaceName
	Replicas := common.FilterStr(KubernetesDEPLOY, "status.replicas")
	REPLICA.Replicas = StringToInt(Replicas)
	ReadyReplicas := common.FilterStr(KubernetesDEPLOY, "status.readyReplicas")
	REPLICA.ReadyReplicas = StringToInt(ReadyReplicas)
	UpdatedReplicas := common.FilterStr(KubernetesDEPLOY, "status.updatedReplicas")
	REPLICA.UpdatedReplicas = StringToInt(UpdatedReplicas)
	AvailableReplicas := common.FilterStr(KubernetesDEPLOY, "status.availableReplicas")
	REPLICA.AvailableReplicas = StringToInt(AvailableReplicas)
	UnavailableReplicas := common.FilterStr(KubernetesDEPLOY, "status.unavailableReplicas")
	REPLICA.UnavailableReplicas = StringToInt(UnavailableReplicas)
	DeploymentModel.Replica = REPLICA
	strategyType := common.Finding(KubernetesDEPLOY, "spec.strategy", "type")
	strategyMaxUnavailable := common.Finding(KubernetesDEPLOY, "spec.strategy", "MaxUnavailable")
	strategyMaxSurge := common.Finding(KubernetesDEPLOY, "spec.strategy", "MaxSurge")
	STRATEGY.Type = strategyType
	STRATEGY.MaxUnavailable = strategyMaxUnavailable
	STRATEGY.MaxSurge = strategyMaxSurge
	DeploymentModel.Strategy = STRATEGY
	containers := common.FilterStr(KubernetesDEPLOY, "spec.template.spec.containers")
	containersData := []model.CONTAINER{}
	err4 := json.Unmarshal([]byte(containers), &containersData)
	if err4 != nil {
		panic(err4)
	}

	DeploymentModel.Containers = containersData
	updateTime := common.Finding(KubernetesDEPLOY, "status.conditions", "lastUpdateTime")
	timer2, err := time.Parse(time.RFC3339, updateTime)
	DeploymentModel.UpdateAt = timer2
	fmt.Printf("[#]updateTime : %+v\n", updateTime)

	return DeploymentModel
}
