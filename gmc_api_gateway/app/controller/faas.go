package controller

import (
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)


func GetEnvironments(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind: "environments",
		Type:      "faas",
		QueryString: c.QueryString(),
		Project:   c.QueryParam("project"),
		Method:      c.Request().Method,
		Body:        responseBody(c.Request().Body),
	}

	getData, err := common.DataRequest_scheduler(params)
	if err != nil {
			return c.JSON(http.StatusBadRequest, getData)
		} else {
			return c.JSON(http.StatusOK, common.StringToInterface(getData))
			// return c.JSON(http.StatusOK, echo.Map{
			// "data": common.StringToInterface(getData),
			// })
		} 
}

func GetEnvironment(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind: "environments",
		Type: "faas",
		Name:      c.Param("name"),
		QueryString: c.QueryString(),
		Project:   c.QueryParam("project"),
		Method:      c.Request().Method,
		Body:        responseBody(c.Request().Body),
	}
	
	getData, err := common.DataRequest_scheduler(params)
	if err != nil   {
			return c.JSON(http.StatusBadRequest, getData)
		} else {
			return c.JSON(http.StatusOK, common.StringToInterface(getData))
		} 
}

func CreateEnvironment(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind: "environments",
		Type: "faas",
		QueryString: c.QueryString(),
		Project:   c.QueryParam("project"),
		Method:      c.Request().Method,
		Body:        responseBody(c.Request().Body),
	}
	
	postData, err := common.DataRequest_scheduler(params)

	if err != nil || strings.Contains(postData, "Error") {
			return c.String(http.StatusBadRequest, postData )
		} else {
			// return json.NewEncoder(c.Response()).Encode(u)
			return c.JSON(http.StatusCreated, common.StringToInterface(postData))
		} 
}

func DeleteEnvironment(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind: "environments",
		Type: "faas",
		Name:      c.Param("name"),
		QueryString: c.QueryString(),
		Project:   c.QueryParam("project"),
		Method:      c.Request().Method,
		Body:        responseBody(c.Request().Body),
	}
	
	postData, err := common.DataRequest_scheduler(params)

	if err != nil || strings.Contains(postData, "Error") {
			return c.String(http.StatusBadRequest, postData )
		} else {
			// return json.NewEncoder(c.Response()).Encode(u)
			return c.JSON(http.StatusCreated, common.StringToInterface(postData))
		} 
}