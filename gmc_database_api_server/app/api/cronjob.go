package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func GetCronJobs(c echo.Context) (err error) {
	cronjob_name := c.Param("name")
	workspace_name := c.QueryParam("workspace")
	project_name := c.QueryParam("project")
	cluster_name := c.QueryParam("cluster")
	containerInfo := []model.CONTAINERS{}
	container := model.CONTAINERS{}

	jobrecord := []model.ACTIVE{}
	jobinfo := model.ACTIVE{}
	label := make(map[string]string)
	annotation := make(map[string]string)
	fmt.Printf("cronjob_name is %s\n, workspace name is %s\n, project name is %s", cronjob_name, workspace_name, project_name)

	// KubernetesJOB := get(job_name)
	//
	KubernetesCronJOB, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/cluster2/cronjobs", false)

	// fmt.Printf("[56] data is %s", KubernetesCronJOB)
	itemsmap := gjson.Get(KubernetesCronJOB, "items").Array()
	// fmt.Printf("[55555] data is %s", itemsmap)
	var crondjobDetail model.CRONJOBDETAIL
	for t, _ := range itemsmap {
		containers := itemsmap[t].Get(`spec.jobTemplate.spec.template.spec.containers`).Array()
		fmt.Printf("[#contaienr ] is %s\n", containers)
		for n, _ := range containers {
			fmt.Printf("containerdata is %s\n", containers[n])
			err := json.Unmarshal([]byte(containers[n].String()), &container)
			if err != nil {
				panic(err)
			}
			containerInfo = append(containerInfo, container)
		}
		labels := itemsmap[t].Get("metadata.labels")
		err_label := json.Unmarshal([]byte(labels.String()), &label)
		if err_label != nil {
			fmt.Printf("Error : %s\n", err_label)
		}

		annotations := itemsmap[t].Get("metadata.annotations")
		err_annotation := json.Unmarshal([]byte(annotations.String()), &annotation)
		if err_annotation != nil {
			fmt.Printf("Error : %s\n", err_annotation)
		}
		schedule := itemsmap[t].Get(`spec.schedule`).String()
		successfulJobsHistoryLimit := itemsmap[t].Get(`spec.successfulJobsHistoryLimit`).String()
		creationTimestamp := itemsmap[t].Get(`metadata.creationTimestamp`).Time()
		fmt.Printf("[#creationTimestamp ] is %s\n", creationTimestamp)
		concurrencyPolicy := itemsmap[t].Get(`spec.concurrencyPolicy`).String()
		failedJobsHistoryLimit := itemsmap[t].Get(`spec.failedJobsHistoryLimit`).String()
		status := itemsmap[t].Get(`spec.suspend`).String()
		fmt.Printf("[#1234777status] is %s\n", status)
		jobstatus := itemsmap[t].Get(`status.lastScheduleTime`).Time()
		fmt.Printf("[#74111] is %s\n", jobstatus)
		now := time.Now()
		fmt.Printf("[#9999] is %s\n", now)
		jobinfos := itemsmap[t].Get(`status.active`).Array()
		// fmt.Printf("[#123456jobinfo ] is %s\n", jobinfos)
		for n, _ := range jobinfos {
			fmt.Printf("jobinfo is %s\n", jobinfos[n])
			err := json.Unmarshal([]byte(jobinfos[n].String()), &jobinfo)
			if err != nil {
				panic(err)
			}
			jobrecord = append(jobrecord, jobinfo)
		}
		sepcAll := itemsmap[t].Get(`spec`).String()
		fmt.Printf("[#1sepcAll ] is %s\n", sepcAll)
		var statusJob string
		if jobstatus != now {
			statusJob = "Completed"
		} else {
			statusJob = "Running"
		}
		var stausCondtion string
		if status == "false" {
			stausCondtion = "Running"
		} else {
			stausCondtion = "Pending"
		}

		crondjobDetail.Name = cronjob_name
		crondjobDetail.Workspace = workspace_name
		crondjobDetail.Project = project_name
		crondjobDetail.Cluster = cluster_name
		crondjobDetail.Lable = label
		crondjobDetail.Annotations = annotation
		crondjobDetail.ConcurrencyPolicy = concurrencyPolicy
		crondjobDetail.FailedJobsHistoryLimit = StringToInt(failedJobsHistoryLimit)
		crondjobDetail.Schedule = schedule
		crondjobDetail.SuccessfulJobsHistoryLimit = StringToInt(successfulJobsHistoryLimit)
		crondjobDetail.Status = stausCondtion
		crondjobDetail.JOBSTATUS = statusJob
		crondjobDetail.ACTIVE = jobrecord
		// crondjobDetail.SPEC = sepcAll
		crondjobDetail.CONTAINERS = containerInfo
		crondjobDetail.CreationTimestamp = creationTimestamp

		// crondjobDetail.LastScheduleTime = successfulJobsHistoryLimit
	}
	return c.JSON(http.StatusOK, echo.Map{"items": crondjobDetail})
}
