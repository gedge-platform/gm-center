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

func Get_Deployment(c echo.Context) (err error) {
	namespaceName := c.QueryParam("namespace")
	clusterName := c.QueryParam("cluster")
	workspaceName := c.QueryParam("workspace")
	KubernetesDEPLOY, err := common.GetModel(c, "deployments")
	var DeploymentModel model.DEPLOYMENT
	var STRATEGY model.STRATEGY
	var REPLICA model.REPLICA
	DeploymentModel.Lable = common.Filter(KubernetesDEPLOY, "metadata.labels")
	DeploymentModel.Annotation = common.Filter(KubernetesDEPLOY, "metadata.labels")
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

	return c.JSON(http.StatusOK, echo.Map{"data": DeploymentModel})
}
