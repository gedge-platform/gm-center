package main

import (
	"net/http"
	"os"
	"strconv"

	"gmc_database_api_server/app/db"
	"gmc_database_api_server/app/routes"
	"gmc_database_api_server/config"

	_ "gmc_database_api_server/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Gedge GM-Center Swagger API
// @version 1.0
// @description This is a Gedge GM-Center Swagger API.

// @contact.name consine2c
// @contact.url https://gedge-platform.github.io/gm-center/
// @contact.email consine2c@innogrid.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host g-api-test.innogrid.tech
// @BasePath /api/v1
// @schemes http https
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
				<h1>Welcome to GEdge API-Gateway!</h1>
				<h3>GEdge Platform :: GM-Center API Server :)</h3>
		`)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.GEdgeRoute(e)

	if err := e.Start(GetListenPort(config)); err != nil {
		panic(err)
	}
}

// Environment Value ("LISTEN_PORT")
func GetListenPort(config *config.Config) string {
	port := os.Getenv("LISTEN_PORT")

	if len(port) == 0 {
		port = config.COMMON.Port
	}
	intPort, err := strconv.Atoi(port)
	if err != nil || intPort < 1 || 65535 < intPort {
		port = config.COMMON.Port
	}

	return ":" + port
}