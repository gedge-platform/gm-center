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

	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	getData0 := common.FindData(getData, "", "")                                          // 빈칸으로 둘 시, 전체 조회
	getData1 := common.FindData(getData, "metadata", "name")                              // metadata 의 name 찾기
	getData1Str := common.InterfaceToString(common.FindData(getData, "metadata", "name")) // interface to String 처리

	log.Println("getData0 is", getData0)
	log.Println("getData1 is", getData1)
	fmt.Println("getData1 type:", reflect.ValueOf(getData1).Type())
	log.Println("getData1Str is", getData1Str)
	fmt.Println("getData1Str type:", reflect.ValueOf(getData1Str).Type())

	var testServices model.Service
	common.Transcode(getData0, &testServices) // interface{} -> struct 적용

	log.Println("Service Model is", testServices)
	fmt.Println("[#32] type:", reflect.ValueOf(testServices).Type())

	services := model.SERVICE{
		Name:      common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Workspace: params.Workspace,
		Cluster:   params.Cluster,
		Project:   params.Project,
	}

	getData99, _ := common.GetModelRelatedList(params)

	return c.JSON(http.StatusOK, echo.Map{
		"services": services,
		"lists":    getData99,
	})
}

func GetServices(c echo.Context) (err error) {

	return nil
}
