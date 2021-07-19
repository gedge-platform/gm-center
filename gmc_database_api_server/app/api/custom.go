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
	"encoding/json"

	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"

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
	detailUrl := getDetailURL(c)
	returnUrl := ""

	models := FindClusterDB(db, "Name", cluster_name)
	if models == nil {
		return "nf"
	}

	var kindUrl, kindName, returnUrl2 string

	switch detailUrl {
		case "jobs_podlist":
			kindUrl, kindName = GetKindURL("pods")
			returnUrl2 = "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(namespace_name) + kindName
		case "cronjobs_joblist":
			kindUrl, kindName = GetKindURL("jobs")
			returnUrl2 = "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(namespace_name) + kindName
		case "deployments_list":
			kindUrl, kindName = GetKindURL("replicasets")
			returnUrl2 = "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(namespace_name) + kindName
	}

	kindUrl, kindName = GetKindURL(kind_name)
	returnUrl = "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(namespace_name) + kindName


	switch c.Param("*") {
		case "":
			returnUrl = returnUrl + c.Param("*")
		case "\n":
			returnUrl = returnUrl + c.Param("*")
		default:
			returnUrl = returnUrl + "/" + c.Param("*")
	}

	switch detailUrl {
		case "jobs_podlist":
			returnUrl = strings.TrimRight(returnUrl, "/list")
		case "cronjobs_joblist":
			returnUrl = strings.TrimRight(returnUrl, "/list")
		case "deployments_list":
			returnUrl = strings.TrimRight(returnUrl, "/list")
	}

	log.Printf("# cluster name : " + cluster_name)
	log.Printf("# API : %s + %s",returnUrl, c.Param("*"))
	log.Printf("# returnUrl Url is : " + returnUrl)
	log.Printf("# returnUrl2 Url #2 is : " + returnUrl2)
	log.Printf("detailUrl is : %s", detailUrl)


	switch detailUrl {
		case "jobs_podlist":
			if data, err := HttpRequest(c, returnUrl, false); err != nil {
				log.Println("HttpRequest error")
				return returnUrl
			} else {
				Uniq := gjson.Get(data, "metadata.labels.job-name")
				if url, _ := getDetailList(c, detailUrl, returnUrl2, Uniq.String()); url=="nf" {
					return returnUrl
				} else {
					return url
				}
			}
		case "cronjobs_joblist":
			if data, err := HttpRequest(c, returnUrl, false); err != nil {
				log.Println("HttpRequest error")
				return returnUrl
			} else {
				// log.Printf("cronjobs_joblist IN # returnUrl : %s", returnUrl)
				Uniq := gjson.Get(data, "metadata.uid")
				log.Printf("Uniq : %s", Uniq)
				url, _ := getDetailList(c, detailUrl, returnUrl2, Uniq.String());
				switch url {
					case "nf":
						return returnUrl
					case "em":
						return ""
					default:
						return url
				}
			}
		case "deployments_list":
			if data, err := HttpRequest(c, returnUrl, false); err != nil {
				log.Println("HttpRequest error")
				return returnUrl
			} else {
				log.Println("########")
				log.Printf("# returnUrl : %s", returnUrl)
				Uniq := gjson.Get(data, "metadata.uid")
				log.Printf("Uniq : %s", Uniq)
				url, _ := getDetailList(c, detailUrl, returnUrl2, Uniq.String());
				switch url {
					case "nf":
						return returnUrl
					case "em":
						return ""
					default:
						return url
				}
			}
	}

	return returnUrl
}

