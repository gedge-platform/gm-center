package routes

import (
	"os"

	c "gmc_api_gateway/app/controller"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

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

	e.POST("/gmcapi/v2/auth", c.LoginUser)

	r0 := e.Group("/gmcapi/v2/restricted")

	// decoded, err := base64.URLEncoding.DecodeString(os.Getenv("SIGNINGKEY"))
	// if err != nil {
	// 	fmt.Println("signingkey base64 decoded Error")
	// }

	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(os.Getenv("SIGNINGKEY")),
	}

	r0.Use(middleware.JWTWithConfig(config))
	r0.GET("/test", c.ListMember)

	// /gmcapi/v2
	r := e.Group("/gmcapi/v2", middleware.JWTWithConfig(config))

	r.POST("/members", c.CreateMember)
	r.GET("/members", c.ListMember)
	r.GET("/members/:memberId", c.FindMember)
	r.DELETE("/members/:memberId", c.DeleteMember)
	r.PUT("/members/:memberId", c.UpdateMember)

	r.POST("/cluster", c.CreateCluster)
	r.GET("/cluster", c.ListCluster)
	r.GET("/cluster/:name", c.FindCluster)
	r.DELETE("/cluster/:name", c.DeleteCluster)
	r.PUT("/cluster/:name", c.UpdateCluster)

	r.POST("/workspace", c.CreateWorkspace)
	r.GET("/workspace", c.ListWorkspace)
	r.GET("/workspace/:name", c.FindWorkspace)
	r.DELETE("/workspace/:name", c.DeleteWorkspace)
	r.PUT("/workspace/:name", c.UpdateWorkspace)

	r.POST("/projects", c.CreateProject)
	r.GET("/userProjects", c.ListUserProject)
	r.GET("/systemProjects", c.ListSystemProject)
	r.GET("/userProjects/:name", c.GetUserProject)
	r.GET("/systemProjects/:name", c.GetSystemProject)
	r.DELETE("/projects/:name", c.DeleteProject)
	r.PUT("/projects/:name", c.UpdateProject)

	r.POST("/request", c.CreateRequest)
	r.GET("/request", c.ListRequest)
	r.GET("/request/:requestId", c.FindRequest)
	r.DELETE("/request/:requestId", c.DeleteRequest)
	r.PUT("/request/:requestId", c.UpdateRequest)

	r.GET("/vm/:vm_name/suspend", c.VmSuspend)
	r.GET("/vm/:vm_name/resume", c.VmResume)
	r.GET("/vm/:vm_name/reboot", c.VmReboot)
	r.DELETE("/vm/:vm_name/terminate", c.VmTerminate)

	r.GET("/vm", c.GetALLVm)
	r.GET("/view/:name", c.GetView)

	r.GET("/deployments", c.GetDeployments)
	r.POST("/deployments", c.CreateDeployment)
	r.GET("/deployments/:name", c.GetDeployment)
	// r.PUT("/deployments/:name", c.UpdateDeployment)
	r.DELETE("/deployments/:name", c.DeleteDeployment)

	r.GET("/pods", c.GetAllPods)
	r.POST("/pods", c.CreatePod)
	r.GET("/pods/:name", c.GetPods)
	// r.PUT("/pods/:name", c.UpdatePods)
	r.DELETE("/pods/:name", c.DeletePod)

	r.GET("/jobs", c.GetAllJobs)
	r.POST("/jobs", c.CreateJob)
	r.GET("/jobs/:name", c.GetJobs)
	// r.PUT("/jobs/:name", c.UpdateJob)
	r.DELETE("/jobs/:name", c.DeleteJob)

	r.GET("/cronjobs", c.GetCronAllJobs)
	r.POST("/cronjobs", c.CreateCronJob)
	r.GET("/cronjobs/:name", c.GetCronJobs)
	// r.PUT("/cronjobs/:name", c.UpdateCronJobs)
	r.DELETE("/cronjobs/:name", c.DeleteCronJob)

	r.GET("/services", c.GetServices)
	r.POST("/services", c.CreateService)
	r.GET("/services/:name", c.GetService)
	// r.PUT("/services/:name", c.UpdateService)
	r.DELETE("/services/:name", c.DeleteService)

	r.GET("/pvs", c.GetAllPVs)
	r.POST("/pvs", c.CreatePV)
	r.GET("/pvs/:name", c.GetPV)
	// // r.PUT("/services/:name", c.UpdateService)
	r.DELETE("/pvs/:name", c.DeletePV)

	r.GET("/pvcs", c.GetAllPVCs)
	r.POST("/pvcs", c.CreatePVC)
	r.GET("/pvcs/:name", c.GetPVC)
	// // r.PUT("/services/:name", c.UpdateService)
	r.DELETE("/pvcs/:name", c.DeletePVC)

	r.GET("/secrets", c.GetAllSecrets)
	// r.POST("/pvs", c.CreateService)
	r.GET("/secrets/:name", c.GetSecret)
	// // r.PUT("/services/:name", c.UpdateService)
	// r.DELETE("/pvs/:name", c.DeleteService)

	r.GET("/storageclasses/:name", c.GetStorageclass)
	r.GET("/storageclasses", c.GetStorageclasses)
	r.POST("/storageclasses", c.CreateStorageclasses)
	r.DELETE("/storageclasses/:name", c.DeleteStorageclasses)

	r.GET("/clusterroles/:name", c.GetClusterRole)
	r.GET("/clusterroles", c.GetClusterRoles)

	r.GET("/roles/:name", c.GetRole)
	r.GET("/roles", c.GetRoles)

	r.GET("/configmaps", c.GetAllConfigmaps)
	r.GET("/configmaps/:name", c.GetConfigmap)
	r.GET("/daemonsets", c.GetAllDaemonsets)
	// r.POST("/pvs", c.CreateService)
	r.GET("/daemonsets/:name", c.GetDaemonset)

	r.GET("/statefulsets", c.GetAllStatefulset)
	// r.POST("/pvs", c.CreateService)
	r.GET("/statefulsets/:name", c.GetStatefulset)
	// // r.PUT("/services/:name", c.UpdateService)
	// r.DELETE("/pvs/:name", c.DeleteService)

	r.GET("/serviceaccounts", c.GetAllServiceaccounts)
	r.GET("/serviceaccounts/:name", c.GetServiceaccount)

	r.GET("/clusterrolebindings", c.GetAllClusterrolebindings)
	r.GET("/clusterrolebindings/:name", c.GetClusterrolebinding)
	r.POST("/clusterrolebindings", c.CreateClusterRolebinding)
	r.DELETE("/clusterrolebindings/:name", c.DeleteClusterRolebinding)

	r.GET("/duplicateCheck/:name", c.DuplicateCheckDB)

	r.GET("/view/:name", c.GetView)

	e.GET("/clusterInfo", c.GetClusterInfo)

	r2 := e.Group("/kube/v1", middleware.BasicAuth(func(id, password string, echo echo.Context) (bool, error) {
		userChk, _ := c.AuthenticateUser(id, password)
		return userChk, nil
	}))
	// r2.Any("/:cluster_name", api.Kubernetes)
	// r2.Any("/:cluster_name/:namespace_name", api.Kubernetes)
	// r2.Any("/:cluster_name/:namespace_name/:kind_name", api.Kubernetes)
	// r2.Any("/:cluster_name/:namespace_name/:kind_name/*", api.Kubernetes)

	r2.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	r2.GET("/monitoring", echo.WrapHandler(promhttp.Handler()))
	r2.Any("/monitoring/:kind", c.Monit)
	// r2.Any("/monitoring/:kind/:name", api.Monit)
	r2.Any("/monitoring/realtime/:kind", c.RealMetrics)
	r2.Any("/monitoring/realtime", c.RealMetrics)
}
