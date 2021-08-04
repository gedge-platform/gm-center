package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/model"
	"net/http"

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
	fmt.Printf("job_name is %s\n, workspace name is %s\n, project name is %s", cronjob_name, workspace_name, project_name)

	// KubernetesJOB := get(job_name)
	//
	KubernetesCronJOB, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/cluster2/cronjobs", false)
	// https: //g-api-test.innogrid.tech/kube/v1/test1/default/jobs/hello-27133072
	// fmt.Printf("[56] data is %s", KubernetesCronJOB)
	itemsmap := gjson.Get(KubernetesCronJOB, "items").Array()
	// fmt.Printf("[55555] data is %s", itemsmap)
	var crondjobDetail model.CRONJOBDETAIL
	for t, _ := range itemsmap {

		containers := itemsmap[t].Get(`spec.spec.template.spec.containers`).Array()
		fmt.Printf("[#contaienr ] is %s\n", containers)
		for n, _ := range containers {
			fmt.Printf("containerdata is %s\n", containers[n])
			err := json.Unmarshal([]byte(containers[n].String()), &container)
			if err != nil {
				panic(err)
			}
			containerInfo = append(containerInfo, container)
		}

		schedule := itemsmap[t].Get(`spec.schedule`).String()
		successfulJobsHistoryLimit := itemsmap[t].Get(`spec.successfulJobsHistoryLimit`).String()
		creationTimestamp := itemsmap[t].Get(`metadata.creationTimestamp`).Time()
		fmt.Printf("[#creationTimestamp ] is %s\n", creationTimestamp)
		crondjobDetail.Workspace = workspace_name
		crondjobDetail.Project = project_name
		crondjobDetail.Cluster = cluster_name
		crondjobDetail.Name = cronjob_name
		crondjobDetail.Schedule = schedule
		crondjobDetail.SuccessfulJobsHistoryLimit = StringToInt(successfulJobsHistoryLimit)
		crondjobDetail.CONTAINERS = containerInfo
		crondjobDetail.CreationTimestamp = creationTimestamp
		// crondjobDetail.LastScheduleTime = successfulJobsHistoryLimit
	}
	return c.JSON(http.StatusOK, echo.Map{"items": crondjobDetail})
}
