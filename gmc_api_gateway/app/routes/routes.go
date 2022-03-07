package routes

import (

	// "gmc_api_gateway/app/api"
	c "gmc_api_gateway/app/controller"

	"github.com/labstack/echo/v4"
)

// type jwtCustomClaims struct {
// 	Name string `json:"name"`
// 	Role string `json:"role"`
// 	jwt.StandardClaims
// }

// type DataValidator struct {
// 	validator *validator.Validate
// }

// func NewValidator() *DataValidator {
// 	return &DataValidator{
// 		validator: validator.New(),
// 	}
// }

// func (dv *DataValidator) Validate(i interface{}) error {
// 	return dv.validator.Struct(i)
// }

func GEdgeRoute(e *echo.Echo) {
	// e.Validator = NewValidator()

	// e.POST("/gmcapi/v1/auth", api.LoginUser)

	// r0 := e.Group("/gmcapi/v1/restricted")

	// decoded, err := base64.URLEncoding.DecodeString(os.Getenv("SIGNINGKEY"))
	// if err != nil {
	// 	fmt.Println("signingkey base64 decoded Error")
	// }

	// config := middleware.JWTConfig{
	// 	Claims:     &jwtCustomClaims{},
	// 	SigningKey: []byte(os.Getenv("SIGNINGKEY")),
	// }

	// r0.Use(middleware.JWTWithConfig(config))
	// r0.GET("/test", api.GetAllMembers)

	// /gmcapi/v1
	r := e.Group("/gmcapi/v1")
	// r := e.Group("/gmcapi/v1", middleware.BasicAuth(func(id, password string, c echo.Context) (bool, error) {
	// 	userChk, _ := api.AuthenticateUser(id, password)
	// 	return userChk, nil
	// }))
	r.GET("/members", c.ListMember)
	r.GET("/members/:id", c.FindMember)
	r.DELETE("/members/:id", c.DeleteMember)
	r.POST("/members", c.CreateMember)
	// r.PUT("/members/:id", c.UpdateMember)
	r.POST("/workspace", c.CreateWorkspace)
	r.GET("/workspace", c.ListWorkspace)
	r.GET("/workspace/:workspaceName", c.FindWorkspace)
	r.DELETE("/workspace/:workspaceName", c.DeleteWorkspace)
	r.POST("/cluster", c.CreateCluster)
	r.GET("/cluster", c.ListCluster)
	r.GET("/cluster/:clusterName", c.FindCluster)
	r.DELETE("/cluster/:clusterName", c.DeleteCluster)
}
