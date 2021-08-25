package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetJobs godoc
// @Summary Show detail job
// @Description get job Details
// @Accept  json
// @Produce  json
// @Success 200 {object} model.JOB
// @Header 200 {string} Token "qwerty"
// @Router /jobs/:name [get]
func GetJobs(c echo.Context) error {
	params := model.PARAMS{
		Kind:      "jobs",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}

	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	containerData := common.FindData(getData, "spec.template.spec", "containers")
	var containerInfo []model.Containers
	common.Transcode(containerData, &containerInfo)

	conditionData := common.FindData(getData, "status", "conditions")
	var conditionInfo []model.Conditions
	common.Transcode(conditionData, &conditionInfo)

	ownerReferencesData := common.FindData(getData, "metadata", "ownerReferences")
	var ownerReferencesInfo []model.OwnerReference
	common.Transcode(ownerReferencesData, &ownerReferencesInfo)

	referData, _ := common.GetModelRelatedList(params)
	log.Printf("#####getdata99 ", referData)

	jobinfos := model.JOB{
		Workspace:      params.Workspace,
		Cluster:        params.Cluster,
		Project:        params.Project,
		Name:           common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Lable:          common.FindData(getData, "metadata", "labels"),
		Annotations:    common.FindData(getData, "metadata", "annotations"),
		Kind:           common.InterfaceToString(common.FindData(getData, "kind", "")),
		OwnerReference: ownerReferencesInfo,
		BackoffLimit:   StringToInt(common.InterfaceToString(common.FindData(getData, "spec", "backoffLimit"))),
		Completions:    StringToInt(common.InterfaceToString(common.FindData(getData, "spec", "completions"))),
		Parallelism:    StringToInt(common.InterfaceToString(common.FindData(getData, "spe", "parallelism"))),
		Status:         StringToInt(common.InterfaceToString(common.FindData(getData, "status", "succeeded"))),
		CreationTime:   common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		StartTime:      common.InterfaceToTime(common.FindData(getData, "status", "startTime")),
		CompletionTime: common.InterfaceToTime(common.FindData(getData, "status", "completionTime")),
		Conditions:     conditionInfo,
		Containers:     containerInfo,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"jobDetail": jobinfos,
		"referData": referData,
	})
}

// func GetAllJobs(c echo.Context) error {
// 	params := model.PARAMSAll{
// 		Kind: "jobs",
// 		Name: "",
// 		// Cluster: c.QueryParam("cluster"),
// 		// Workspace: c.QueryParam("workspace"),
// 		Project: c.QueryParam("project"),
// 		Method:  c.Request().Method,
// 		Body:    c.Request().Body,
// 	}
// 	getData, err := common.GetModel2(params)
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, err)
// 		return nil
// 	}
// 	log.Printf("testing data ", getData)
// 	testing, err := common.GetModel2(params)
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, err)
// 		return nil
// 	}
// 	// log.Printf("[#1111]", common.SetJobs)
// 	// GetAllClusters(c)
// 	// GetCluster2(c)
// 	log.Printf("###11111", GetAllClusters(c))
// 	// log.Printf("###22222", GetCluster2(c))
// 	// ClusterAll

// 	List := []model.JOBALL{}
// 	Only := model.JOBALL{}
// 	n := gjson.Parse(testing)
// 	log.Printf("testing data ", n)
// 	k := n.Get("items").Array()
// 	log.Printf("testing data items  ", k)

// 	for num, _ := range k {

// 		test := k[num].String()
// 		log.Printf("testing data only items  ", test)
// 		err := json.Unmarshal([]byte(test), &Only)
// 		if err != nil {
// 			fmt.Println("[!53] error")
// 		}
// 		List = append(List, Only)

// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"allData": List,
// 		// "test":    testing,
// 	})
// }
