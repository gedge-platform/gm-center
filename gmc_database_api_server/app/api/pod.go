package api

import (
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetPods godoc
// @Summary Show detail pods
// @Description get pods Details
// @Accept  json
// @Produce  json
// @Success 200 {object} model.POD
// @Header 200 {string} Token "qwerty"
// @Router /pods/:name [get]
func GetPods(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "pods",
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

	ownerReferencesData := common.FindData(getData, "metadata", "ownerReferences")
	var ownerReferencesInfo []model.OwnerReference
	common.Transcode(ownerReferencesData, &ownerReferencesInfo)

	podIPsData := common.FindData(getData, "status", "podIPs")
	var podIPsInfo []model.PodIPs
	common.Transcode(podIPsData, &podIPsInfo)

	containerStatusesData := common.FindData(getData, "status", "containerStatuses")
	var containerStatusesInfo []model.ContainerStatuses
	common.Transcode(containerStatusesData, &containerStatusesInfo)

	podcontainersData := common.FindData(getData, "spec", "containers")
	var podcontainersDataInfo []model.PODCONTAINERS
	common.Transcode(podcontainersData, &podcontainersDataInfo)

	// volumeMountsData := common.FindData(getData, "spec.containers", "volumeMounts")
	// var volumeMountsInfo []model.VolumeMounts
	// common.Transcode(volumeMountsData, &volumeMountsInfo)
	referData, _ := common.GetModelRelatedList(params)
	log.Printf("#####getdata99 ", referData)

	podinfos := model.POD{
		Workspace:         params.Workspace,
		Cluster:           params.Cluster,
		Project:           params.Project,
		Name:              common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Namespace:         common.InterfaceToString(common.FindData(getData, "metadata", "namespace")),
		CreationTimestamp: common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		NodeName:          common.InterfaceToString(common.FindData(getData, "spec", "nodeName")),
		Lable:             common.FindData(getData, "metadata", "labels"),
		Annotations:       common.FindData(getData, "metadata", "annotations"),
		QosClass:          common.InterfaceToString(common.FindData(getData, "status", "qosClass")),
		OwnerReference:    ownerReferencesInfo,
		Status:            common.InterfaceToString(common.FindData(getData, "status", "phase")),
		HostIP:            common.InterfaceToString(common.FindData(getData, "status", "hostIP")),
		PodIP:             common.InterfaceToString(common.FindData(getData, "status", "podIP")),
		PodIPs:            podIPsInfo,
		ContainerStatuses: containerStatusesInfo,
		Podcontainers:     podcontainersDataInfo,
		// VolumeMounts:      volumeMountsInfo,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"podDetail": podinfos,
		"referData": referData,
	})
}
