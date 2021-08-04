package api

import (
	"encoding/json"
	"fmt"
	"gmc_database_api_server/app/common"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func Get_Cluster(c echo.Context) (err error) {
	KubernetesNODE, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/test1/nodes", false)
	clusterModel := GetCluster2(c)
	k := gjson.Get(KubernetesNODE, "items").Array()
	label := make(map[string]string)
	annotation := make(map[string]string)
	clusterRole := ""
	// createAt := time.Time{}
	for t, _ := range k {
		labels := k[t].Get("metadata.labels")
		err_label := json.Unmarshal([]byte(labels.String()), &label)
		if err_label != nil {
			fmt.Printf("Error : %s\n", err_label)
		}
		for key, _ := range label {
			if key == "node-role.kubernetes.io/master" {
				clusterRole = "master"
			}
			// fmt.Printf("%s\n", key)
		}
		clusterModel.Lable = label
		annotations := k[t].Get("metadata.annotations")
		err_annotation := json.Unmarshal([]byte(annotations.String()), &annotation)
		if err_annotation != nil {
			fmt.Printf("Error : %s\n", err_annotation)
		}
		clusterModel.Annotation = annotation
		clusterModel.Name = k[t].Get("metadata.name").String()
		createTime := k[t].Get("metadata.creationTimestamp").Time()
		clusterModel.Version = k[t].Get("status.nodeInfo.kubeletVersion").String()
		clusterModel.Os = k[t].Get("status.nodeInfo.operatingSystem").String() + " / " + k[t].Get("status.nodeInfo.osImage").String()
		clusterModel.Kernel = k[t].Get("status.nodeInfo.kernelVersion").String()
		clusterModel.Role = clusterRole

		// createAt := createTime.Format("2006-01-02 15:04:05") // format 변경은 완료, time.Time type X

		clusterModel.CreateAt = createTime
		// fmt.Printf("version : %s\n", k[t].Get("status.nodeInfo.kubeletVersion").String())
		// return c.JSON(http.StatusOK, echo.Map{"data": clusterModel})
		// return clusterModel
	}
	// return err
	return c.JSON(http.StatusOK, echo.Map{"data": clusterModel})
}
func Get_Project(c echo.Context) (err error) {

	data, err := common.GetModel(c, "pods")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	return c.JSON(http.StatusOK, echo.Map{"data": data})

	// KubernetesNS, _ := HttpRequest(c, "https://g-api-test.innogrid.tech/kube/v1/test1/namespaces", false)
	// projectModel := GetAllProjects(c)

	// fmt.Printf("test : %+v\n", projectModel)

	// k := gjson.Get(KubernetesNS, "items").Array()
	// label := make(map[string]string)
	// annotation := make(map[string]string)
	// clusterRole := ""
	// createAt := time.Time{}
	// for t, _ := range k {
	// 	labels := k[t].Get("metadata.labels")
	// 	err_label := json.Unmarshal([]byte(labels.String()), &label)
	// 	if err_label != nil {
	// 		fmt.Printf("Error : %s\n", err_label)
	// 	}
	// 	// projectModel.Lable = label
	// 	annotations := k[t].Get("metadata.annotations")
	// 	err_annotation := json.Unmarshal([]byte(annotations.String()), &annotation)
	// 	if err_annotation != nil {
	// 		fmt.Printf("Error : %s\n", err_annotation)
	// 	}
	// projectModel.Annotation = annotation

	// createAt := createTime.Format("2006-01-02 15:04:05") // format 변경은 완료, time.Time type X

	// projectModel.CreateAt = createTime
	// fmt.Printf("version : %s\n", k[t].Get("status.nodeInfo.kubeletVersion").String())
	// return c.JSON(http.StatusOK, echo.Map{"data": clusterModel})
	// return clusterModel
	// }
	return
	// return c.JSON(http.StatusOK, echo.Map{"data": projectModel})
}
