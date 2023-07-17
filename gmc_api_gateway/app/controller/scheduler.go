package controller

import (
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"
	"log"
	"github.com/labstack/echo/v4"
)

func PostScheduler(c echo.Context) (err error) {
	params := model.PARAMS{
		QueryString: c.QueryString(),
		Method:      c.Request().Method,
		Body:        responseBody(c.Request().Body),
	}

	// postData, err := common.DataRequest(params)
	// if err != nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, err)
	// 	return nil
	// } else {
	// 	return c.JSON(http.StatusCreated, echo.Map{
	// 		"status": "Created",
	// 		"code":   http.StatusCreated,
	// 		"data":   postData,
	// 	})
	// }

	postData, err := common.DataRequest_scheduler(params)
	if err != nil {
		// common.ErrorMsg(c, http.StatusNotFound, err)
		// return nil

		return c.JSON(http.StatusNotFound, echo.Map{
			"status": "error",
			"code":   http.StatusNotFound,
			"data":   err,
		})
	} else {
		return c.JSON(http.StatusCreated, echo.Map{
			"status": "Created",
			"code":   http.StatusCreated,
			"data":   postData,
		})
	}
}

func CallbackScheduler(c echo.Context) (err error) {
	params := model.PARAMS{
		// QueryString: c.QueryString(),
		Method:      c.Request().Method,
		Body:        responseBody(c.Request().Body),
	}
log.Println("body : ", params.Body)
		return c.JSON(http.StatusCreated, echo.Map{
			"err":   err,
			"data":  params.Body,
		})
	}
