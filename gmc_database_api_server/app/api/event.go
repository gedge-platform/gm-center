package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func getCallEvent(c echo.Context) {
	pod_name := c.Param("name")
	namespace_name := c.QueryParam("namespace")
	workspace_name := c.QueryParam("workspace")
	project_name := c.QueryParam("project")
	cluster_name := c.QueryParam("cluster")
	fmt.Printf("cronjob_name is %s\n, workspace name is %s\n, project name is %s", pod_name, workspace_name, project_name, cluster_name)
	// clusterName := c.QueryParam("cluster")
	podEvent, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/"+cluster_name+"/"+namespace_name+"/pods/"+pod_name+"/events", false)
	fmt.Printf("[#33333 podEvnet", podEvent)

}
