package routes

import (
	// "fmt"
	// "os"

	"gmc_api_gateway/app/api"

	// "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type DataValidator struct {
	validator *validator.Validate
}

func NewValidator() *DataValidator {
	return &DataValidator{
		validator: validator.New(),
	}
}

func (dv *DataValidator) Validate(i interface{}) error {
	return dv.validator.Struct(i)
}

func GEdgeRoute(e *echo.Echo) {
	e.Validator = NewValidator()

	// /gmcapi/v1
	r := e.Group("/gmcapi/v1")
	r.GET("/members", api.GetAllMembers)
	r.POST("/members", api.CreateMember)
	r.GET("/members/:id", api.GetMember)
	r.PUT("/members/:id", api.UpdateMember)
	r.DELETE("/members/:id", api.DeleteMember)

	r.GET("/apps", api.GetAllApps)
	r.POST("/apps", api.CreateApp)
	r.GET("/apps/:name", api.GetApp)
	r.PUT("/apps/:name", api.UpdateApp)
	r.DELETE("/apps/:name", api.DeleteApp)

	r.GET("/clusters", api.GetClusters)
	// r.POST("/clusters", api.Create_Cluster)
	r.GET("/clusters/:name", api.GetCluster)
	// r.PUT("/clusters/:name", api.Update_Cluster)
	// r.DELETE("/clusters/:name", api.Delete_Cluster)

	r.GET("/projects", api.GetProjects)
	r.GET("/projects/:name", api.GetProject)
	r.POST("/projects", api.CreateProject)
	// r.PUT("/projects/:name", api.ReplaceProject)
	// r.PATCH("/projects/:name", api.UpdateProject)
	r.DELETE("/projects/:name", api.DeleteProject)

	r.GET("/deployments", api.GetDeployments)
	r.POST("/deployments", api.CreateDeployment)
	r.GET("/deployments/:name", api.GetDeployment)
	// r.PUT("/deployments/:name", api.Update_Deployment)
	// r.DELETE("/deployments/:name", api.Delete_Deployment)

	r.GET("/workspaces", api.GetAllWorkspaces)
	r.POST("/workspaces", api.CreateWorkspace)
	r.GET("/workspaces/:name", api.GetWorkspace)
	r.PUT("/workspaces/:name", api.UpdateWorkspace)
	r.DELETE("/workspaces/:name", api.DeleteWorkspace)

	r.GET("/pods", api.GetAllPods)
	r.POST("/pods", api.CreatePod)
	r.GET("/pods/:name", api.GetPods)
	// r.PUT("/pods/:name", api.UpdatePods)
	// r.DELETE("/pods/:name", api.DeletePods)

	r.GET("/jobs", api.GetAllJobs)
	r.POST("/jobs", api.CreateJob)
	r.GET("/jobs/:name", api.GetJobs)
	// r.PUT("/jobs/:name", api.UpdateJob)
	// r.DELETE("/jobs/:name", api.DeleteJob)

	r.GET("/cronjobs", api.GetCronAllJobs)
	r.POST("/cronjobs", api.CreateCronJob)
	r.GET("/cronjobs/:name", api.GetCronJobs)
	// r.PUT("/cronjobs/:name", api.UpdateCronJobs)
	// r.DELETE("/cronjobs/:name", api.DeleteCronJobs)

	r.GET("/services", api.GetServices)
	r.POST("/services", api.CreateService)
	r.GET("/services/:name", api.GetService)
	// r.PUT("/services/:name", api.UpdateService)
	// r.DELETE("/services/:name", api.DeleteService)

	// r.GET("/clusters", api.GetAllClusters)
	// r.POST("/clusters", api.CreateCluster)
	// r.GET("/clusters/:name", api.GetCluster)
	// r.PUT("/clusters/:name", api.UpdateCluster)
	// r.DELETE("/clusters/:name", api.DeleteCluster)

	// r.GET("/projects", api.GetAllProjects)
	// r.POST("/projects", api.CreateProject)
	// r.GET("/projects/:name", api.GetProject)
	// r.PUT("/projects/:name", api.UpdateProject)
	// r.DELETE("/projects/:name", api.DeleteProject)

	r2 := e.Group("/kube/v1", middleware.BasicAuth(func(id, password string, c echo.Context) (bool, error) {
		return api.AuthenticateUser(id, password), nil
	}))
	r2.Any("/:cluster_name", api.Kubernetes)
	r2.Any("/:cluster_name/:namespace_name", api.Kubernetes)
	r2.Any("/:cluster_name/:namespace_name/:kind_name", api.Kubernetes)
	r2.Any("/:cluster_name/:namespace_name/:kind_name/*", api.Kubernetes)

	r2.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	r2.GET("/monitoring", echo.WrapHandler(promhttp.Handler()))
	r2.Any("/monitoring/:kind", api.Monit)
	r2.Any("/monitoring/:kind/:name", api.Monit)
	r2.Any("/monitoring/realtime/:kind", api.RealMetrics)
}
