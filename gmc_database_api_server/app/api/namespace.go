package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/db"
	"gmc_database_api_server/app/model"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Namespace struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
	} `json:"spec"`
	Status struct {
	} `json:"status"`
}

//프로젝트 생성
func CreateProjects(c echo.Context) (err error) {

	err, models := CreateProjectDB(c)

	if err != nil {
		return err
	}

	fmt.Println("1")
	selectCluster := models.SelectCluster
	slice := strings.Split(selectCluster, ",")

	for _, cluster := range slice {

		clusters := GetCluster3(cluster)

		namesapce := Namespace{}

		namesapce.APIVersion = "v1"
		namesapce.Kind = "Namespace"
		namesapce.Metadata.Name = models.Name

		url := "https://" + clusters.Endpoint + ":6443/api/v1/namespaces/"
		Token := clusters.Token

		data, err := json.Marshal(namesapce)

		if err != nil {
			common.ErrorMsg(c, http.StatusBadRequest, err)
			return err
		}

		var jsonStr = []byte(fmt.Sprint(string(data)))

		code, _ := RequsetKube(url, "POST", jsonStr, Token)

		if code != 201 {
			common.ErrorMsg(c, http.StatusBadRequest, err)
			return err
		}
	}
	fmt.Println("2")

	return nil
}

func RequsetKube(url string, method string, reqdata []byte, token string) (int, string) {

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqdata))

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, ""
	}
	fmt.Println(string(body))

	return res.StatusCode, string(body)
}

func CreateProjectDB(c echo.Context) (err error, st *model.Project) {
	db := db.DbManager()
	models := new(model.Project)

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return err, models
	}
	if err = c.Validate(models); err != nil {
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return err, models
	}

	if check := strings.Compare(models.Name, "") == 0; check {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return err, models
	}

	if err != nil {
		panic(err)
	}

	if err := db.Create(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return err, models
	}

	return nil, models
}

func GetCluster3(cluster string) *model.Cluster {
	search_val := cluster
	db := db.DbManager()
	// search_val := c.Param("name")
	models := FindClusterDB(db, "Name", search_val)

	// if models == nil {
	// 	// common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
	// 	// return nil
	// }

	return models
}
