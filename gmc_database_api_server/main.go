package main

import (
	"net/http"
  
	"gmc_api_gateway/config"
	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.Init()
	config := config.GetConfig()

	app := &db.DB{}
	app.Initialize(config)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// e.Use(middleware.Gzip())
	// e.Use(middleware.Secure())

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{config.COMMON.CorsOrigin},
	// 	AllowHeaders: []string{"Authorization"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	// }))

	e.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, "GEdge Platform :: GM-Center API Server")
	})

	routes.GEdgeRoute(e)

	port := config.COMMON.Port
	e.Logger.Fatal(e.Start(port))
}
