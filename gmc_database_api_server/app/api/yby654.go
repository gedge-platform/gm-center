package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func Get_Cluster(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.Param("name"),
		Workspace: c.Param("name"),
		Project:   c.Param("name"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
	var Cluster model.Node
	common.Transcode(getData0, &Cluster)
	clusterModel := GetCluster2(c)
	clusterModel.Label = common.FindData(getData, "metadata", "labels")
	clusterModel.Annotation = common.FindData(getData, "metadata", "annotations")
	clusterModel.CreateAt = common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp"))
	clusterModel.Version = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kubeletVersion"))
	clusterModel.Os = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "operatingSystem"))
	clusterModel.Kernel = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kernelVersion"))
	// common.Transcode(getData0, &clusterModel)
	return c.JSON(http.StatusOK, echo.Map{
		"cluster": clusterModel,
		// "getData":  getData98,
	})

	// return nil
}

func Get_Project(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.Param("name"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	if common.FindData(getData, "status", "") == "Failure" {
		fmt.Printf("error code : %s\n", getData)
		err_data := common.InterfaceToString(getData)
		// errReturn := KubernetesNS.Array()
		errJson := make(map[string]string)
		err_ns := json.Unmarshal([]byte(err_data), &errJson)
		if err_ns != nil {
			fmt.Printf("err_ns : %s\n", err_ns)
		}
		// 	common.ErrorMsg(c, http.StatusNotFound, err)
		// return c.JSON(http.StatusNotFound, errJson)
		return c.JSON(http.StatusNotFound, errJson)
	}
	projectModel := GetProject2(c)
	getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
	var Namespace model.Namespace
	common.Transcode(getData0, &Namespace)

	projectModel.Label = common.FindData(getData, "metadata", "labels")
	projectModel.Annotation = common.FindData(getData, "metadata", "annotations")
	params.Workspace = projectModel.WorkspaceName
	projectModel.Status = common.InterfaceToString(common.FindData(getData, "status", "phase"))
	projectModel.ClusterName = c.QueryParam("cluster")

	ResourceCnt := model.PROJECT_RESOURCE{
		DeploymentCount: ResourceCnt(params, "deployments"),
		PodCount:        ResourceCnt(params, "pods"),
		ServiceCount:    ResourceCnt(params, "services"),
		CronjobCount:    ResourceCnt(params, "cronjobs"),
		JobCount:        ResourceCnt(params, "jobs"),
		// VolumeCount:     ResourceCnt(params, "deployments"),
	}

	projectModel.Resource = ResourceCnt
	return c.JSON(http.StatusOK, echo.Map{
		"project": projectModel,
	})
}

func ResourceCnt(params model.PARAMS, kind string) int {
	params.Kind = kind
	params.Name = ""
	deployments, _ := common.GetModel(params)
	deployment := common.FindDataStr(deployments, "items", "")
	deployment_cnt := common.FindingLen(deployment)
	fmt.Printf("deployment_cnt : %d\n", deployment_cnt)
	return deployment_cnt
}

func Get_Clusters(c echo.Context) (err error) {
	clusterModel := GetAllClusters2(c)
	for k, _ := range clusterModel {
		fmt.Printf("value : %+v\n", clusterModel[k].Name)
		params := model.PARAMS{
			Kind:      "nodes",
			Name:      clusterModel[k].Name,
			Cluster:   clusterModel[k].Name,
			Workspace: clusterModel[k].Name,
			Project:   clusterModel[k].Name,
			Method:    c.Request().Method,
			Body:      c.Request().Body,
		}
		// params.Name = value.Name
		getData, err := common.GetModel(params)
		if err != nil {
			common.ErrorMsg(c, http.StatusNotFound, err)
			return nil
		}
		getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
		var Cluster model.Node
		common.Transcode(getData0, &Cluster)
		clusterModel[k].Label = common.FindData(getData, "metadata", "labels")
		clusterModel[k].Annotation = common.FindData(getData, "metadata", "annotations")
		clusterModel[k].CreateAt = common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp"))
		clusterModel[k].Version = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kubeletVersion"))
		clusterModel[k].Os = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "operatingSystem"))
		clusterModel[k].Kernel = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kernelVersion"))
		// clusterModel[k].Kernel = "123"
	}
	return c.JSON(http.StatusOK, echo.Map{
		"clusters": clusterModel,
	})
}
func Get_Projects(c echo.Context) (err error) {
	var Projects model.Projects
	// ProjectModel := model.Project{}
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("cluster"),
		Project:   c.QueryParam("cluster"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	var Namespace model.Namespace
	getData0 := common.FindingArray(common.Finding(getData, "items"))
	// getData0 := common.FindData(getData, "items", "")
	// data234 := common.InterfaceToString(getData0)
	for k, _ := range getData0 {
		params.Name = (gjson.Get(getData0[k].String(), "metadata.name")).String()

		ProjectModel := GetProject3(params)
		fmt.Printf("[######]ProjectModel  : %+v\n", ProjectModel)
		// for key, value := range getData0 {
		Projects = append(Projects, ProjectModel)
	}
	// }
	// fmt.Printf("{###]stirng: %s\n", getData0)

	common.Transcode(getData0, &Namespace)
	return c.JSON(http.StatusOK, echo.Map{
		// "projects": projectModel,
		"data": ProjectModels,
	})
}
