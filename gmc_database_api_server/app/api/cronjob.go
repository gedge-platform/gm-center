package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func GetCronJobs(c echo.Context) (err error) {

// 	cronjob_name := c.Param("name")
// 	workspace_name := c.QueryParam("workspace")
// 	project_name := c.QueryParam("project")
// 	cluster_name := c.QueryParam("cluster")
// 	var cronjobDetail model.CRONJOBDETAIL
// 	data, err := common.GetModel(c, "cronjobs")
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, err)
// 		return nil
// 	}
// 	fmt.Printf("[#55]cronjob data is info %s", data)

// 	fmt.Printf("cronjob_name is %s\n, workspace name is %s\n, project name is %s", cronjob_name, workspace_name, project_name, cluster_name)
// 	schedule := common.FilterStr(data, `spec.schedule`)
// 	successfulJobsHistoryLimit := common.FilterStr(data, `spec.successfulJobsHistoryLimit`)
// 	creationTimestamp := common.FilterStr(data, `metadata.creationTimestamp`)
// 	creationTimestamp_format, err := time.Parse(time.RFC3339, creationTimestamp)
// 	fmt.Printf("[#creationTimestamp ] is %s\n", creationTimestamp)
// 	concurrencyPolicy := common.FilterStr(data, `spec.concurrencyPolicy`)
// 	failedJobsHistoryLimit := common.FilterStr(data, `spec.failedJobsHistoryLimit`)
// 	status := common.FilterStr(data, `spec.suspend`)
// 	fmt.Printf("[#1234777status] is %s\n", status)
// 	jobstatus := common.FilterStr(data, `status.lastScheduleTime`)
// 	fmt.Printf("[#74111] is %s\n", jobstatus)
// 	now := time.Now()
// 	fmt.Printf("[#9999] is %s\n", now)
// 	//labels
// 	label := make(map[string]string)
// 	labels := common.FilterStr(data, "metadata.labels")
// 	fmt.Printf("[labels] is %s\n", labels)
// 	err_labels := json.Unmarshal([]byte(labels), &label)

// 	if err_labels != nil {
// 		fmt.Printf("Error : %s\n", err_labels)
// 	}
// 	annotation := make(map[string]string)
// 	annotations := common.FilterStr(data, "metadata.annotations")
// 	err_annotation := json.Unmarshal([]byte(annotations), &annotation)
// 	if err_annotation != nil {
// 		fmt.Printf("Error : %s\n", err_annotation)
// 	}
// 	//active
// 	activejob := []model.ACTIVE{}
// 	log.Printf("[#523] t is %s\n", activejob)
// 	activejobs := common.FilterStr(data, `status.active`)
// 	fmt.Printf("[#5231321313131] t is %s\n", activejobs)
// 	infoActive := json.Unmarshal([]byte(activejobs), &activejob)
// 	if err != nil {
// 		panic(infoActive)
// 	}
// 	log.Printf("[#53] t is %s\n", activejob)

// 	var stausCondtion string
// 	if status == "false" {
// 		stausCondtion = "Running"
// 	} else {
// 		stausCondtion = "Pending"
// 	}
// 	fmt.Printf("[#1234777status] is %s\n", status)

// 	containerspec := []model.CONTAINERS{}
// 	containers := common.FilterStr(data, `spec.jobTemplate.spec.template.spec.containers`)
// 	fmt.Printf("data1 type is %s\n", common.Typeof(containers))
// 	log.Printf("[#523] t is %s\n", containers)

// 	sepcContainer := json.Unmarshal([]byte(containers), &containerspec)
// 	if err != nil {
// 		panic(sepcContainer)
// 	}
// 	log.Printf("[#53] t is %s\n", containerspec)

// 	lastScheduleTime := common.FilterStr(data, `status.lastScheduleTime`)
// 	lastScheduleTime_format, err := time.Parse(time.RFC3339, lastScheduleTime)

// 	cronjobDetail.Workspace = workspace_name
// 	cronjobDetail.Project = project_name
// 	cronjobDetail.Cluster = cluster_name
// 	cronjobDetail.Name = cronjob_name
// 	cronjobDetail.Schedule = schedule
// 	cronjobDetail.SuccessfulJobsHistoryLimit = StringToInt(successfulJobsHistoryLimit)
// 	cronjobDetail.ConcurrencyPolicy = concurrencyPolicy
// 	cronjobDetail.FailedJobsHistoryLimit = StringToInt(failedJobsHistoryLimit)
// 	cronjobDetail.ACTIVE = activejob
// 	cronjobDetail.Lable = label
// 	cronjobDetail.CONTAINERS = containerspec
// 	cronjobDetail.Annotations = annotation
// 	cronjobDetail.Status = stausCondtion
// 	cronjobDetail.CreationTimestamp = creationTimestamp_format
// 	cronjobDetail.LastScheduleTime = lastScheduleTime_format

// 	cronjobReferJob, _ := HttpRequest3(c, "https://g-api-test.innogrid.tech/kube/v1/"+cluster_name+"/"+project_name+"/cronjobs/"+cronjob_name+"/list", false)

// 	fmt.Printf("[#7897979]jobReferPod : %+v\n", cronjobReferJob)
// 	itemsmap := common.FilterStr(cronjobReferJob, "items")
// 	fmt.Printf("[33333#]event items : %+v\n", itemsmap)
// 	// var job model.REFERJOB
// 	JobInfo := model.JOB{}
// 	eventstatus := json.Unmarshal([]byte(itemsmap), &JobInfo)
// 	if err != nil {
// 		panic(eventstatus)
// 	}
// 	cronjobDetail.JOB = JobInfo
// 	// fmt.Printf("[#7897979]jobReferPod : %+v\n", job)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"items": cronjobDetail,
// 	})
// }
func GetCronJobs(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:      "cronjobs",
		Name:      "",
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}

	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	cronjobs := model.CRONJOB{

		// Name:      common.GetModel2(params, "metadata", "name"),
		Workspace: params.Workspace,
		Cluster:   params.Cluster,
		Project:   params.Project,
		// Kind:      params.Kind,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"services": cronjobs,
		"getData":  getData,
	})
}
