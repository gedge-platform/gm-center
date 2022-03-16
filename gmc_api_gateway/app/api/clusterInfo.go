package api

import (
"fmt"
	"net/http"
"strings"

	"gmc_api_gateway/app/common"

	"gmc_api_gateway/app/model"


	"github.com/labstack/echo/v4"
)


func GetClusterInfo(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	var ClusterInfoList []model.ClusterInfo
		params.Name = c.QueryParam("project")
		project := GetDBProject(params)
		fmt.Printf("[3##]cluster : %+v\n", project)
		if project == nil {
			common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
			return
		}
		selectCluster := project.SelectCluster
		slice := strings.Split(selectCluster, ",")
		for i, _ := range slice {
			params.Name = slice[i]
			cluster := GetDBCluster(params)
			if cluster == nil {
				common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
				return nil
			}
			var tsCluster model.Cluster
			var clusterModel model.CLUSTER
			common.Transcode(cluster, &tsCluster)
			common.Transcode(tsCluster, &clusterModel)	
			fmt.Printf("[3##]cluster : %+v\n", clusterModel)
			var NodeInfoList []model.NodeInfo
			params.Name = ""
			params.Cluster = slice[i]
		getData, err := common.DataRequest(params)
		if err != nil {
				common.ErrorMsg(c, http.StatusNotFound, err)
				return nil
		}
		fmt.Printf("[3##]getData : %+v\n", getData)
		Nodes,  _ := common.FindDataLabelKey(getData, "items", "labels", "node-role.kubernetes.io/master")
 		fmt.Printf("[3##]Nodes : %+v\n", Nodes)
		for n, _ := range Nodes {
			Node := model.NodeInfo {
				Name : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
				Type : common.InterfaceToString(common.FindData(Nodes[n], "nodeType", "")),
				Ip :  common.InterfaceToString(common.FindData(Nodes[n], "status", "addresses.0.address")),
			}
			NodeInfoList = append(NodeInfoList, Node)
		}

		fmt.Printf("[3##]NodeInfoList : %+v\n", NodeInfoList)
		clusterInfo := model.ClusterInfo {
			ClusterName : clusterModel.Name,
			Type  : clusterModel.Type,
			Nodes : NodeInfoList,
		}

		fmt.Printf("[3##]clusterInfo : %+v\n", clusterInfo)
		ClusterInfoList = append(ClusterInfoList, clusterInfo)
		fmt.Printf("[3##]ClusterInfoList : %+v\n", ClusterInfoList)
	}
	// clusters := GetAllDBClusters(params)
	// if clusters == nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
	// 	return nil
	// }
	// for k, _ := range clusters {
	// 	var NodeInfoList []model.NodeInfo
	// 	// params.Name = clusters[k].Name
	// 	params.Cluster = clusters[k].Name
	// 	getData, err := common.DataRequest(params)
	// 	if err != nil {
	// 			common.ErrorMsg(c, http.StatusNotFound, err)
	// 			return nil
	// 	}
	// 	Nodes,  _ := common.FindDataLabelKey(getData, "items", "labels", "node-role.kubernetes.io/master")
 	// 	for n, _ := range Nodes {
	// 		Node := model.NodeInfo {
	// 			Name : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
	// 			Type : common.InterfaceToString(common.FindData(Nodes[n], "nodeType", "")),
	// 			Ip :  common.InterfaceToString(common.FindData(Nodes[n], "status", "addresses.0.address")),
	// 		}
	// 		NodeInfoList = append(NodeInfoList, Node)
	// 	}

	// 	clusterInfo := model.ClusterInfo {
	// 		ClusterName : clusters[k].Name,
	// 		Type  : clusters[k].Type,
	// 		Nodes : NodeInfoList,
	// 	}

	// 	clusterInfoList = append(clusterInfoList, clusterInfo)
	// }
		
	
	return c.JSON(http.StatusOK, echo.Map{
		"data": ClusterInfoList,
	})
}
