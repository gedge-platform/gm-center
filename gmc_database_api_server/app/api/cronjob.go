package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCronJobs(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:      "cronjobs",
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
	containerData := common.FindData(getData, "spec.jobTemplate.spec.template.spec", "containers")
	var containerInfo []model.Containers
	common.Transcode(containerData, &containerInfo)

	activeData := common.FindData(getData, "status", "active")
	var activeInfo []model.Active
	common.Transcode(activeData, &activeInfo)

	cronjobInfos := model.CRONJOB{
		Workspace:                  params.Workspace,
		Cluster:                    params.Cluster,
		Project:                    params.Project,
		Name:                       common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Lable:                      common.FindData(getData, "metadata", "labels"),
		Annotations:                common.FindData(getData, "metadata", "annotations"),
		Schedule:                   common.InterfaceToString(common.FindData(getData, "spec", "schedule")),
		ConcurrencyPolicy:          common.InterfaceToString(common.FindData(getData, "spec", "concurrencyPolicy")),
		SuccessfulJobsHistoryLimit: StringToInt(common.InterfaceToString(common.FindData(getData, "spec", "successfulJobsHistoryLimit"))),
		FailedJobsHistoryLimit:     StringToInt(common.InterfaceToString(common.FindData(getData, "spec", "failedJobsHistoryLimits"))),
		LastScheduleTime:           common.InterfaceToTime(common.FindData(getData, "status", "lastScheduleTime")),
		CreationTimestamp:          common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		Containers:                 containerInfo,
		Active:                     activeInfo,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"cronjobDetail": cronjobInfos,
	})
}

// func GetAllCronJobs(c echo.Context) (err error) {

// 	params := model.PARAMS{
// 		Kind:      "cronjobs",
// 		Name:      "",
// 		Cluster:   c.QueryParam("cluster"),
// 		Workspace: c.QueryParam("workspace"),
// 		Project:   c.QueryParam("project"),
// 		Method:    c.Request().Method,
// 		Body:      c.Request().Body,
// 	}
// 	getData99, err := common.GetModel(params, "cronjobs")
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, err)
// 		return nil
// 	}
// 	getData98 := common.FindData(getData99, "", "")
// 	// cronjobInfos := []model.CRONJOB{
// 	// 	getData98
// 	// }
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"cronjobList": getData98,
// 	})
// }
