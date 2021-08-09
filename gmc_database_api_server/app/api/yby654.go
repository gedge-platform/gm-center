package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func Get_Cluster(c echo.Context) (err error) {
	// search_val := c.QueryParam("name")
	clusterName := c.QueryParam("cluster")
	KubernetesNODE, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/nodes", false)
	clusterModel := GetCluster2(c)
	k := gjson.Get(KubernetesNODE, "items").Array()
	label := make(map[string]string)
	annotation := make(map[string]string)
	clusterRole := ""
	for t, _ := range k {
		labels := k[t].Get("metadata.labels")
		err_label := json.Unmarshal([]byte(labels.String()), &label)
		if err_label != nil {
			fmt.Printf("Error : %s\n", err_label)
		}
		for key, _ := range label {
			if key == "node-role.kubernetes.io/master" {
				clusterRole = "master"
			}
		}
		clusterModel.Lable = label
		annotations := k[t].Get("metadata.annotations")
		err_annotation := json.Unmarshal([]byte(annotations.String()), &annotation)
		if err_annotation != nil {
			fmt.Printf("Error : %s\n", err_annotation)
		}
		clusterModel.Annotation = annotation
		clusterModel.Name = k[t].Get("metadata.name").String()
		createTime := k[t].Get("metadata.creationTimestamp").Time()
		clusterModel.Version = k[t].Get("status.nodeInfo.kubeletVersion").String()
		clusterModel.Os = k[t].Get("status.nodeInfo.operatingSystem").String() + " / " + k[t].Get("status.nodeInfo.osImage").String()
		clusterModel.Kernel = k[t].Get("status.nodeInfo.kernelVersion").String()
		clusterModel.Role = clusterRole
		clusterModel.CreateAt = createTime
	}

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
	search_val := c.QueryParam("name")
	clusterName := c.QueryParam("cluster")
	KubernetesNS, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/namespaces/"+search_val, false)
	ResourceCount, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/"+search_val+"/application_resource", false)
	projectModel := GetProject2(c)
	var ResourceModel model.PROJECT_RESOURCE
	fmt.Printf("[#]KubernetesNS : %+v\n", KubernetesNS)
	if gjson.Get(KubernetesNS, "status").String() == "Failure" {
		common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
		return
	}
	fmt.Printf("[#]ResourceCount : %+v\n", gjson.Get(ResourceCount, "count.pod_count"))
	label := make(map[string]string)
	annotation := make(map[string]string)
	labels := gjson.Get(KubernetesNS, "metadata.labels")
	err_label := json.Unmarshal([]byte(labels.String()), &label)
	if err_label != nil {
		fmt.Printf("err_label : %s\n", err_label)
	}
	projectModel.Lable = label
	annotations := gjson.Get(KubernetesNS, "metadata.annotations")
	err_annotation := json.Unmarshal([]byte(annotations.String()), &annotation)
	if err_annotation != nil {
		fmt.Printf("err_annotation : %s\n", err_annotation)
	}
	projectModel.Annotation = annotation
	projectModel.Name = gjson.Get(KubernetesNS, "metadata.name").String()
	projectModel.Status = gjson.Get(KubernetesNS, "status.phase").String()
	createTime := gjson.Get(KubernetesNS, "metadata.creationTimestamp").Time()
	projectModel.Created_at = createTime
	projectModel.ClusterName = clusterName

	ResourceModel.PodCount = StringToInt(gjson.Get(ResourceCount, "count.pod_count").String())
	ResourceModel.DeploymentCount = StringToInt(gjson.Get(ResourceCount, "count.deployment_count").String())
	ResourceModel.JobCount = StringToInt(gjson.Get(ResourceCount, "count.job_count").String())
	ResourceModel.CronjobCount = StringToInt(gjson.Get(ResourceCount, "count.cronjob_count").String())
	ResourceModel.ServiceCount = StringToInt(gjson.Get(ResourceCount, "count.service_count").String())
	projectModel.Resource = ResourceModel
	return c.JSON(http.StatusOK, echo.Map{"data": projectModel})
}
func Get_Deployment(c echo.Context) (err error) {
	namespaceName := c.QueryParam("namespace")
	clusterName := c.QueryParam("cluster")
	workspaceName := c.QueryParam("workspace")
	// KubernetesDEPLOY, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+clusterName+"/"+namespaceName+"/deployments/"+search_val, false)
	KubernetesDEPLOY, err := common.GetModel(c, "deployments")
	var DeploymentModel model.DEPLOYMENT
	var STRATEGY model.STRATEGY
	var REPLICA model.REPLICA
	// var CONTAINER model.CONTAINER
	// label := make(map[string]string)
	// annotation := make(map[string]string)
	// fmt.Printf("[#]KubernetesDEPLOY : %+v\n", KubernetesDEPLOY)
	DeploymentModel.Lable = common.Filter(KubernetesDEPLOY, "metadata.labels")
	// DeploymentModel.Lable = labels
	DeploymentModel.Annotation = common.Filter(KubernetesDEPLOY, "metadata.labels")
	//  = annotaitions
	createTime := common.FilterStr(KubernetesDEPLOY, "metadata.creationTimestamp")

	timer, err := time.Parse(time.RFC3339, createTime)
	DeploymentModel.CreateAt = timer
	Name := common.FilterStr(KubernetesDEPLOY, "metadata.name")
	DeploymentModel.Name = Name
	DeploymentModel.Namespace = namespaceName
	DeploymentModel.ClusterName = clusterName
	DeploymentModel.WorkspaceName = workspaceName
	Replicas := common.FilterStr(KubernetesDEPLOY, "status.replicas")
	REPLICA.Replicas = StringToInt(Replicas)
	ReadyReplicas := common.FilterStr(KubernetesDEPLOY, "status.readyReplicas")
	REPLICA.ReadyReplicas = StringToInt(ReadyReplicas)
	UpdatedReplicas := common.FilterStr(KubernetesDEPLOY, "status.updatedReplicas")
	REPLICA.UpdatedReplicas = StringToInt(UpdatedReplicas)
	AvailableReplicas := common.FilterStr(KubernetesDEPLOY, "status.availableReplicas")
	REPLICA.AvailableReplicas = StringToInt(AvailableReplicas)
	UnavailableReplicas := common.FilterStr(KubernetesDEPLOY, "status.unavailableReplicas")
	REPLICA.UnavailableReplicas = StringToInt(UnavailableReplicas)
	DeploymentModel.Replica = REPLICA
	strategyType := common.Finding(KubernetesDEPLOY, "spec.strategy", "type")
	strategyMaxUnavailable := common.Finding(KubernetesDEPLOY, "spec.strategy", "MaxUnavailable")
	strategyMaxSurge := common.Finding(KubernetesDEPLOY, "spec.strategy", "MaxSurge")
	STRATEGY.Type = strategyType
	STRATEGY.MaxUnavailable = strategyMaxUnavailable
	STRATEGY.MaxSurge = strategyMaxSurge
	DeploymentModel.Strategy = STRATEGY
	containers := common.FilterStr(KubernetesDEPLOY, "spec.template.spec.containers")
	containersData := []model.CONTAINER{}
	err4 := json.Unmarshal([]byte(containers), &containersData)
	if err4 != nil {
		panic(err4)
	}

	DeploymentModel.Containers = containersData
	updateTime := common.Finding(KubernetesDEPLOY, "status.conditions", "lastUpdateTime")
	timer2, err := time.Parse(time.RFC3339, updateTime)
	DeploymentModel.UpdateAt = timer2
	fmt.Printf("[#]updateTime : %+v\n", updateTime)

	return c.JSON(http.StatusOK, echo.Map{"data": DeploymentModel})
}
