package api

import (
// "fmt"
	"net/http"


	"gmc_api_gateway/app/common"

	"gmc_api_gateway/app/model"


	"github.com/labstack/echo/v4"
)


func GetClusterInfo(c echo.Context) (err error) {
	var clusterInfoList []model.ClusterInfo
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	clusters := GetAllDBClusters(params)
	if clusters == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}
	for k, _ := range clusters {
		var NodeInfoList []model.NodeInfo
		// params.Name = clusters[k].Name
		params.Cluster = clusters[k].Name
		getData, err := common.DataRequest(params)
		if err != nil {
				common.ErrorMsg(c, http.StatusNotFound, err)
				return nil
		}
		Master, Worker, _ := common.FindDataLabelKey(getData, "items", "labels", "node-role.kubernetes.io/master")
		for m, _ := range Master {
			Masternode := model.NodeInfo {
				Name : common.InterfaceToString(common.FindData(Master[m], "metadata", "name")),
				Type : "master",
				Ip :  common.InterfaceToString(common.FindData(Master[m], "status", "addresses.0.address")),
			}
			NodeInfoList = append(NodeInfoList, Masternode)
		}
		for w, _ := range Worker {
			Workernode := model.NodeInfo {
				Name : common.InterfaceToString(common.FindData(Worker[w], "metadata", "name")),
				Type : "worker",
				Ip :  common.InterfaceToString(common.FindData(Worker[w], "status", "addresses.0.address")),
			}
			NodeInfoList = append(NodeInfoList, Workernode)
		}
		clusterInfo := model.ClusterInfo {
			ClusterName : clusters[k].Name,
			Type  : clusters[k].Type,
			Nodes : NodeInfoList,
		}

		clusterInfoList = append(clusterInfoList, clusterInfo)
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"data": clusterInfoList,
	})
}
