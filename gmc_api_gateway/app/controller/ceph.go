package controller

import (
	"encoding/json"
	"fmt"
	"gmc_api_gateway/app/common"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetCephAuthToken() (token string) {
	url := os.Getenv("CEPH") + "api/auth"
	Id := os.Getenv("CEPH_ID")
	Pw := os.Getenv("CEPH_PW")
	method := "POST"
	data := make(map[string]interface{})
	data["username"] = Id
	data["password"] = Pw
	payload, _ := json.Marshal(data)
	payloadstr := strings.NewReader(string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payloadstr)
	req.Header.Add("Accept", "application/vnd.ceph.api.v1.0+json")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	body_str := string(body)
	temp := common.StringToMapInterface(body_str)
	result := temp["token"]
	fmt.Println(temp["token"])
	return common.InterfaceToString(result)
}

func GetCephHealth(c echo.Context) (err error) {
	token := GetCephAuthToken()
	// params := model.PARAMS{
	// 	Kind:      "nodes",
	// 	Name:      c.Param("name"),
	// 	Cluster:   c.QueryParam("cluster"),
	// 	Workspace: c.QueryParam("workspace"),
	// 	Project:   c.QueryParam("project"),
	// 	Method:    c.Request().Method,
	// 	Body:      responseBody(c.Request().Body),
	// }
	// config.Init()
	// addr := os.Getenv("CEPH")
	// Id := os.Getenv("CEPH_ID")
	// Pw := os.Getenv("CEPH_PW")

	return c.JSON(http.StatusOK, echo.Map{
		"data": token,
	})
}
