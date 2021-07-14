package api

import (
	"log"
	"net/http"
	"strings"
	"crypto/tls"
	"time"
	"io"
	"io/ioutil"
	"bytes"

	"gmc_api_gateway/app/db"
	"github.com/labstack/echo"
	"github.com/tidwall/gjson"
	"github.com/go-resty/resty/v2"
)



func Kubernetes(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("cluster_name")
	models := FindClusterDB(db, "Name", search_val)

	if models == nil {
		var msgError messageFormat
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}
	
	url := GetClusterEP(c)

	if url == "" {
		return nil
	} else if _, err := HttpRequest(c, url, true); err != nil {
		log.Println("HttpRequest error")
	} 

	return nil
}

func GetClusterEP(c echo.Context) (url string) {
	db := db.DbManager()
	cluster_name := c.Param("cluster_name")
	namespace_name := c.Param("namespace_name")
	kind_name := c.Param("kind_name")
	returnUrl := ""
	detailUrl := getDetailURL(c)
	models := FindClusterDB(db, "Name", cluster_name)
	if models == nil {
		return "nf"
	}

	var kindUrl, kindName, returnUrl2 string
	if detailUrl == "jobs_podlist" {
		kindUrl, kindName = GetKindURL("pods")
		returnUrl2 = "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(namespace_name) + kindName
	}
	
	kindUrl, kindName = GetKindURL(kind_name)
	returnUrl = "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(namespace_name) + kindName


	if c.Param("*") == "\n" || c.Param("*") == "" {
		returnUrl = returnUrl + c.Param("*")
	} else {
		returnUrl = returnUrl + "/" + c.Param("*")
	}
	
	if detailUrl == "jobs_podlist" {
		returnUrl = strings.TrimRight(returnUrl, "/list")
	} else if detailUrl == "cronjobs_joblist" {
		returnUrl = strings.TrimRight(returnUrl, "/list")
	}


	log.Printf("# cluster name : " + cluster_name)
	log.Printf("# API : %s + %s",returnUrl, c.Param("*"))
	log.Printf("# API Url is : " + returnUrl)
	log.Printf("# API Url #2 is : " + returnUrl2)


	if detailUrl == "jobs_podlist" {
		if data, err := HttpRequest(c, returnUrl, false); err != nil {
			log.Println("HttpRequest error")
			return returnUrl
		} else {
			log.Printf("jobs_podlist IN # returnUrl : %s", returnUrl)
			job_name := gjson.Get(data, "metadata.labels.job-name")
			if err := getDetailList(c, detailUrl, returnUrl2, job_name.String()); err=="nf" {
				log.Println("err")
				return returnUrl
			} else {
				return err
			}
	}
}

	return returnUrl
}

func GetKindURL(kindParam string) (url string, kind string) {
	var kindUrl string
	if kindParam == "deployments" {
		kindUrl = "apis/apps/v1/namespaces/"
	} else if kindParam == "jobs" || kindParam == "cronjobs" {
		kindUrl = "apis/batch/v1/namespaces/"
	} else if kindParam == "clusterroles" || kindParam == "roles" || kindParam == "clusterrolebindings" {
		kindUrl = "apis/rbac.authorization.k8s.io/v1/namespaces/"
	} else if kindParam == "networkpolicies" {
		kindUrl = "/apis/networking.k8s.io/v1/namespaces/"
	} else if kindParam == "nodes" {
		kindUrl = "api/v1/namespaces/"
	} else if kindParam == "events" {
		kindUrl = "apis/events.k8s.io/v1/namespaces/"
	} else if kindParam == "\n" || kindParam == "" {
		kindUrl = "api/v1/"
	} else {
		kindUrl = "api/v1/namespaces/"
	}

	return kindUrl, kindParam
}

