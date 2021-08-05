package api

import (
	"fmt"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func GetPods(c echo.Context) (err error) {
	pod_name := c.Param("name")
	workspace_name := c.QueryParam("workspace")
	project_name := c.QueryParam("project")
	cluster_name := c.QueryParam("cluster")
	KubernetesPod, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/cluster3/default/pods", false)
	fmt.Printf("[56] data is %s", KubernetesPod)
	fmt.Printf("cronjob_name is %s\n, workspace name is %s\n, project name is %s", pod_name, workspace_name, project_name, cluster_name)

	// envInfo := []model.CONTAINERSTATUSES{}
	// env := model.CONTAINERSTATUSES{}
	itemsmap := gjson.Get(KubernetesPod, "items").Array()
	// fmt.Printf("[55555] data is %s", itemsmap)
	var podDetail model.PODDETAIL
	for pod, _ := range itemsmap {
		envCall := itemsmap[pod].Get(`spec.dnsPolicy`).String()
		fmt.Printf("[#envCall ] is %s\n", envCall)
		// for n, _ := range envCall {
		// 	fmt.Printf("135envCall is %s\n", envCall[n])
		// 	err := json.Unmarshal([]byte(envCall[n].String()), &env)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	envInfo = append(envInfo, env)
		// }

		podDetail.Name = pod_name
		podDetail.Workspace = workspace_name
		podDetail.Project = project_name
		podDetail.Cluster = cluster_name
		// podDetail.ContainerStatuses = envInfo
		podDetail.Test = envCall
	}
	return c.JSON(http.StatusOK, echo.Map{"items": podDetail})
}
