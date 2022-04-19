package common

import (
	"encoding/json"
	"fmt"
	"gmc_api_gateway/app/model"
	"strings"

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
}

var nsTemplates = map[string]string{
	"pods":                   "/api/v1/namespaces/$1/pods/$2",
	"services":               "/api/v1/namespaces/$1/services/$2",
	"endpoints":              "/api/v1/namespaces/$1/endpoints/$2",
	"configmaps":             "/api/v1/namespaces/$1/configmaps/$2",
	"serviceaccounts":        "/api/v1/namespaces/$1/serviceaccounts/$2",
	"resourcequota":          "/api/v1/namespaces/$1/resourcequotas/$2",
	"deployments":            "/apis/apps/v1/namespaces/$1/deployments/$2",
	"replicasets":            "/apis/apps/v1/namespaces/$1/replicasets/$2",
	"daemonsets":             "/apis/apps/v1/namespaces/$1/daemonsets/$2",
	"statefulsets":           "/apis/apps/v1/namespaces/$1/statefulsets/$2",
	"jobs":                   "/apis/batch/v1/namespaces/$1/jobs/$2",
	"cronjobs":               "/apis/batch/v1/namespaces/$1/cronjobs/$2",
	"clusterroles":           "/apis/rbac.authorization.k8s.io/v1/clusterroles/$2",
	"roles":                  "/apis/rbac.authorization.k8s.io/v1/namespaces/$1/roles/$2",
	"clusterrolebindings":    "/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/$2",
	"namespaces":             "/api/v1/namespaces/$2",
	"nodes":                  "/api/v1/nodes/$2",
	"events":                 "/apis/events.k8s.io/v1/namespaces/$1/events/$2",
	"storageclasses":         "/apis/storage.k8s.io/v1/storageclasses/$2",
	"persistentvolumes":      "/api/v1/persistentvolumes/$2",
	"persistentvolumeclaims": "/api/v1/namespaces/$1/persistentvolumeclaims/$2",
	"secrets":                "/api/v1/namespaces/$1/secrets/$2",
}

var (
	SetCredential = new(model.Credential)
)

func DataRequest() (data string, err error) {
	var response string

	client := resty.New()
	resp, _ := client.R().
		EnableTrace().
		Get("http://210.207.104.188:1024/spider/credential")
	response = string(resp.Body())
	fmt.Println(StringToMapInterface(response))

	return response, nil
}

func DataFind(param model.PARAM) (data string, err error) {
	var response string
	method := param.Method

	client := resty.New()
	fmt.Println(param.Method)
	switch method {
	case "GET":
		if resp, err := client.R().EnableTrace().
			SetPathParams(map[string]string{
				"CredentialName": param.CredentialName,
			}).
			Get("http://210.207.104.188:1024/spider/credential/{CredentialName}"); err != nil {
			panic(err)
		} else {
			response = string(resp.Body())
		}
	case "POST":
		if resp, err := client.R().
			SetBody(map[string]interface{}{
				"CredentialName": param.CredentialName,
				"ProviderName":   "OPENSTACK",
				"KeyValueInfoList": []interface{}{
					map[string]interface{}{"Key": "IdentityEndpoint", "Value": param.IdentityEndPoint},
					map[string]interface{}{"Key": "Username", "Value": param.Username},
					map[string]interface{}{"Key": "Password", "Value": param.Password},
					map[string]interface{}{"Key": "DomainName", "Value": param.DomainName},
					map[string]interface{}{"Key": "ProjectID", "Value": param.ProjectID}},
			}).Post("http://210.207.104.188:1024/spider/credential"); err != nil {
			panic(err)
		} else {
			response = string(resp.Body())
		}
	case "DELETE":
		if resp, err := client.R().
			SetPathParams(map[string]string{
				"CredentialName": param.CredentialName,
			}).
			Delete("http://210.207.104.188:1024/spider/credential/{CredentialName}"); err != nil {
			panic(err)
		} else {
			response = string(resp.Body())
		}
	}
	return response, nil
}

func StringToMapInterface(i string) map[string]interface{} {
	x := make(map[string]interface{})
	if err := json.Unmarshal([]byte(i), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	return x
}

func UrlExpr(endpoint, project, item, kind string) string {
	check_project := strings.Compare(project, "") != 0
	check_item := strings.Compare(item, "") != 0

	// defaultUrl := "https://" + endpoint + ":6443"
	defaultUrl := "http://" + endpoint + ":8010"
	var returnUrl string
	fmt.Printf("##### teststest : %v,%v \n", check_project, check_item)
	if check_project || check_item {
		// project or item value exist
		// if err := errCheck(project, item, kind); err != "" {
		// 	return err
		// }
		returnUrl = defaultUrl + ProjectExpr(nsTemplates[kind], project, item)
	} else {
		returnUrl = defaultUrl + listTemplates[kind]
	}

	return returnUrl
}

func ProjectExpr(url, project, item string) string {
	check_project := strings.Compare(project, "") != 0
	check_item := strings.Compare(item, "") != 0
	returnVal := url

	if check_project && check_item {
		returnVal = strings.Replace(returnVal, "$1", project, -1)
		returnVal = strings.Replace(returnVal, "$2", item, -1)
	} else if check_project {
		returnVal = strings.Replace(returnVal, "$1", project, -1)
		returnVal = strings.Replace(returnVal, "$2", "", -1)
	} else if check_item {
		returnVal = strings.Replace(returnVal, "$2", item, -1)
	}

	return returnVal
}

func errCheck(project, item, kind string) string {
	check_project := strings.Compare(project, "") != 0
	check_item := strings.Compare(item, "") != 0

	if !check_project {
		if strings.Compare(kind, "clusterroles") == 0 || strings.Compare(kind, "namespaces") == 0 || strings.Compare(kind, "nodes") == 0 {
			if !check_item {
				return "nodetail"
			}
		} else {
			if !check_project {
				return "noname"
			}
		}
	}

	return ""
}

func validate(params model.PARAMS) error {
	workspaceCheck := strings.Compare(params.Workspace, "") != 0
	clusterCheck := strings.Compare(params.Cluster, "") != 0
	projectCheck := strings.Compare(params.Project, "") != 0
	nameCheck := strings.Compare(params.Name, "") != 0

	Method := params.Method
	// Body := responseBody(params.Body)
	Body := params.Body
	BodyCheck := strings.Compare(Body, "") != 0

	if Method == "POST" {
		if !BodyCheck {
			return ErrBodyEmpty
		}
		if !clusterCheck {
			return ErrClusterInvalid
		}
		if !projectCheck {
			return ErrProjectInvalid
		}
	} else if Method == "DELETE" {
		if !clusterCheck {
			return ErrClusterInvalid
		}
		if !projectCheck {
			return ErrProjectInvalid
		}
		if !nameCheck {
			return ErrDetailNameInvalid
		}
	} else if Method == "GET" {
		// if !clusterCheck {
		// 	return ErrClusterInvalid
		// }
	} else {
		if !clusterCheck {
			return ErrClusterInvalid
		}
		if !workspaceCheck {
			return ErrWorkspaceInvalid
		}
	}

	return nil
}
