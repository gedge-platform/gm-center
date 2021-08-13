package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
		"jobDetails": jobinfos,
	})
}

// func GetAllJobs(c echo.Context) (err error) {

// 	params := model.PARAMS{
// 		Kind:      "jobs",
// 		Name:      "",
// 		Cluster:   c.QueryParam("cluster"),
// 		Workspace: c.QueryParam("workspace"),
// 		Project:   c.QueryParam("project"),
// 		Method:    c.Request().Method,
// 		Body:      c.Request().Body,
// 	}
// 	getData99, err := common.GetModel(params, "jobs")
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, err)
// 		return nil
// 	}
// 	getData98 := common.FindData(getData99, "", "")

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"jobList": getData98,
// 	})
// }
