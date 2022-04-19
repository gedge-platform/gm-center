package api

import (
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCredentials(c echo.Context) (err error) {
	cdb, _ := common.DataRequest()
	x := common.StringToMapInterface(cdb)

	return c.JSON(http.StatusOK, echo.Map{"data": x})
}

func GetCredential(c echo.Context) (err error) {
	param := model.PARAM{
		CredentialName: c.Param("CredentialName"),
		Method:         c.Request().Method,
	}
	cdb, _ := common.DataFind(param)
	x := common.StringToMapInterface(cdb)
	return c.JSON(http.StatusOK, echo.Map{"data": x})
}

func CreateCredential(c echo.Context) (err error) {
	param := model.PARAM{
		CredentialName:   c.QueryParam("CredentialName"),
		DomainName:       c.Param("DomainName"),
		IdentityEndPoint: c.Param("IdentityEndPoint"),
		Password:         c.Param("Password"),
		ProjectID:        c.Param("ProjectID"),
		Username:         c.Param("Username"),
		Method:           c.Request().Method}
	cdb, _ := common.DataFind(param)
	x := common.StringToMapInterface(cdb)
	return c.JSON(http.StatusCreated, echo.Map{"credential": x})
}

func DeleteCredential(c echo.Context) (err error) {
	param := model.PARAM{
		CredentialName: c.Param("CredentialName"),
		Method:         c.Request().Method,
	}
	cdb, _ := common.DataFind(param)
	x := common.StringToMapInterface(cdb)
	return c.JSON(http.StatusOK, echo.Map{"data": x})
}
