package common

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gmc_api_gateway/app/model"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

var listTemplates = map[string]string{
	"pods":                   "/api/v1/pods",
	"services":               "/api/v1/services",
	"endpoints":              "/api/v1/endpoints",
	"configmaps":             "/api/v1/configmaps",
	"serviceaccounts":        "/api/v1/serviceaccounts",
	"resourcequota":          "/api/v1/resourcequotas",
	"deployments":            "/apis/apps/v1/deployments",
	"replicasets":            "/apis/apps/v1/replicasets",
	"daemonsets":             "/apis/apps/v1/daemonsets",
	"statefulsets":           "/apis/apps/v1/statefulsets",
	"jobs":                   "/apis/batch/v1/jobs",
	"cronjobs":               "/apis/batch/v1/cronjobs",
	"clusterroles":           "/apis/rbac.authorization.k8s.io/v1/clusterroles",
	"roles":                  "/apis/rbac.authorization.k8s.io/v1/roles",
	"clusterrolebindings":    "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings",
	"namespaces":             "/api/v1/namespaces",
	"nodes":                  "/api/v1/nodes",
	"events":                 "/apis/events.k8s.io/v1/events",
	"storageclasses":         "/apis/storage.k8s.io/v1/storageclasses",
	"persistentvolumes":      "/api/v1/persistentvolumes",
	"persistentvolumeclaims": "/api/v1/persistentvolumeclaims",
	"secrets":                "/api/v1/secrets",
	"credential":             "/spider/credential",
	"vm":                     "/spider/vm",
}

func HttpRequest(method string, url string, _body io.ReadCloser) (map[string]string, error) {

	client := resty.New()
	client.SetAllowGetMethodPayload(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(1 * time.Minute)
	client.SetHeaders(map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type":                "application/json; charset=utf-8",
		"Accept":                      "application/json; charset=utf-8",
	})

	var responseString string
	var responseStatus int
	body := ResponseBody(_body)

	switch method {
	case "GET":
		if resp, err := client.R().SetBody([]byte(body)).Get(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}

	case "POST":
		if resp, err := client.R().SetBody([]byte(body)).Post(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	case "PATCH":
		if resp, err := client.R().SetBody([]byte(body)).Patch(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	case "PUT":
		if resp, err := client.R().SetBody([]byte(body)).Put(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	case "DELETE":
		if resp, err := client.R().Delete(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	}
	return map[string]string{
		"data":   responseString,
		"status": strconv.Itoa(responseStatus),
	}, nil
}

func ResponseBody(req io.ReadCloser) string {
	var bodyBytes []byte
	if req != nil {
		bodyBytes, _ = ioutil.ReadAll(req)
	}

	// Restore the io.ReadCloser to its original state
	req = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	buf := new(bytes.Buffer)
	buf.ReadFrom(req)
	newStr := buf.String()

	return newStr
}

func DataRequest(params model.PARAMS) (data string, err error) {
	var endPoint, token_value string

	endPoint = "210.207.104.188"

	url := UrlExpr(endPoint, params.Name, params.Kind)

	log.Println("url is", url)

	switch url {
	case "noname":
		return "", ErrWorkspaceInvalid
	case "nodetail":
		return "", ErrDetailNameInvalid
	}

	log.Printf("[#31] url is %s", url)
	var responseString, token string
	r := io.NopCloser(strings.NewReader(params.Body))
	reqMethod := params.Method
	passBody := ResponseBody(r)
	//passBody := params.Body

	//body := ResponseBody(_body)

	log.Printf("[#32] passBody is %s", passBody)
	token = token_value

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(1 * time.Minute)
	client.SetHeaders(map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type":                "application/json; charset=utf-8",
		"Accept":                      "application/json; charset=utf-8",
	})
	client.SetAllowGetMethodPayload(true)
	client.SetDebug(true)

	switch reqMethod {
	case "GET":
		if resp, err := client.R().SetBody([]byte(params.Body)).Get(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
		}
	case "POST":
		if resp, err := client.R().SetBody([]byte(string(passBody))).
			SetAuthToken(token).
			Post(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())

		}

	case "PATCH":
		if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Patch(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
		}
	case "PUT":
		if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Put(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
		}
	case "DELETE":
		if resp, err := client.R().SetAuthToken(token).SetBody([]byte(params.Body)).Delete(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
		}
	}

	return responseString, nil
}

func UrlExpr(endpoint, item, kind string) string {
	check_item := strings.Compare(item, "") != 0
	var returnUrl string
	defaultUrl := "http://" + endpoint + ":1024"

	fmt.Println(item)

	if check_item {
		// project or item value exist
		returnUrl = defaultUrl + listTemplates[kind] + "/" + item
	} else {
		returnUrl = defaultUrl + listTemplates[kind]
	}

	return returnUrl

}
