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

	getData, _ := common.GetModel2(params, "metadata.managedFields", "fieldsType")
	if getData == nil {
		return nil
	}

	log.Println("getData is", getData)
	fmt.Println("[#32] type:", reflect.ValueOf(getData).Type())

	var testServices model.Service
	common.Transcode(getData, &testServices)

	log.Println("Service Model is", testServices)
	fmt.Println("[#32] type:", reflect.ValueOf(testServices).Type())

	services := model.SERVICE{
		// Name:      common.Finding(getData, "metadata", "name"),
		Workspace: params.Workspace,
		Cluster:   params.Cluster,
		Project:   params.Project,
	}

	// fmt.Printf("getData2 is %+v\n", getData2["Protocol"])
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
		"services": services,
		"getData":  getData,
	})
}

func GetServices(c echo.Context) (err error) {

	return nil
}
