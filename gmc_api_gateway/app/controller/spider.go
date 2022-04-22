package controller

import (
	"fmt"
	"gmc_api_gateway/app/common"
	"net/http"

	"gmc_api_gateway/app/model"

	"github.com/labstack/echo/v4"
)

func GetALLVm(c echo.Context) (err error) {

	/* params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("name"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)

	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,

	}) */

	params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("name"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	body := make(map[string]string)
	_ = c.Bind(&body)

	getData, err := common.DataRequest(params)

	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})

}

func GetVm(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("vmName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)
	vm := StringToInterface(getData)
	fmt.Println(params.Body)
	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})

}

func CreateVm(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("vmName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)

	vm := StringToInterface(getData)
	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})
}

func DeleteVm(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("vmName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	fmt.Println(params.Body)
	getData, err := common.DataRequest(params)

	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})

}

func GetALLCredential(c echo.Context) (err error) {

	/* params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("name"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)

	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,

	}) */

	params := model.PARAMS{
		Kind:   "credential",
		Name:   c.Param("name"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)

	credential := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})

}

func GetCredential(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "credential",
		Name:   c.Param("credentialName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)
	credential := StringToInterface(getData)
	fmt.Println(params.Body)
	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})

}

func CreateCredential(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:   "credential",
		Name:   c.Param("credentialName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)

	credential := StringToInterface(getData)
	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})
}

func DeleteCredential(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "credential",
		Name:   c.Param("credentialName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody(c.Request().Body),
	}

	fmt.Println(params.Body)
	getData, err := common.DataRequest(params)

	credential := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})

}
