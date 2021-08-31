package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"log"
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
	getData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	service := model.SERVICE{
		Name:            common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Workspace:       params.Workspace,
		Cluster:         params.Cluster,
		Project:         params.Project,
		Type:            common.InterfaceToString(common.FindData(getData, "spec", "type")),
		ClusterIp:       common.InterfaceToString(common.FindData(getData, "spec", "clusterIP")),
		Selector:        common.FindData(getData, "spec", "selector"),
		Ports:           common.FindData(getData, "spec", "ports"),
		SessionAffinity: common.InterfaceToString(common.FindData(getData, "spec", "type")),
		Events:          getCallEvent(params),
		CreateAt:        common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		// UpdateAt:        common.InterfaceToTime(common.FindData(getData, "metadata.managedFields.#", "time")),
	}

	involvesData, _ := common.GetModelRelatedList(params) // Pods, Deployments
	log.Printf("#####involvesData ", involvesData)

	return c.JSON(http.StatusOK, echo.Map{
		"service":      service,
		"involvesData": involvesData,
	})
}

func GetServices(c echo.Context) (err error) {
	var services []model.SERVICE
	params := model.PARAMS{
		Kind:      "services",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	data := GetModelList(params)
	// fmt.Printf("#################dataerr : %s", data)
	for i, _ := range data {
		service := model.SERVICE{
			Name:      common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Cluster:   common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			Project:   common.InterfaceToString(common.FindData(data[i], "metadata", "namespace")),
			Type:      common.InterfaceToString(common.FindData(data[i], "spec", "type")),
			ClusterIp: common.InterfaceToString(common.FindData(data[i], "spec", "clusterIP")),
			Ports:     common.FindData(data[i], "spec", "ports"),
			CreateAt:  common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
		}
		services = append(services, service)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"services": services,
	})
}
