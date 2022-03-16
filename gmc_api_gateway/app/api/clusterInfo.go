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
		Nodes,  _ := common.FindDataLabelKey(getData, "items", "labels", "node-role.kubernetes.io/master")
 		for n, _ := range Nodes {
			Node := model.NodeInfo {
				Name : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
				Type : common.InterfaceToString(common.FindData(Nodes[n], "nodeType", "")),
				Ip :  common.InterfaceToString(common.FindData(Nodes[n], "status", "addresses.0.address")),
			}
			NodeInfoList = append(NodeInfoList, Node)
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
