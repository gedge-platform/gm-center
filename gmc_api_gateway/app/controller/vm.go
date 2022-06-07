package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gmc_api_gateway/app/common"

	"gmc_api_gateway/app/model"

	"github.com/labstack/echo/v4"
)

const (
	defaultbaseURL = "http://210.207.104.188:1024/spider/controlvm/"
)

func VmSuspend(c echo.Context) error {

	method := c.Request().Method

	log.Println("[method] is : ", c.Request().Method)
	log.Println("[vm_name] is : ", c.Param("vm_name"))

	url := defaultbaseURL + c.Param("vm_name") + "?action=suspend"

	log.Println("[url] is : ", url)

	body := c.Request().Body
	var i map[string]interface{}

	data, _ := common.HttpRequest(method, url, body)

	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	if err := json.Unmarshal([]byte(jsonResp), &i); err != nil {
		fmt.Printf("Error : %s\n", err)
	}

	return c.JSON(http.StatusOK, &i)
}

func VmResume(c echo.Context) error {

	method := c.Request().Method

	log.Println("[method] is : ", c.Request().Method)
	log.Println("[vm_name] is : ", c.Param("vm_name"))

	url := defaultbaseURL + c.Param("vm_name") + "?action=resume"

	log.Println("[url] is : ", url)

	body := c.Request().Body
	var i map[string]interface{}

	data, _ := common.HttpRequest(method, url, body)

	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	if err := json.Unmarshal([]byte(jsonResp), &i); err != nil {
		fmt.Printf("Error : %s\n", err)
	}

	return c.JSON(http.StatusOK, &i)
}

func VmReboot(c echo.Context) error {

	method := c.Request().Method

	log.Println("[method] is : ", c.Request().Method)
	log.Println("[vm_name] is : ", c.Param("vm_name"))

	url := defaultbaseURL + c.Param("vm_name") + "?action=reboot"

	log.Println("[url] is : ", url)

	body := c.Request().Body
	var i map[string]interface{}

	data, _ := common.HttpRequest(method, url, body)

	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	if err := json.Unmarshal([]byte(jsonResp), &i); err != nil {
		fmt.Printf("Error : %s\n", err)
	}

	return c.JSON(http.StatusOK, &i)
}

func VmTerminate(c echo.Context) error {

	method := c.Request().Method

	log.Println("[method] is : ", c.Request().Method)
	log.Println("[vm_name] is : ", c.Param("vm_name"))

	url := defaultbaseURL + c.Param("vm_name") + "?action=terminate"

	log.Println("[url] is : ", url)

	body := c.Request().Body
	var i map[string]interface{}

	data, _ := common.HttpRequest(method, url, body)

	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	if err := json.Unmarshal([]byte(jsonResp), &i); err != nil {
		fmt.Printf("Error : %s\n", err)
	}

	return c.JSON(http.StatusOK, &i)
}

func GetALLVm(c echo.Context) (err error) {

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
