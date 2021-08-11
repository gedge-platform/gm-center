package api

import (
	"bytes"
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

func GetJobs(c echo.Context) (err error) {
	job_name := c.Param("name")
	workspace_name := c.QueryParam("workspace")
	project_name := c.QueryParam("project")
	cluster_name := c.QueryParam("cluster")
	var jobDetail model.JOBDETAIL
	data, err := common.GetModel(c, "jobs")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	fmt.Printf("[#55]job data is info %s", data)

	fmt.Printf("job_name is %s\n, workspace name is %s\n, project name is %s", job_name, workspace_name, project_name, cluster_name)

	status := common.FilterStr(data, `status.succeeded`)
	//labels
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
	kind := common.FilterStr(data, `kind`)
	backoffLimit := common.FilterStr(data, `spec.backoffLimit`)
	completions := common.FilterStr(data, `spec.completions`)
	parallelism := common.FilterStr(data, `spec.parallelism`)
	creationTimestamp := common.FilterStr(data, `metadata.creationTimestamp`)
	create_at, err := time.Parse(time.RFC3339, creationTimestamp)

	//ownerReferences
	ownerReferencesInfo := []model.OwnerReference{}

	ownerReferences := common.FilterStr(data, "metadata.ownerReferences")
	fmt.Printf("data1 type is %s\n", common.Typeof(ownerReferences))
	log.Printf("[#523] t is %s\n", ownerReferences)
	ownerReferences_err := json.Unmarshal([]byte(ownerReferences), &ownerReferencesInfo)
	if err != nil {
		panic(ownerReferences_err)
	}

	log.Printf("[#53] t is %s\n", ownerReferencesInfo)
	containerspec := []model.CONTAINERS{}
	containers := common.FilterStr(data, "spec.template.spec.containers")
	fmt.Printf("data1 type is %s\n", common.Typeof(containers))
	log.Printf("[#523] t is %s\n", containers)

	sepcContainer := json.Unmarshal([]byte(containers), &containerspec)
	if err != nil {
		panic(sepcContainer)
	}
	log.Printf("[#53] t is %s\n", containerspec)
	startTime := common.FilterStr(data, "status.startTime")
	completionTime := common.FilterStr(data, "status.completionTime")
	start_timeformat, err := time.Parse(time.RFC3339, startTime)
	completion_timeformat, err := time.Parse(time.RFC3339, completionTime)
	// containers := gjson.Get(KubernetesJOB, `spec.template.spec.containers`).Array()
	// fmt.Printf("[#contaienr ] is %s\n", containers)
	// for n, _ := range containers {
	// 	fmt.Printf("containerdata is %s\n", containers[n])
	// 	err := json.Unmarshal([]byte(containers[n].String()), &container)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	containerInfo = append(containerInfo, container)
	// }
	conditionInfo := []model.CONDITIONS{}
	conditions := common.FilterStr(data, "status.conditions")
	fmt.Printf("[#conditions ] is %s\n", conditions)
	fmt.Printf("data1 type is %s\n", common.Typeof(conditionInfo))
	log.Printf("[#523] t is %s\n", conditionInfo)

	infoContainer := json.Unmarshal([]byte(conditions), &conditionInfo)
	if err != nil {
		panic(infoContainer)
	}
	log.Printf("[#53] t is %s\n", infoContainer)
	var Str string
	if status == "1" {
		Str = "Running"
	} else {
		Str = "Pending"
	}

	jobDetail.Workspace = workspace_name
	jobDetail.Project = project_name
	jobDetail.Cluster = cluster_name
	jobDetail.Name = job_name
	jobDetail.Kind = kind
	jobDetail.OwnerReference = ownerReferencesInfo
	jobDetail.Lable = label
	jobDetail.Annotations = annotation
	jobDetail.Status = Str
	jobDetail.CreationTime = create_at
	jobDetail.CONTAINERS = containerspec
	jobDetail.StartTime = start_timeformat
	jobDetail.CONDITIONS = conditionInfo
	jobDetail.CompletionTime = completion_timeformat
	jobDetail.BackoffLimit = StringToInt(backoffLimit)
	jobDetail.Completions = StringToInt(completions)
	jobDetail.Parallelism = StringToInt(parallelism)
	test33 := "https://g-api-test.innogrid.tech/kube/v1/" + cluster_name + "/" + project_name + "/jobs/" + job_name + "/list"
	fmt.Println("test33 : ", test33)

	jobReferPod, _ := HttpRequest3(c, "https://g-api-test.innogrid.tech/kube/v1/"+cluster_name+"/"+project_name+"/jobs/"+job_name+"/list", false)

	fmt.Printf("[#7897979]jobReferPod : %+v\n", jobReferPod)
	var jobReferPodDetail model.POD
	// jobmodel := []model.POD{}
	fmt.Printf("[#]podDeploy : %+v\n", jobReferPodDetail)
	gjson.Get(jobReferPod, "kind").String()
	refer_name := gjson.Get(jobReferPod, `metadata.name`).String()
	refer_node := gjson.Get(jobReferPod, `spec.nodeName`).String()
	refer_podIP := gjson.Get(jobReferPod, `status.podIP`).String()
	refer_creationTimestamp := gjson.Get(jobReferPod, `metadata.creationTimestamp`).Time()

	containerStatuses := []model.ContainerStatuses{}
	ContainerStatuses_info := common.FilterStr(jobReferPod, "status.containerStatuses")
	fmt.Printf("###888 envs type is %s\n", ContainerStatuses_info)
	fmt.Printf("###555 envs type is %s\n", common.Typeof(ContainerStatuses_info))
	log.Printf("[#798] t is %s\n", ContainerStatuses_info)

	ContainerStatuses_pod := json.Unmarshal([]byte(ContainerStatuses_info), &containerStatuses)
	if err != nil {
		panic(ContainerStatuses_pod)
	}
	jobDetail.POD = jobReferPodDetail
	jobReferPodDetail.Name = refer_name
	jobReferPodDetail.NodeName = refer_node
	jobReferPodDetail.PodIP = refer_podIP
	jobReferPodDetail.CreatedAt = refer_creationTimestamp
	// jobReferPodDetail.ContainerStatuses = containerStatuses
	jobEvent, _ := HttpRequest3(c, "https://g-api-test.innogrid.tech/kube/v1/"+cluster_name+"/"+project_name+"/jobs/"+job_name+"/events", false)
	fmt.Printf("[123456]event : %+v\n", jobEvent)
	itemsmap := gjson.Get(jobEvent, "items").Array()
	fmt.Printf("[33333#]event items : %+v\n", itemsmap)
	// var event model.EVENT
	eventINFO := []model.EVENT{}
	// for t, _ := range itemsmap {
	// 	action := itemsmap[t].Get(`metadata.name`).String()
	// 	fmt.Printf("[33333#]event items : %+v\n", action)
	// 	reason := itemsmap[t].Get(`reason`).String()
	// 	namespace := itemsmap[t].Get(`metadata.namespace`).String()
	// 	fmt.Printf("[555#]event note : %+v\n", reason)
	// 	note := itemsmap[t].Get(`note`).String()
	// 	fmt.Printf("[9999#]event note : %+v\n", note)
	// 	kind := itemsmap[t].Get(`regarding.kind`).String()
	// 	event.Name = action
	// 	event.Reason = reason
	// 	event.Message = note
	// 	event.Namespace = namespace
	// 	event.Kind = kind

	// }
	jobDetail.EVENT = eventINFO
	return c.JSON(http.StatusOK, echo.Map{
		"jobDetails":  jobDetail,
		"jobReferPod": jobReferPodDetail,
		"events":      eventINFO,
	})
}

func Transcode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
