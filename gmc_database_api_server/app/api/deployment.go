package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get_Deployment(c echo.Context) (err error) {
	// var ServicePorts []model.PORT
	params := model.PARAMS{
		Kind:      "deployments",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	getData0 := common.FindData(getData, "", "")
	var Deployment model.Deployment
	common.Transcode(getData0, &Deployment)

	// log.Println("Service Model is", Deployment)
	// fmt.Println("[#32] type:", reflect.ValueOf(Deployment).Type())
	// replica := StringToInt(common.InterfaceToString(common.FindData(getData, "spec", "replicas")))
	replicas := model.REPLICA{
		Replicas:            StringToInt(common.InterfaceToString(common.FindData(getData, "status", "replicas"))),
		ReadyReplicas:       StringToInt(common.InterfaceToString(common.FindData(getData, "status", "readyReplicas"))),
		UpdatedReplicas:     StringToInt(common.InterfaceToString(common.FindData(getData, "status", "updatedReplicas"))),
		AvailableReplicas:   StringToInt(common.InterfaceToString(common.FindData(getData, "status", "availableReplicas"))),
		UnavailableReplicas: StringToInt(common.InterfaceToString(common.FindData(getData, "status", "unavailableReplicas"))),
	}

	deployments := model.DEPLOYMENT{
		Name:          common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		WorkspaceName: params.Workspace,
		ClusterName:   params.Cluster,
		Namespace:     params.Project,
		Label:         common.FindData(getData, "metadata", "labels"),
		Annotation:    common.FindData(getData, "metadata", "annotations"),
		CreateAt:      common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		UpdateAt:      common.InterfaceToTime(common.FindData(getData, "status.conditions", "lastUpdateTime")),
		Replica:       replicas,
		Strategy:      common.FindData(getData, "spec", "strategy"),
		Containers:    common.FindData(getData, "spec.template.spec", "containers"),
	}
	testData, _ := common.GetModelRelatedList(params)
	// fmt.Printf("[####]data : %+v\n", testData)
	return c.JSON(http.StatusOK, echo.Map{
		"deployments": deployments,
		"getData":     testData,
	})
}
