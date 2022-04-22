package routes

import (
	c "gmc_api_gateway/app/controller"

	"github.com/labstack/echo/v4"
)

func GEdgeRoute(e *echo.Echo) {

	r := e.Group("/gmcapi/v2")

	r.POST("/members", c.CreateMember)
	r.GET("/members", c.ListMember)
	r.GET("/members/:memberId", c.FindMember)
	r.DELETE("/members/:memberId", c.DeleteMember)
	r.PUT("/members/:memberId", c.UpdateMember)

	r.POST("/workspace", c.CreateWorkspace)
	r.GET("/workspace", c.ListWorkspace)
	r.GET("/workspace/:workspaceName", c.FindWorkspace)
	r.DELETE("/workspace/:workspaceName", c.DeleteWorkspace)
	r.PUT("/workspace/:workspaceName", c.UpdateWorkspace)

	r.POST("/cluster", c.CreateCluster)
	r.GET("/cluster", c.ListCluster)
	r.GET("/cluster/:clusterName", c.FindCluster)
	r.DELETE("/cluster/:clusterName", c.DeleteCluster)
	r.PUT("/cluster/:clusterName", c.UpdateCluster)

	r.POST("/project", c.CreateProject)
	r.GET("/project", c.ListProject)
	r.GET("/project/:projectName", c.FindProject)
	r.DELETE("/project/:projectName", c.DeleteProject)
	r.PUT("/project/:projectName", c.UpdateProject)

	r.POST("/request", c.CreateRequest)
	r.GET("/request", c.ListRequest)
	r.GET("/request/:requestId", c.FindRequest)
	r.DELETE("/request/:requestId", c.DeleteRequest)
	r.PUT("/request/:requestId", c.UpdateRequest)

	r.POST("/credentials", c.CreateCredentials)
	r.GET("/credentials", c.ListCredentials)
	r.GET("/credentials/:credentialsName", c.FindCredentials)
	r.DELETE("/credentials/:credentialsName", c.DeleteCredentials)
	r.PUT("/credentials/:credentialsName", c.UpdateCredentials)

	r2 := e.Group("/gmcapi/v2/spider")
	r2.GET("/credentials", c.GetALLCredential)
	r2.GET("/credentials/:credentialName", c.GetCredential)
	r2.POST("/credentials", c.CreateCredential)
	r2.DELETE("/credentials/:credentialName", c.DeleteCredential)

	r2.GET("/vm", c.GetALLVm)
	r2.GET("/vm/:vmName", c.GetVm)
	r2.POST("/vm", c.CreateVm)
	r2.DELETE("/vm/:vmName", c.DeleteVm)

}