func GetKindURL(kindParam string) (url string, kind string) {
	var kindUrl string
	if kindParam == "deployments" || kindParam == "replicasets" {
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

	switch namespaceUrl {
		case "":
			namespaceUrl = namespaceParam + "/"
		case "\n":
			namespaceUrl = namespaceParam + "/"
		case "default":
			namespaceUrl = namespaceParam + "/"
		default:
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

	if multiParamLen >= 2 {
		switch kind_name {
			case "jobs":
				log.Printf("check Param : " + multiParamChk[1])
				if multiParamChk[1] == "list" {
					value = "jobs_podlist"
					return value
				}
			case "cronjobs":
				log.Printf("check Param : " + multiParamChk[1])
				if multiParamChk[1] == "list" {
					value = "cronjobs_joblist"
					return value
				}
			case "deployments":
				log.Printf("check Param : " + multiParamChk[1])
				if multiParamChk[1] == "list" {
					value = "deployments_list"
					return value
				}
			default:
				return "nf"
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

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
	client.SetTimeout(1 * time.Minute)
	client.SetHeaders(map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type": "application/json; charset=utf-8",
		"Accept": "application/json; charset=utf-8",
	})

	switch reqMethod {
		case "GET":
			if resp, err := client.R().SetAuthToken(token).Get(url); err != nil {
				panic(err)
			} else {
				responseString =string(resp.Body())
			}
		case "POST":
			if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Post(url); err != nil {
				panic(err)
			} else {
				responseString =string(resp.Body())
			}
		case "PATCH":
			if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Patch(url); err != nil {
				panic(err)
			} else {
				responseString =string(resp.Body())
			}
		case "PUT":
			if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Put(url); err != nil {
				panic(err)
			} else {
				responseString =string(resp.Body())
			}
		case "DELETE":
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

func getDetailList(c echo.Context, kind string, url string, uniq string) (urls string, err error) {

	jobList := []model.Job{}
	jobOnly := model.Job{}
	podList := []model.Pod{}
	// List := []model.Pod{}

	log.Printf("###getDetailURL Start###")
	log.Printf("val is %s", uniq)
	log.Printf("url is %s", url)
	if data := getData(c, url, false); data!="nf" {
		if uniq == "" || uniq == "\n" {
			return "nf", nil
	}

		switch kind {
			case "jobs_podlist":
				var returnVal string
				result := gjson.Get(data, `items.#.metadata.name`)
				for _, name := range result.Array() {
					if strings.Contains(name.String(), uniq) == true {
						returnVal = name.String()
					}
				}

				url = url + "/" + returnVal
				return url, nil

			case "cronjobs_joblist":
				k := gjson.Get(data, "items").Array()
				for t, _ := range k {
					arr := k[t].Get("metadata.ownerReferences").Array()
					if len(arr) > 0 {
						// log.Printf("[`1] %s, [`2] %d", arr, len(arr))
						for t2, _ := range arr {
							if strings.Contains(arr[t2].Get("kind").String(), "CronJob") == true && strings.Contains(arr[t2].Get("uid").String(), uniq) == true {
								// log.Printf("Same uid -> [1] %s, [2] %s", arr[t2].Get("uid").String(), uniq)
								// log.Printf("[34] %T, %s\n\n\n", k2[t], k2[t])
								err := json.Unmarshal([]byte(k[t].String()), &jobOnly)
								if err != nil {
									panic(err)
								}
								jobList = append(jobList, jobOnly)
							}
						}
					}
				}

				return "em", c.JSON(http.StatusOK, echo.Map{"items": jobList})
				

			case "deployments_list":
				k := gjson.Get(data, "items").Array()
				if chk, num, ok := testtttt(data, "metadata.ownerReferences", uniq); ok==true {
					log.Printf("[#1] %s, k[%s]", chk, k[num])
				}
				for t, _ := range k {
					arr := k[t].Get("metadata.ownerReferences").Array()
					if len(arr) > 0 {
						// log.Printf("[`1] %s, [`2] %d", arr, len(arr))
						for t2, _ := range arr {
							if arr[t2].Get("kind").String() == "Deployment" && arr[t2].Get("uid").String() == uniq {
								podList = getPods(c, "pods", k[t].Get("metadata").Get("uid").String())
								log.Printf("[`3] %s", k[t].Get("metadata").String())
								// endpointList = getEndpoints(c, "endpoints")
							}
						}
					}
				}

				return "em", c.JSON(http.StatusOK, echo.Map{
					"pods": podList,
					"services": "test",
				})				
		}
	}

	return "nf", nil
}

func SearchNested(obj interface{}, key string) (interface{}, bool) {
	switch t := obj.(type) {
	case map[string]interface{}:
			if v, ok := t[key]; ok {
					return v, ok
			}
			for _, v := range t {
					if result, ok := SearchNested(v, key); ok {
							return result, ok
					}
			}
	case []interface{}:
			for _, v := range t {
					if result, ok := SearchNested(v, key); ok {
							return result, ok
					}
			}
	}

	// key not found
	return nil, false
}



func getData(c echo.Context, url string, check bool) string {

	if data, err := HttpRequest(c, url, check); err != nil {
		log.Println("HttpRequest error")
		return "nf"
	} else {
		return data
	}
}

func getPods(c echo.Context, kind string, uniq string) []model.Pod {

	List := []model.Pod{}
	Only := model.Pod{}

	db := db.DbManager()
	models := FindClusterDB(db, "Name", c.Param("cluster_name"))
	if models == nil {
		panic(models)
	}

	kindUrl, kindName := GetKindURL(kind)
	url := "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(c.Param("namespace_name")) + kindName
	log.Printf("[53] url is %s", url)
					
	if data := getData(c, url, false); data!="nf" {
		n := gjson.Parse(data)
		// log.Printf("[20] %T, %s\n\n\n", n, n)
		k := n.Get("items").Array()
		for t, _ := range k {
			t2 := k[t].Get("metadata").Get("ownerReferences").Array()
			if len(t2) > 0 {
				// log.Printf("[`20] %s, [`21] %d", t2, len(t2))
				for t3, _ := range t2 {
					if t2[t3].Get("kind").String() == "ReplicaSet" && t2[t3].Get("uid").String() == uniq {
						// log.Printf("[22] Same uid -> [1] %s, [2] %s", t2[t3].Get("uid").String(), uniq)
						// log.Printf("[23] %s", k[t].Get("metadata").Get("uid"))
						// log.Printf("[34] %T, %s\n\n\n", k[t], k[t])
						err := json.Unmarshal([]byte(k[t].String()), &Only)
						if err != nil {
							panic(err)
						}
						List = append(List, Only)
					}
				}
			}
		}
	}
	return List
}

// func testtttt(data string, path string, chkData []gjson.Result) {
func testtttt(data string, path string, uniq string) (interface{}, int, bool) {

	var j interface{}
	log.Printf("\ndata is %s\n, path is %s", data, path)

	
	k := gjson.Get(data, "items").Array()
	for t, _ := range k {
		arr := k[t].Get(path).Array()
		if len(arr) > 0 {
			// log.Printf("[`1] %s, [`2] %d", arr, len(arr))
			for t2, _ := range arr {
			log.Printf("[`1] %s", arr[t2])

			err := json.Unmarshal([]byte(arr[t2].String()), &j)
			if err != nil {
				panic(err)
			}
		
			if v, ok := SearchNested(j, "kind"); ok {
				log.Printf("[#53] %+v\n", v)
				return v, t, true
			} else {
				log.Println("Key not found")
			}
				
			}
		}
	}
	return j, 0, false
}

func getEndpoints(c echo.Context, kind string, uniq string) []model.Endpoints {

	List := []model.Endpoints{}
	Only := model.Endpoints{}

	db := db.DbManager()
	models := FindClusterDB(db, "Name", c.Param("cluster_name"))
	if models == nil {
		panic(models)
	}

	kindUrl, kindName := GetKindURL(kind)
	url := "https://" + models.Endpoint + ":6443/" + kindUrl + GetNamespaceURL(c.Param("namespace_name")) + kindName
	log.Printf("[54] url is %s", url)
					
	if data := getData(c, url, false); data!="nf" {
		n := gjson.Parse(data)
		// log.Printf("[20] %T, %s\n\n\n", n, n)
		k := n.Get("items").Array()
		for t, _ := range k {
			t2 := k[t].Get("metadata").Get("ownerReferences").Array()
			if len(t2) > 0 {
				log.Printf("[`20] %s, [`21] %d", t2, len(t2))
				for t3, _ := range t2 {
					if t2[t3].Get("kind").String() == "ReplicaSet" && t2[t3].Get("uid").String() == uniq {
						log.Printf("[22] Same uid -> [1] %s, [2] %s", t2[t3].Get("uid").String(), uniq)
						log.Printf("[23] %s", k[t].Get("metadata").Get("uid"))
						// log.Printf("[34] %T, %s\n\n\n", k[t], k[t])
						err := json.Unmarshal([]byte(k[t].String()), &Only)
						if err != nil {
							panic(err)
						}
						List = append(List, Only)
					}
				}
			}
		}
	}
	return List
}