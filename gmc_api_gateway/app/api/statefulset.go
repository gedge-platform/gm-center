package api

import (
	"fmt"
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"
	"log"
	"github.com/labstack/echo/v4"
)

func GetStatefulset(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "statefulsets",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	// fmt.Println("[###########ingress]", common.InterfaceToString(common.FindDataStr(getData, "status.loadBalancer.ingress.0", "ip")))
	daemonset := model.WORKLOAD{
		Name:          common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Namespace:     common.InterfaceToString(common.FindData(getData, "metadata", "namespace")),
		// Replica:       replicas,
		ClusterName:   common.InterfaceToString(common.FindData(getData, "clusterName", "")),
		CreateAt:      common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		// UpdateAt:      common.InterfaceToTime(common.FindData(data[i], "status.conditions", "lastUpdateTime")),
		// Stauts:        common.FindData(getData, "status", ""),
		WorkspaceName: common.InterfaceToString(common.FindData(getData, "workspaceName", "")),
		// UpdateAt:        common.InterfaceToTime(common.FindData(getData, "metadata.managedFields.#", "time")),
	}

	involvesData, _ := common.GetModelRelatedList(params) // Pods, Deployments
	log.Printf("#####involvesData ", involvesData)

	return c.JSON(http.StatusOK, echo.Map{
		"data":         daemonset,
		"involvesData": involvesData,
	})
}
func GetAllStatefulset(c echo.Context) (err error) {
	var daemonsets []model.WORKLOAD
	params := model.PARAMS{
		Kind:      "statefulsets",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	data := GetModelList(params)
	fmt.Printf("#################dataerr : %s", data)
	for i, _ := range data {
		daemonset := model.WORKLOAD{
			Name:          common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Namespace:     common.InterfaceToString(common.FindData(data[i], "metadata", "namespace")),
			// Replica:       replicas,
			ClusterName:   common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			CreateAt:      common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
			// UpdateAt:      common.InterfaceToTime(common.FindData(data[i], "status.conditions", "lastUpdateTime")),
			// Stauts:        common.FindData(data[i], "status", ""),
			WorkspaceName: common.InterfaceToString(common.FindData(data[i], "workspaceName", "")),
		}
		daemonsets = append(daemonsets, daemonset)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": daemonsets,
	})
}

// func CreateDeployment(c echo.Context) (err error) {
// 	params := model.PARAMS{
// 		Kind:    "deployments",
// 		Cluster: c.QueryParam("cluster"),
// 		Project: c.QueryParam("project"),
// 		Method:  c.Request().Method,
// 		Body:    responseBody(c.Request().Body),
// 	}

// 	postData, err := common.DataRequest(params)
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, err)
// 		return nil
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"info": common.StringToInterface(postData),
// 	})
// }

// func DeleteDeployment(c echo.Context) (err error) {
// 	params := model.PARAMS{
// 		Kind:    "deployments",
// 		Name:    c.Param("name"),
// 		Cluster: c.QueryParam("cluster"),
// 		Project: c.QueryParam("project"),
// 		Method:  c.Request().Method,
// 		Body:    responseBody(c.Request().Body),
// 	}

// 	postData, err := common.DataRequest(params)
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, err)
// 		return nil
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"info": common.StringToInterface(postData),
// 	})
// }