func GetNamespaceURL(namespaceParam string) (url string) {
	var namespaceUrl string
	if namespaceParam == "\n" || namespaceParam == "" {
		namespaceUrl = namespaceParam + "/"
	} else if namespaceParam == "default" {
		namespaceUrl = namespaceParam + "/"
	}	else {
		namespaceUrl = namespaceParam + "/"
	} 

	return namespaceUrl
}

func getDetailURL(c echo.Context) (value string) {
	var multiParamChk []string
	var multiParamLen int
	// var SearchVal string

	multiParamChk = strings.Split(c.Param("*"),"/")
	multiParamLen = len(multiParamChk)
	kind_name := c.Param("kind_name")

	log.Printf("###getDetailURL Start###")
	log.Printf("kind : " + kind_name)
	log.Printf("Resource Name : " + multiParamChk[0])
	log.Printf("Resource Len : %d", multiParamLen)

	if multiParamLen >= 2 {
		if kind_name == "jobs" {
			// testVal = testString( strings.Split(c.Param("*"),"/") )
			log.Printf("check Param : " + multiParamChk[1])
			if multiParamChk[1] == "list" {
				log.Printf("###Deployment list###")
				value = "jobs_podlist"
				return value
			}
		}
	}
	return "nf"
}

func HttpRequest(c echo.Context, url string, check bool) (data string, err error) {
	var responseString, token string
	reqMethod := c.Request().Method
	passBody := responseBody(c.Request().Body)
	
	tokens, ok := c.Request().Header["Authorization"]
	if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
	}
	
	// log.Printf("URL : %v", url)
	// log.Printf("reqMethod : %v", reqMethod)
	// log.Printf("token : %v", token)
	// log.Printf("passBody : %v", passBody)


	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
	client.SetTimeout(1 * time.Minute)
	client.SetHeaders(map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type": "application/json; charset=utf-8",
		"Accept": "application/json; charset=utf-8",
	})

	// client.SetAuthToken(token)

	if reqMethod == "GET" {
		if resp, err := client.R().SetAuthToken(token).Get(url); err != nil {
			panic(err)
		} else {
			responseString =string(resp.Body())
		}
	} else if reqMethod == "POST" {
		if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Post(url); err != nil {
			panic(err)
		} else {
			responseString =string(resp.Body())
		}
	} else if reqMethod == "PATCH" {
		if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Patch(url); err != nil {
			panic(err)
		} else {
			responseString =string(resp.Body())
		}
	} else if reqMethod == "PUT" {
		if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Put(url); err != nil {
			panic(err)
		} else {
			responseString =string(resp.Body())
		}
	} else if reqMethod == "DELETE" {
		if resp, err := client.R().SetAuthToken(token).Delete(url); err != nil {
			panic(err)
		} else {
			responseString =string(resp.Body())
		}
	}

	content, ok := gjson.Parse(responseString).Value().(map[string]interface{})
	if !ok {
		panic("err")
	}

	if check == true {
		return responseString, c.JSON(http.StatusOK, content)
	} else {
		return responseString, nil
	}
	
}

func lastString(ss []string) string {
	return ss[len(ss)-1]
}

func responseBody(req io.ReadCloser) string {
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

func getDetailList(c echo.Context, kind string, url string, val string) (urls string) {

	log.Printf("###getDetailURL Start###")
	log.Printf("val is %s", val)
	log.Printf("url is %s", url)

	if data, err := HttpRequest(c, url, false); err != nil {
		log.Println("HttpRequest error")
	} else {

		if val == "" || val == "\n" {
			return "nf"
		}
		if kind == "jobs_podlist" {
			// log.Println(data)
			var podList []string
			var returnVal string
			result := gjson.Get(data, `items.#.metadata.name`)
			for _, name := range result.Array() {
				log.Printf("podlist : %s", name.String())
				if strings.Contains(name.String(), val) == true {
					returnVal = name.String()
				}
			}
			log.Printf("podList is %s", podList)
			log.Printf("returnVal is %s", returnVal)

			url = url + "/" + returnVal
			log.Printf("url is %s", url)
			return url
		}
	}

	return "nf"
}