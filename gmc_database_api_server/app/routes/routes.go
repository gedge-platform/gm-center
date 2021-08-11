package routes

import (
	// "fmt"
	// "os"

	"gmc_database_api_server/app/api"

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

	r0 := e.Group("/test/v1", middleware.BasicAuth(func(id, password string, c echo.Context) (bool, error) {
		return api.AuthenticateUser(id, password), nil
	}))
	// r0.GET("/cluster", api.Cluster)
	r0.GET("/getClusters/:name", api.Get_Cluster)
	r0.GET("/getClusters", api.Get_Clusters)
	r0.GET("/getProjects/:name", api.Get_Project)
	r0.GET("/getDeployments/:name", api.Get_Deployment)

	r5 := e.Group("/testing/v1", middleware.BasicAuth(func(id, password string, c echo.Context) (bool, error) {
		return api.AuthenticateUser(id, password), nil
	}))
	r5.GET("/services", api.GetService)
	r5.GET("/services/:name", api.GetService)

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

	r.GET("/clusters", api.GetAllClusters)
	r.POST("/clusters", api.CreateCluster)
	r.GET("/clusters/:name", api.GetCluster)
	r.PUT("/clusters/:name", api.UpdateCluster)
	r.DELETE("/clusters/:name", api.DeleteCluster)

	r.GET("/projects", api.GetAllProjects)
	r.POST("/projects", api.CreateProject)
	r.GET("/projects/:name", api.GetProject)
	r.PUT("/projects/:name", api.UpdateProject)
	r.DELETE("/projects/:name", api.DeleteProject)

	r.GET("/workspaces", api.GetAllWorkspaces)
	r.POST("/workspaces", api.CreateWorkspace)
	r.GET("/workspaces/:name", api.GetWorkspace)
	r.PUT("/workspaces/:name", api.UpdateWorkspace)
	r.DELETE("/workspaces/:name", api.DeleteWorkspace)

	// r.GET("/jobs", api.GetAllJobs)
	// r.POST("/jobs", api.CreateJobs)
	r.GET("/jobs/:name", api.GetJobs)
	// r.PUT("/jobs/:name", api.UpdateJobs)
	// r.DELETE("/jobs/:name", api.DeleteJobs)
	// r.GET("/cronjobs", api.GetAllCronJobs)

	r.GET("/cronjobs/:name", api.GetCronJobs)

	r.GET("/pods/:name", api.GetPods)

	r2 := e.Group("/kube/v1", middleware.BasicAuth(func(id, password string, c echo.Context) (bool, error) {
		return api.AuthenticateUser(id, password), nil
	}))
	r2.Any("/:cluster_name", api.Kubernetes)
	r2.Any("/:cluster_name/:namespace_name", api.Kubernetes)
	r2.Any("/:cluster_name/:namespace_name/:kind_name", api.Kubernetes)
	r2.Any("/:cluster_name/:namespace_name/:kind_name/*", api.Kubernetes)

	r2.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	r2.GET("/monitoring", echo.WrapHandler(promhttp.Handler()))
	r2.Any("/monitoring/:kind", api.Metrics)
	r2.Any("/monitoring/:kind/:name", api.Metrics)
	r2.Any("/monitoring/realtime/:kind", api.RealMetrics)
}
