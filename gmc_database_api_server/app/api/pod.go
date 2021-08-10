package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func GetPods(c echo.Context) (err error) {
	pod_name := c.Param("name")
	workspace_name := c.QueryParam("workspace")
	project_name := c.QueryParam("project")
	cluster_name := c.QueryParam("cluster")
	fmt.Printf("GETpods  is %s\n, workspace name is %s\n, project name is %s", pod_name, workspace_name, project_name, cluster_name)
	var podDetail model.PODDETAIL
	data, err := common.GetModel(c, "pods")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	fmt.Printf("[#55555]data is info %s", data)
	pod_uid := common.FilterStr(data, "metadata.uid")
	fmt.Printf("[#666]pod_uid type is %s\n", pod_uid)

	name := common.Finding(data, "metadata", "name")
	kind := common.FilterStr(data, "kind")
	// data1, _ := common.FilterStr(data, "metadata.ownerReferences")
	fmt.Printf("kind is info %s", kind)

	namespace := common.Finding(data, "metadata", "namespace")
	startTime := common.FilterStr(data, "metadata.creationTimestamp")
	status := common.Finding(data, "status", "phase")

	timeformat, err := time.Parse(time.RFC3339, startTime)

	fmt.Println("[#100 ] time data format is\n", timeformat)

	qosClass := common.Finding(data, "status", "qosClass")

	podIP := common.Finding(data, "status", "podIP")
	nodeName := common.Finding(data, "spec", "nodeName")
	log.Println("[#4] data is", name)
	ownerReferencesInfo := []model.OwnerReference{}
	ownerReferences := common.FilterStr(data, "metadata.ownerReferences")
	fmt.Printf("data1 type is %s\n", common.Typeof(ownerReferences))
	log.Printf("[#523] t is %s\n", ownerReferences)
	ownerReferences_err := json.Unmarshal([]byte(ownerReferences), &ownerReferencesInfo)
	if err != nil {
		panic(ownerReferences_err)
	}

	log.Printf("[#53] t is %s\n", ownerReferencesInfo)
	containerspec := []model.PODCONTAINERS{}
	containers := common.FilterStr(data, "spec.containers")
	fmt.Printf("data1 type is %s\n", common.Typeof(containers))
	log.Printf("[#523] t is %s\n", containers)

	sepcContainer := json.Unmarshal([]byte(containers), &containerspec)
	if err != nil {
		panic(sepcContainer)
	}
	log.Printf("[#53] t is %s\n", containerspec)

	containerStatuses := []model.ContainerStatuses{}
	ContainerStatuses_info := common.FilterStr(data, "status.containerStatuses")
	fmt.Printf("###888 envs type is %s\n", ContainerStatuses_info)
	fmt.Printf("###555 envs type is %s\n", common.Typeof(ContainerStatuses_info))
	log.Printf("[#798] t is %s\n", ContainerStatuses_info)

	ContainerStatuses_pod := json.Unmarshal([]byte(ContainerStatuses_info), &containerStatuses)
	if err != nil {
		panic(ContainerStatuses_pod)
	}

	label := make(map[string]string)
	labels := common.FilterStr(data, "metadata.labels")
	fmt.Printf("[labels] is %s\n", labels)
	err_labels := json.Unmarshal([]byte(labels), &label)

	if err_labels != nil {
		fmt.Printf("Error : %s\n", err_labels)
	}
	annotation := make(map[string]string)
	annotations := common.FilterStr(data, "metadata.annotations")
	err_annotation := json.Unmarshal([]byte(annotations), &annotation)
	if err_annotation != nil {
		fmt.Printf("Error : %s\n", err_annotation)
	}

	podDetail.Name = pod_name
	podDetail.Workspace = workspace_name
	podDetail.Project = project_name
	podDetail.Cluster = cluster_name
	podDetail.Namespace = namespace
	podDetail.Status = status
	podDetail.CreatedAt = timeformat
	podDetail.NodeName = nodeName
	podDetail.PodIP = podIP
	podDetail.QosClass = qosClass
	podDetail.Podcontainers = containerspec
	podDetail.ContainerStatuses = containerStatuses
	podDetail.OwnerReference = ownerReferencesInfo
	podDetail.Lable = label
	podDetail.Annotations = annotation
	podDetail.Kind = kind

	namespace_name := c.QueryParam("namespace")

	podDeploy, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+cluster_name+"/"+namespace_name+"/pods/"+pod_name+"/list", false)

	var referDeploy model.DEPLOYMENT
	fmt.Printf("[#]podDeploy : %+v\n", podDeploy)
	gjson.Get(podDeploy, "kind").String()
	fmt.Printf("[#033333 pod refer", podDeploy)
	fmt.Printf("#122334", gjson.Get(podDeploy, "kind").String())
	refer_name := gjson.Get(podDeploy, `metadata.name`).String()
	fmt.Printf("#12metaname", gjson.Get(podDeploy, `metadata.name`).String())
	refer_namespace := gjson.Get(podDeploy, `metadata.namespace`).String()
	refer_status := gjson.Get(podDeploy, `status.availableReplicas`).String()
	referDeploy.Name = refer_name
	referDeploy.Namespace = refer_namespace
	referDeploy.Stauts = refer_status
	referDeploy.ClusterName = cluster_name
	referDeploy.WorkspaceName = workspace_name

	podEvent, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+cluster_name+"/"+namespace_name+"/pods/"+pod_name+"/events", false)
	fmt.Printf("[44444444888#]event : %+v\n", podEvent)
	itemsmap := gjson.Get(podEvent, "items").Array()
	fmt.Printf("[33333#]event items : %+v\n", itemsmap)
	var event model.EVENT
	for t, _ := range itemsmap {
		action := itemsmap[t].Get(`metadata.name`).String()
		fmt.Printf("[33333#]event items : %+v\n", action)
		reason := itemsmap[t].Get(`reason`).String()
		namespace := itemsmap[t].Get(`metadata.namespace`).String()
		fmt.Printf("[555#]event note : %+v\n", reason)
		note := itemsmap[t].Get(`note`).String()
		fmt.Printf("[9999#]event note : %+v\n", note)
		kind := itemsmap[t].Get(`regarding.kind`).String()
		event.Name = action
		event.Reason = reason
		event.Message = note
		event.Namespace = namespace
		event.Kind = kind

	}

	return c.JSON(http.StatusOK, echo.Map{
		"items":           podDetail,
		"referdeployment": referDeploy,
		"events":          event,
	})
}
