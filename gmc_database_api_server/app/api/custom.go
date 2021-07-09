package api

import (
	"log"
	"net/http"
	"io/ioutil"
	"github.com/labstack/echo"
)

func Kubernetes(c echo.Context) (err error) {
	url := kubernetes_api + c.Param("*")
	// reqMethod := c.Request().Method
	// authorizationKey := c.Request().Header["Authorization"]

	req, err := http.NewRequest(c.Request().Method, url, c.Request().Body)

	req.Header = c.Request().Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
			return nil
	} else {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)

			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			c.Response().WriteHeader(http.StatusOK)

			c.String(http.StatusOK, string(data))
	}

	return nil
}