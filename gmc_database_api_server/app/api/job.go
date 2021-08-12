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
	infos := model.JOB{
		Name:         common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Workspace:    params.Workspace,
		Cluster:      params.Cluster,
		Project:      params.Project,
		BackoffLimit: StringToInt(common.InterfaceToString(common.FindData(getData, "spec", "backoffLimit"))),
		Completions:  StringToInt(common.InterfaceToString(common.FindData(getData, "spec", "completions"))),
		Parallelism:  StringToInt(common.InterfaceToString(common.FindData(getData, "spe", "parallelism"))),
	}
	// Name:           common.InterfaceToString(common.FindData(getData, "metadata", "name")),
	// Workspace: params.Workspace
	// Cluster:   params.Cluster
	// Project:   params.Project
	// Lable:       common.FindData(getData, "metadata", "labels"),
	// Annotations: common.FindData(getData, "metadata", "annotations"),
	// Kind:           common.InterfaceToString(common.FindData(getData, "kind", "")),
	// BackoffLimit: StringToInt(common.InterfaceToString(common.FindData(getData, "spec","backoffLimit")))
	// Completions:    common.InterfaceToString(common.FindData(getData, "spec", "completions")),
	// Parallelism:    common.InterfaceToInt(common.FindData(getData, "spe", "parallelism")),
	// CreationTime:   common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
	// OwnerReference: common.InterfaceToString(common.FindData(getData, "metadata", "ownerReferences")),
	// containers:     common.InterfaceToString(common.FindData(getData, "spec.template.spec", "containers")),
	// startTime:      common.InterfaceToTime(common.FindData(getData, "status", "startTime")),
	// completionTime: common.InterfaceToString(common.FindData(getData, "status", "completionTime")),
	// conditions:     common.FindData(getData, "status", "conditions"),
	// }

	return c.JSON(http.StatusOK, echo.Map{
		"jobDetails": infos,
	})
}
