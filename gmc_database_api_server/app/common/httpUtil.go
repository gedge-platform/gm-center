package common

import (
	"bytes"
	"crypto/tls"
	"gmc_database_api_server/app/db"
	"gmc_database_api_server/app/model"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

var listTemplates = map[string]string{
	"pods":                "/api/v1/pods",
	"services":            "/api/v1/services",
	"endpoints":           "/api/v1/endpoints",
	"configmaps":          "/api/v1/configmaps",
	"serviceaccounts":     "/api/v1/serviceaccounts",
	"resourcequota":       "/api/v1/resourcequotas",
	"deployments":         "/apis/apps/v1/deployments",
	"replicasets":         "/apis/apps/v1/replicasets",
	"jobs":                "/apis/batch/v1/jobs",
	"cronjobs":            "/apis/batch/v1/cronjobs",
	"clusterroles":        "/apis/rbac.authorization.k8s.io/v1/clusterroles",
	"roles":               "/apis/rbac.authorization.k8s.io/v1/roles",
	"clusterrolebindings": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings",
	"namespaces":          "/api/v1/namespaces",
	"nodes":               "/api/v1/nodes",
	"events":              "/apis/events.k8s.io/v1/events",
}

var nsTemplates = map[string]string{
	"pods":                "/api/v1/namespaces/$1/pods/$2",
	"services":            "/api/v1/namespaces/$1/services/$2",
	"endpoints":           "/api/v1/namespaces/$1/endpoints/$2",
	"configmaps":          "/api/v1/namespaces/$1/configmaps/$2",
	"serviceaccounts":     "/api/v1/namespaces/$1/serviceaccounts/$2",
	"resourcequota":       "/api/v1/namespaces/$1/resourcequotas/$2",
	"deployments":         "/apis/apps/v1/namespaces/$1/deployments/$2",
	"replicasets":         "/apis/apps/v1/namespaces/$1/replicasets/$2",
	"jobs":                "/apis/batch/v1/namespaces/$1/jobs/$2",
	"cronjobs":            "/apis/batch/v1/namespaces/$1/cronjobs/$2",
	"clusterroles":        "/apis/rbac.authorization.k8s.io/v1/clusterroles/$2",
	"roles":               "/apis/rbac.authorization.k8s.io/v1/namespaces/$1/roles/$2",
	"clusterrolebindings": "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/$2",
	"namespaces":          "/api/v1/namespaces/$2",
	"nodes":               "/api/v1/nodes/$2",
	"events":              "/apis/events.k8s.io/v1/namespaces/$1/events/$2",
}

var (
	SetPod                = new(model.Pod)
	SetService            = new(model.Service)
	SetEndpoint           = new(model.Endpoints)
	SetConfigmap          = new(model.ConfigMap)
	SetServiceaccount     = new(model.ServiceAccount)
	SetResourcequota      = new(model.ResourceQuota)
	SetDeployment         = new(model.Deployment)
	SetReplicaset         = new(model.ReplicaSet)
	SetJob                = new(model.Job)
	SetCronjob            = new(model.CronJob)
	SetClusterrole        = new(model.ClusterRole)
	SetRole               = new(model.Role)
	SetClusterrolebinding = new(model.ClusterRoleBinding)
	SetNamespace          = new(model.Namespace)
	SetNode               = new(model.Node)
	SetEvent              = new(model.Event)

	SetPods                = new(model.PodList)
	SetServices            = new(model.ServiceList)
	SetEndpoints           = new(model.EndpointsList)
	SetConfigmaps          = new(model.ConfigMapList)
	SetServiceaccounts     = new(model.ServiceAccountList)
	SetResourcequotas      = new(model.ResourceQuotaList)
	SetDeployments         = new(model.DeploymentList)
	SetReplicasets         = new(model.ReplicaSetList)
	SetJobs                = new(model.JobList)
	SetCronjobs            = new(model.CronJobList)
	SetClusterroles        = new(model.ClusterRoleList)
	SetRoles               = new(model.RoleList)
	SetClusterrolebindings = new(model.ClusterRoleBindingList)
	SetNamespaces          = new(model.NamespaceList)
	SetNodes               = new(model.NodeList)
	SetEvents              = new(model.EventList)
)

func GetModel(c echo.Context, kind string) (data string, err error) {
	var endPoint, namespace_name, item_name string

	if err := validate(c); err != nil {
		return "", err
	}

	if data, err := FindClusterDB(c.QueryParam("cluster")); err != nil {
		return "", err
	} else {
		endPoint = data.Endpoint
	}
	if strings.Compare(c.QueryParam("namespace"), "") != 0 {
		namespace_name = c.QueryParam("namespace")
	}

	if strings.Compare(c.Param("name"), "") != 0 {
		item_name = c.Param("name")
	}

	// models := ReturnModel(c.Param("name"), kind)
	url := UrlExpr(endPoint, namespace_name, item_name, kind)

	log.Println("url is", url)

	switch url {
	case "noname":
		return "", ErrNamespaceInvalid
	case "nodetail":
		return "", ErrDetailNameInvalid
	}

	log.Printf("[#31] url is %s", url)
	var responseString, token string
	reqMethod := c.Request().Method
	passBody := responseBody(c.Request().Body)

	tokens, ok := c.Request().Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
	}

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(1 * time.Minute)
	client.SetHeaders(map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type":                "application/json; charset=utf-8",
		"Accept":                      "application/json; charset=utf-8",
	})

	switch reqMethod {
	case "GET":
		if resp, err := client.R().SetAuthToken(token).Get(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
		}
	case "POST":
		if resp, err := client.R().SetBody([]byte(string(passBody))).SetAuthToken(token).Post(url); err != nil {
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
		if resp, err := client.R().SetAuthToken(token).Delete(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
		}
	}

	// if err := json.Unmarshal([]byte(responseString), &models); err != nil {
	// 	fmt.Println("옷이 맞지 않음")
	// } else {
	// 	return models, nil
	// }

	// x := make(map[string]interface{})
	// if err := json.Unmarshal([]byte(responseString), &x); err != nil {
	// 	fmt.Printf("Error : %s\n", err)
	// }
	// return x, nil

	return responseString, nil
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

func FindClusterDB(name string) (*model.Cluster, error) {
	log.Println("in FindClusterDB")
	db := db.DbManager()
	var models model.Cluster

	if err := db.First(&models, model.Cluster{Name: name}).Error; err != nil {
		return &models, err
	}
	return &models, nil
}

func ReturnModel(detail_name string, kind string) interface{} {
	var models interface{}
	check := strings.Compare(detail_name, "") != 0

	switch kind {
	case "pods":
		if check {
			models = SetPod
		} else {
			models = SetPods
		}
	case "services":
		if check {
			models = SetService
		} else {
			models = SetServices
		}
	case "endpoints":
		if check {
			models = SetEndpoint
		} else {
			models = SetEndpoints
		}
	case "configmaps":
		if check {
			models = SetConfigmap
		} else {
			models = SetConfigmaps
		}
	case "serviceaccounts":
		if check {
			models = SetServiceaccount
		} else {
			models = SetServiceaccounts
		}
	case "resourcequota":
		if check {
			models = SetResourcequota
		} else {
			models = SetResourcequotas
		}
	case "deployments":
		if check {
			models = SetDeployment
		} else {
			models = SetDeployments
		}
	case "replicasets":
		if check {
			models = SetReplicaset
		} else {
			models = SetReplicasets
		}
	case "jobs":
		if check {
			models = SetJob
		} else {
			models = SetJobs
		}
	case "cronjobs":
		if check {
			models = SetCronjob
		} else {
			models = SetCronjobs
		}
	case "clusterroles":
		if check {
			models = SetClusterrole
		} else {
			models = SetClusterroles
		}
	case "roles":
		if check {
			models = SetRole
		} else {
			models = SetRoles
		}
	case "clusterrolebindings":
		if check {
			models = SetClusterrolebinding
		} else {
			models = SetClusterrolebindings
		}
	case "namespaces":
		if check {
			models = SetNamespace
		} else {
			models = SetNamespaces
		}
	case "nodes":
		if check {
			models = SetNode
		} else {
			models = SetNodes
		}
	case "events":
		if check {
			models = SetEvent
		} else {
			models = SetEvents
		}
	}

	return models
}

func UrlExpr(endpoint, namespace, item, kind string) string {
	check_namespace := strings.Compare(namespace, "") != 0
	check_item := strings.Compare(item, "") != 0

	defaultUrl := "https://" + endpoint + ":6443"
	var returnUrl string

	if check_namespace || check_item {
		// namespace or item value exist
		if err := errCheck(namespace, item, kind); err != "" {
			return err
		}
		returnUrl = defaultUrl + NamespaceExpr(nsTemplates[kind], namespace, item)
	} else {
		returnUrl = defaultUrl + listTemplates[kind]
	}

	return returnUrl
}

func NamespaceExpr(url, namespace, item string) string {
	log.Println("items is ", item)
	check_namespace := strings.Compare(namespace, "") != 0
	check_item := strings.Compare(item, "") != 0
	returnVal := url

	if check_namespace && check_item {
		returnVal = strings.Replace(returnVal, "$1", namespace, -1)
		returnVal = strings.Replace(returnVal, "$2", item, -1)
	} else if check_namespace {
		returnVal = strings.Replace(returnVal, "$1", namespace, -1)
		returnVal = strings.Replace(returnVal, "$2", "", -1)
	} else if check_item {
		returnVal = strings.Replace(returnVal, "$2", item, -1)
	}

	return returnVal
}

func errCheck(namespace, item, kind string) string {
	check_namespace := strings.Compare(namespace, "") != 0
	check_item := strings.Compare(item, "") != 0

	if !check_namespace {
		if strings.Compare(kind, "clusterroles") == 0 || strings.Compare(kind, "namespaces") == 0 || strings.Compare(kind, "nodes") == 0 {
			if !check_item {
				return "nodetail"
			}
		} else {
			if !check_namespace {
				return "noname"
			}
		}
	}

	return ""
}

func validate(c echo.Context) error {
	cluster_name := c.QueryParam("cluster")
	if strings.Compare(cluster_name, "") == 0 {
		return ErrClusterInvalid
	}
	return nil
}
