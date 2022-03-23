package api

import (
	"fmt"
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	// "log"
	"net/http"
	// "github.com/tidwall/sjson"
	"github.com/labstack/echo/v4"
)

// GetPvs godoc
// @Summary Show app PVs
// @Description get pv List
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PV
// @Header 200 {string} Token "qwerty"
// @Router /pvs [get]
func GetAllPVs(c echo.Context) error {
	var pvs []model.PV
	fmt.Printf("## PVs", pvs)
	params := model.PARAMS{
		Kind:      "persistentvolumes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	data := GetModelList(params)
	fmt.Printf("####Pod data confirm : %s", data)

	for i, _ := range data {
		pv := model.PV{
			Name:              common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Capacity:         common.InterfaceToString(common.FindData(data[i], "spec", "capacity.storage")),
			AccessMode:           common.InterfaceToArray(common.FindData(data[i], "spec", "accessModes")),
			ReclaimPolicy: common.InterfaceToString(common.FindData(data[i], "spec", "persistentVolumeReclaimPolicy")),
			Status:            common.InterfaceToString(common.FindData(data[i], "status", "phase")),
			Claim:          common.FindData(data[i], "spec", "claimRef"),
			StorageClass:             common.InterfaceToString(common.FindData(data[i], "spec", "storageClassName")),
			VolumeMode :            common.InterfaceToString(common.FindData(data[i], "spec", "volumeMode")),
			Cluster: common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			CreateAt : common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
		}
		pvs = append(pvs, pv)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": pvs,
	})
}

// GetPvs godoc
// @Summary Show detail PVs
// @Description get PVs Details
// @Accept  json
// @Produce  json
// @Success 200 {object} model.POD
// @Header 200 {string} Token "qwerty"
// @Router /pvs/:name [get]
func GetPV(c echo.Context) error {
	var pvs []model.PV
	fmt.Printf("## PVs", pvs)
	params := model.PARAMS{
		Kind:      "persistentvolumes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	fmt.Printf("####PV data confirm : %s", getData)
	pv := model.PV{
				Name:              common.InterfaceToString(common.FindData(getData, "metadata", "name")),
				Capacity:         common.InterfaceToString(common.FindData(getData, "spec", "capacity.storage")),
				AccessMode:           common.InterfaceToArray(common.FindData(getData, "spec", "accessModes")),
				ReclaimPolicy: common.InterfaceToString(common.FindData(getData, "spec", "persistentVolumeReclaimPolicy")),
				Status:            common.InterfaceToString(common.FindData(getData, "status", "phase")),
				Claim:          common.FindData(getData, "spec", "claimRef"),
				StorageClass:             common.InterfaceToString(common.FindData(getData, "spec", "storageClassName")),
				VolumeMode :            common.InterfaceToString(common.FindData(getData, "spec", "volumeMode")),
				Cluster: c.QueryParam("cluster"),
				CreateAt : common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
				Lable:             common.FindData(getData, "metadata", "labels"),
		Annotations:       common.FindData(getData, "metadata", "annotations"),
				Events: getCallEvent(params),
			}

	return c.JSON(http.StatusOK, echo.Map{
		"data": pv,
	})
}
