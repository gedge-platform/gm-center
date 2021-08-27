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
func CreateNamespace(c echo.Context) (err error) {

	//workspace, user , cluster(selected) , project = namespace를 생성한다.
	/*
	   {
	       "projectName": "terstsee",
	       "projectDescription": "zzz",
	       "projectType": "user",
	       "projectOwner": "innogrid",
	       "projectCreator": "innogrid",
	       "selectCluster": "cluster1",
	       "workspaceName": "multi_test"
	   }
	*/
	//1. Post 요청대로 Project DB에 저장 요청
	fmt.Println("DB 저장")
	// CreateProjectDB(c) <- 여기서 c.bind 해서 아래 c.bind 랑 중복 둘중 하나 제거
	fmt.Println("DB 저장 OK")
	//2. Post 요청대로 kubernetes에 요청하기. (clusterList 만큼 반복하기)

	fmt.Println("Model 만들기")
	models := new(model.Project)

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return err
	}
	if err = c.Validate(models); err != nil {
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return err
	}
	fmt.Println("Model 만들기 OK")

	fmt.Println(models.SelectCluster) //cluster가 여러개 올수도 있음 파싱 함수 및 리스트로 처리하는거 수정해야함.

	models2 := GetCluster3(models.SelectCluster)
	fmt.Println(models2)

	// if err != nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, err)
	// 	return nil
	// }
	// fmt.Printf("[#55555]data is info %s", getData)

	//2-1. model 얻어오기
	//2-2. model 정보에서 url,token 추출하기

	//2-3. clusterList 만큼 반복 호출 (projectName으로 namespace 생성)
	fmt.Println("Kubernetes 요청")
	if check := strings.Compare(models.Name, "") == 0; check {
		fmt.Println("input project name")
	} else {
		namesapce := Namespace{}

		namesapce.APIVersion = "v1"
		namesapce.Kind = "Namespace"
		namesapce.Metadata.Name = models.Name

		url := "https://" + models2.Endpoint + ":6443/api/v1/namespaces/"
		Token := models2.Token

		data, err := json.Marshal(namesapce)

		if err != nil {
			fmt.Println(err)
			// return false
		}

		var jsonStr = []byte(fmt.Sprint(string(data)))

		code, _ := RequsetKube(url, "POST", jsonStr, Token)

		if code == 201 {
			fmt.Println("ok")
		} else {
			fmt.Println("false")
		}
	}
	fmt.Println("Kubernetes 요청완료")
	return nil
}

// func CreateServiceAccount(name string, namespace string, models *model.ClusterToken) bool {
// 	// create service account
// 	accounts := model.Account{}

// 	accounts.Metadata.Name = name
// 	accounts.Metadata.Namespace = namespace

// 	url := "https://" + models.Endpoint + ":6443/api/v1/namespaces/" + namespace + "/serviceaccounts"

// 	data, err := json.Marshal(accounts)
// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	var jsonStr = []byte(fmt.Sprint(string(data)))

// 	code, _ := RequsetKube(url, "POST", jsonStr, models.Token)

// 	if code == 201 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

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

func CreateProjectDB(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.Project)

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return err
	}
	if err = c.Validate(models); err != nil {
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return err
	}

	if err != nil {
		panic(err)
	}

	if err := db.Create(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return err
	}

	return nil
}

func GetCluster3(cluster string) *model.Cluster {
	search_val := cluster
	db := db.DbManager()
	// search_val := c.Param("name")
	models := FindClusterDB(db, "Name", search_val)

	if models == nil {
		// common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}

	return models
}
