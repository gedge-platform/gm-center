package api

import (
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCredentials(c echo.Context) (err error) {
	//	param := model.PARAM{
	//		CredentialName: "",
	//		Method:         c.Request().Method,
	//	}
	cdb, _ := common.DataRequest()
	x := common.StringToMapInterface(cdb)

	return c.JSON(http.StatusOK, echo.Map{"data": x})
}

func GetCredential(c echo.Context) (err error) {
	param := model.PARAM{
		CredentialName: c.Param("credentialName"),
		Method:         c.Request().Method,
	}
	cdb, _ := common.DataFind(param)
	x := common.StringToMapInterface(cdb)
	return c.JSON(http.StatusOK, echo.Map{"data": x})
}

func CreateCredential(c echo.Context) (err error) {
	param := model.PARAM{
		CredentialName:   c.QueryParam("credentialName"),
		DomainName:       c.Param("domainName"),
		IdentityEndPoint: c.Param("identityEndPoint"),
		Password:         c.Param("password"),
		ProjectID:        c.Param("projectID"),
		Username:         c.Param("username"),
		Method:           c.Request().Method}
	cdb, _ := common.DataFind(param)
	x := common.StringToMapInterface(cdb)
	return c.JSON(http.StatusCreated, echo.Map{"credential": x})
}

func DeleteCredential(c echo.Context) (err error) {
	param := model.PARAM{
		CredentialName: c.Param("credentialName"),
		Method:         c.Request().Method,
	}
	cdb, _ := common.DataFind(param)
	x := common.StringToMapInterface(cdb)
	return c.JSON(http.StatusOK, echo.Map{"data": x})
}
