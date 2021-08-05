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

func GetJobs(c echo.Context) (err error) {
	// db := db.DbManager()
	job_name := c.Param("name")
	workspace_name := c.QueryParam("workspace")
	project_name := c.QueryParam("project")
	cluster_name := c.QueryParam("cluster")

	// fmt.Printf("job_name is %s\n, workspace name is %s\n, project name is %s", job_name, workspace_name, project_name)

	// KubernetesJOB := get(job_name)
	//
	KubernetesJOB, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/cluster3/default/jobs/hello-job", false)
	// https: //g-api-test.innogrid.tech/kube/v1/test1/default/jobs/hello-27133072
	fmt.Printf("[56] KubernetesJOB is %s", KubernetesJOB)
	data, err := common.GetModel(c, "cronjobs")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	fmt.Printf("[34] data is %s", data)

	var jobModel model.JOB

	containerInfo := []model.CONTAINERS{}
	container := model.CONTAINERS{}
	conditionInfo := []model.CONDITIONS{}
	condition := model.CONDITIONS{}
	i := gjson.Get(KubernetesJOB, `status.succeeded`).String()
	//labels
	t := make(map[string]string)
	labels := gjson.Get(KubernetesJOB, `metadata.labels`)
	fmt.Printf("[labels] is %s\n", labels)
	err_labels := json.Unmarshal([]byte(labels.String()), &t)
	if err_labels != nil {
		fmt.Printf("Error : %s\n", err_labels)
	}
	//annotations
	x := make(map[string]string)
	annotations := gjson.Get(KubernetesJOB, `metadata.annotations`)
	fmt.Printf("[annotations] is %s\n", annotations)
	err_annotation := json.Unmarshal([]byte(annotations.String()), &x)
	if err_annotation != nil {
		fmt.Printf("Error : %s\n", err_annotation)
	}
	kind := gjson.Get(KubernetesJOB, `kind`).String()
	backoffLimit := gjson.Get(KubernetesJOB, `spec.backoffLimit`).String()
	completions := gjson.Get(KubernetesJOB, `spec.completions`).String()
	parallelism := gjson.Get(KubernetesJOB, `spec.parallelism`).String()
	creationTimestamp := gjson.Get(KubernetesJOB, `metadata.creationTimestamp`).Time()
	startTime := gjson.Get(KubernetesJOB, `status.startTime`).Time()
	fmt.Printf("[startTimep] is %s\n", startTime)
	completionTime := gjson.Get(KubernetesJOB, `status.completionTime`).Time()
	fmt.Printf("[completionTime] is %s\n", completionTime)
	// time := gjson.Get(KubernetesJOB, `metadata.managedFields`).String()
	// fmt.Printf("[time] is %s\n", time)
	fmt.Printf("[creationTimestamp] is %s\n", creationTimestamp)
	//ownerReferences
	l := []model.OwnerReference{}
	o := model.OwnerReference{}
	k := gjson.Get(KubernetesJOB, `metadata.ownerReferences`).Array()
	fmt.Printf("k is %s\n", k)
	for n, _ := range k {
		fmt.Printf("data is %s\n", k[n])
		err := json.Unmarshal([]byte(k[n].String()), &o)
		if err != nil {
			panic(err)
		}
		l = append(l, o)
	}
	containers := gjson.Get(KubernetesJOB, `spec.template.spec.containers`).Array()
	fmt.Printf("[#contaienr ] is %s\n", containers)
	for n, _ := range containers {
		fmt.Printf("containerdata is %s\n", containers[n])
		err := json.Unmarshal([]byte(containers[n].String()), &container)
		if err != nil {
			panic(err)
		}
		containerInfo = append(containerInfo, container)
	}
	conditions := gjson.Get(KubernetesJOB, `status.conditions`).Array()
	fmt.Printf("[#conditions ] is %s\n", conditions)
	for n, _ := range conditions {
		fmt.Printf("containerdata is %s\n", conditions[n])
		err := json.Unmarshal([]byte(conditions[n].String()), &condition)
		if err != nil {
			panic(err)
		}
		conditionInfo = append(conditionInfo, condition)
	}
	var Str string
	if i == "1" {
		Str = "Running"
	} else {
		Str = "Pending"
	}

	jobModel.Workspace = workspace_name
	jobModel.Project = project_name
	jobModel.Cluster = cluster_name
	jobModel.Name = job_name

	jobModel.UpdateAt = creationTimestamp
	jobModel.Status = Str
	//detail
	var jobDetail model.JOBDETAIL
	jobDetail.Kind = kind
	jobDetail.Workspace = workspace_name
	jobDetail.Project = project_name
	jobDetail.Cluster = cluster_name
	jobDetail.Name = job_name
	jobDetail.Lable = t
	jobDetail.Annotations = x
	jobDetail.Parent = l
	jobDetail.Status = Str
	jobDetail.UpdateAt = creationTimestamp
	jobDetail.CONTAINERS = containerInfo
	jobDetail.StartTime = startTime
	jobDetail.CompletionTime = completionTime
	jobDetail.BackoffLimit = StringToInt(backoffLimit)
	jobDetail.Completions = StringToInt(completions)
	jobDetail.Parallelism = StringToInt(parallelism)
	jobDetail.CONDITIONS = conditionInfo

	// return c.JSON(http.StatusOK, echo.Map{"job": jobModel, "jobDetails": jobDetail})
	return c.JSON(http.StatusOK, echo.Map{"jobDetails": jobDetail})
}

// func ReturnBool(i int) string {

// }
