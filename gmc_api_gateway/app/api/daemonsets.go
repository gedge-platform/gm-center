package api

import (
	"fmt"
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDaemonsets(c echo.Context) (err error) {
	// var daemonsets []model.DAEMONSET

	params := model.PARAMS{
		Kind:   "daemonsets",
		Name:   c.Param("name"),
		Method: c.Request().Method,
		Body:   responseBody(c.Request().Body),
	}
	getData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	fmt.Println("[###########daemonsets]", getData)

	// data := GetModelList(params)

	// for i, _ := range data {
	// 	fmt.Printf("#################dataerr : %s", data)
	// 	daemonset := model.DAEMONSET{
	// 		Name:    common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
	// 		Cluster: common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
	// 		// Project:    common.InterfaceToString(common.FindData(data[i], "metadata", "namespace")),
	// 		// Type:       common.InterfaceToString(common.FindData(data[i], "spec", "type")),
	// 		// ClusterIp:  common.InterfaceToString(common.FindData(data[i], "spec", "clusterIP")),
	// 		Workspace: common.InterfaceToString(common.FindData(data[i], "workspaceName", "")),
	// 		// Ports:      common.FindData(data[i], "spec", "ports"),
	// 		// ExternalIp: common.InterfaceToString(common.FindData(data[i], "status.loadBalancer.ingress.0", "ip")),
	// 		// CreateAt:   common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
	// 	}
	// 	daemonsets = append(daemonsets, daemonset)
	// }
	return c.JSON(http.StatusOK, echo.Map{
		"data": "nn",
	})
}
