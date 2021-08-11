package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/db"
	"gmc_database_api_server/app/model"
	"log"

	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func Get_Cluster(c echo.Context) (err error) {
	// search_val := c.QueryParam("name")
	// clusterName := c.Param("name")
	KubernetesNODE, err := common.GetModel(c, "nodes")
	if common.FilterStr(KubernetesNODE, "status") == "Failure" {
		// errReturn := KubernetesNS.Array()
		errJson := make(map[string]string)
		err_ns := json.Unmarshal([]byte(KubernetesNODE), &errJson)
		if err_ns != nil {
			fmt.Printf("err_ns : %s\n", err_ns)
		}
		// 	common.ErrorMsg(c, http.StatusNotFound, err)
		return c.JSON(http.StatusNotFound, errJson)
	}
	// KubernetesNODE2, err := common.GetModel2(c, "nodes")

	fmt.Printf("[#]KubernetesNODE : %+v\n", KubernetesNODE)
	clusterModel := GetCluster2(c)
	clusterModel.Lable = common.Filter(KubernetesNODE, "metadata.labels")
	clusterModel.Annotation = common.Filter(KubernetesNODE, "metadata.annotations")
	clusterModel.Kernel = common.Finding(KubernetesNODE, "status.nodeInfo", "kernelVersion")
	clusterModel.Os = common.Finding(KubernetesNODE, "status.nodeInfo", "operatingSystem") + " / " + common.Finding(KubernetesNODE, "status.nodeInfo", "osImage")
	clusterModel.Version = common.Finding(KubernetesNODE, "status.nodeInfo", "kubeProxyVersion")
	clusterModel.Gpu = ""
	clusterModel.Network = ""
	createTime := common.FilterStr(KubernetesNODE, "metadata.creationTimestamp")

	timer, err := time.Parse(time.RFC3339, createTime)
	clusterModel.CreateAt = timer

	return c.JSON(http.StatusOK, echo.Map{"data": clusterModel})
}

func Get_http(c echo.Context) (err error) {

	data, err := common.GetModel(c, "namespaces")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	fmt.Println(reflect.TypeOf(data))

	return c.JSON(http.StatusOK, echo.Map{"data": data})
	return

}
func Get_Project(c echo.Context) (err error) {
	search_val := c.Param("name")
	clusterName := c.QueryParam("cluster")
	projectModel := GetProject2(c)
	KubernetesNS, err := common.GetModel(c, "namespaces")
	if common.FilterStr(KubernetesNS, "status.phase") != "Active" {
		// errReturn := KubernetesNS.Array()
		errJson := make(map[string]string)
		err_ns := json.Unmarshal([]byte(KubernetesNS), &errJson)
		if err_ns != nil {
			fmt.Printf("err_ns : %s\n", err_ns)
		}
		// 	common.ErrorMsg(c, http.StatusNotFound, err)
		return c.JSON(http.StatusNotFound, errJson)
	}
	fmt.Printf("[#] status : %+v\n", common.FilterStr(KubernetesNS, "status.phase"))
	fmt.Printf("[#]KubernetesNS : %+v\n", KubernetesNS)
	projectModel.Lable = common.Filter(KubernetesNS, "metadata.labels")
	projectModel.Annotation = common.Filter(KubernetesNS, "metadata.annotations")
	projectModel.ClusterName = clusterName
	// KubernetesNS, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/namespaces/"+search_val, false)
	ResourceCount, _ := HttpRequest2(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/"+search_val+"/application_resource", false)
	var ResourceModel model.PROJECT_RESOURCE
	fmt.Printf("[#####] ResourceCount : %+v\n", ResourceCount)
	ResourceModel.PodCount = StringToInt(gjson.Get(ResourceCount, "count.pod_count").String())
	ResourceModel.DeploymentCount = StringToInt(gjson.Get(ResourceCount, "count.deployment_count").String())
	ResourceModel.JobCount = StringToInt(gjson.Get(ResourceCount, "count.job_count").String())
	ResourceModel.CronjobCount = StringToInt(gjson.Get(ResourceCount, "count.cronjob_count").String())
	ResourceModel.ServiceCount = StringToInt(gjson.Get(ResourceCount, "count.service_count").String())
	projectModel.Resource = ResourceModel
	createTime := common.FilterStr(KubernetesNS, "metadata.creationTimestamp")
	timer, err := time.Parse(time.RFC3339, createTime)
	projectModel.CreateAt = timer
	return c.JSON(http.StatusOK, echo.Map{"data": projectModel})
}

func Get_Clusters(c echo.Context) error {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	params.Name = "cluster3"
	// getData, _ := common.GetModel2(params, "metadata", "resourceVersion")
	// if getData == nil {
	// 	return nil
	// }
	// log.Println("getData is", getData)
	// fmt.Println("[#32] type:", reflect.ValueOf(getData).Type())

	// var testServices model.Service
	// common.Transcode(getData, &testServices)

	// log.Println("Service Model is", testServices)
	// fmt.Println("[#32] type:", reflect.ValueOf(testServices).Type())

	db := db.DbManager()
	models := []model.Cluster{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
		return nil
	}
	for _, value := range models {
		params.Name = value.Name
		params.Cluster = value.Name
		// labels := make(map[string]string)
		label, _ := common.GetModel2(params, "metadata", "labels")
		if label == nil {
			return nil
		}
		// labels := label.(map[string]string)

		fmt.Println("[#32] label:", reflect.ValueOf(label).Type())
		// value.Lable = labels
		// if err4 := json.Unmarshal([]byte(label), &labels); err4 != nil {
		// 	panic(err)
		// }
		// getData, _ := common.GetModel2(params, "metadata", "labels")
		// if getData == nil {
		// 	return nil
		// }
		log.Println("getData is", getData)

	}
	return c.JSON(http.StatusOK, echo.Map{"data": models})
}
