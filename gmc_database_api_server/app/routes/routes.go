package routes

import (
	// "fmt"
	// "os"

	"gmc_api_gateway/app/api"

	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	"github.com/go-playground/validator/v10"
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
	r := e.Group("/api/v1/members")
	r.GET("", api.GetAllMembers)
	r.POST("", api.CreateMember)
	r.GET("/:id", api.GetMember)
	r.PUT("/:id", api.UpdateMember)
	r.DELETE("/:id", api.DeleteMember)

	r2 := e.Group("/api/v1/apps")
	r2.GET("", api.GetAllApps)
	r2.POST("", api.CreateApp)
	r2.GET("/:name", api.GetApp)
	r2.PUT("/:name", api.UpdateApp)
	r2.DELETE("/:name", api.DeleteApp)

	r3 := e.Group("/api/v1/clusters")
	r3.GET("", api.GetAllClusters)
	r3.POST("", api.CreateCluster)
	r3.GET("/:name", api.GetCluster)
	r3.PUT("/:name", api.UpdateCluster)
	r3.DELETE("/:name", api.DeleteCluster)

	r4 := e.Group("/api/v1/projects")
	r4.GET("", api.GetAllProjects)
	r4.POST("", api.CreateProject)
	r4.GET("/:name", api.GetProject)
	r4.PUT("/:name", api.UpdateProject)
	r4.DELETE("/:name", api.DeleteProject)

	r5 := e.Group("/api/v1/workspaces")
	r5.GET("", api.GetAllWorkspaces)
	r5.POST("", api.CreateWorkspace)
	r5.GET("/:name", api.GetWorkspace)
	r5.PUT("/:name", api.UpdateWorkspace)
	r5.DELETE("/:name", api.DeleteWorkspace)
	
	e.Any("/api/v1/kubernetes/*", api.Kubernetes)
}