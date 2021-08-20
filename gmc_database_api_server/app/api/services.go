package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetService(c echo.Context) error {

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

	services := model.SERVICE{
		Name:            common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Workspace:       params.Workspace,
		Cluster:         params.Cluster,
		Project:         params.Project,
		Type:            common.InterfaceToString(common.FindData(getData, "spec", "type")),
		ClusterIp:       common.InterfaceToString(common.FindData(getData, "spec", "clusterIP")),
		Selector:        common.FindData(getData, "spec", "selector"),
		Ports:           common.FindData(getData, "spec", "ports"),
		SessionAffinity: common.InterfaceToString(common.FindData(getData, "spec", "type")),
		CreateAt:        common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		// UpdateAt:        common.InterfaceToTime(common.FindData(getData, "metadata.managedFields.#", "time")),
	}

	involveData, _ := common.GetModelRelatedList(params) // Pods, Deployments

	return c.JSON(http.StatusOK, echo.Map{
		"services": services,
		"involves": involveData,
	})
}

func GetServices(c echo.Context) (err error) {

	return nil
}
