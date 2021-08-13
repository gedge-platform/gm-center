package api

import (
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"log"
	"net/http"
	"reflect"

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
	// getData1 := common.FindData(getData, "metadata", "labels")                            // metadata 의 name 찾기
	// getData1Str := common.InterfaceToString(common.FindData(getData, "metadata", "name")) // interface to String 처리

	// log.Println("getData0 is", getData0)
	// log.Println("getData1 is", getData1)
	// fmt.Println("getData1 type:", reflect.ValueOf(getData1).Type())
	// log.Println("getData1Str is", getData1Str)
	// fmt.Println("getData1Str type:", reflect.ValueOf(getData1Str).Type())

	// interface{} -> struct 적용

	log.Println("Service Model is", Deployment)
	fmt.Println("[#32] type:", reflect.ValueOf(Deployment).Type())
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

	return c.JSON(http.StatusOK, echo.Map{
		"deployments": deployments,
		// "getData":  getData98,
	})
}
