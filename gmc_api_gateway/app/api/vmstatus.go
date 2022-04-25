package api

import (
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetVmStatus(c echo.Context) (err error) {
	param := model.VMPARAM{
		ConnectionName: c.QueryParam("ConnectionName"),
		// Method:         c.Request().Method,
	}

	cdb, _ := common.VmStatusFind(param)
	x := common.StringToMapInterface(cdb)

	return c.JSON(http.StatusOK, echo.Map{"data": x})
}

func GetVm(c echo.Context) (err error) {
	param := model.VMPARAM{
		ConnectionName: c.QueryParam("ConnectionName"),
		Method:         c.Request().Method,
	}

	cdb, _ := common.VmFind(param)
	x := common.StringToMapInterface(cdb)
	return c.JSON(http.StatusOK, echo.Map{"data": x})
}
