package api

import (
	"gmc_database_api_server/app/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Cluster(c echo.Context) (err error) {
	data, err := common.GetModel(c, "pods")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	// log.Println("[34] data is ", data)

	// model.Cluster{Monitoring: new(model.Monitoring2)}

	// log.Println("data is ", data)
	return c.JSON(200, echo.Map{
		"items": data,
	})
}
