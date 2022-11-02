package controller

import (
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"

	// "log"
	"net/http"
	// "github.com/tidwall/sjson"
	"github.com/labstack/echo/v4"
)

// Get PV godoc
// @Summary Show List PV
// @Description get PV List
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Param workspace query string false "name of the Workspace"
// @Param cluster query string false "name of the Cluster"
// @Param project query string false "name of the Project"
// @Success 200 {object} model.PV
// @Router /pvs [get]
// @Tags Kubernetes
func GetAllPVs(c echo.Context) error {
	var pvs []model.PV
	params := model.PARAMS{
		Kind:      "persistentvolumes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		User:      c.QueryParam("user"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	data := GetModelList(params)

	for i, _ := range data {
		pv := model.PV{
			Name:          common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Capacity:      common.InterfaceToString(common.FindData(data[i], "spec", "capacity.storage")),
			AccessMode:    common.InterfaceToArray(common.FindData(data[i], "spec", "accessModes")),
			ReclaimPolicy: common.InterfaceToString(common.FindData(data[i], "spec", "persistentVolumeReclaimPolicy")),
			Status:        common.InterfaceToString(common.FindData(data[i], "status", "phase")),
			Claim:         common.FindData(data[i], "spec", "claimRef"),
			StorageClass:  common.InterfaceToString(common.FindData(data[i], "spec", "storageClassName")),
			VolumeMode:    common.InterfaceToString(common.FindData(data[i], "spec", "volumeMode")),
			Cluster:       common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			UserName:      common.InterfaceToString(common.FindData(data[i], "userName", "")),
			CreateAt:      common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
		}
		if params.User != "" {
			if params.User == pv.UserName {
				pvs = append(pvs, pv)
			}
		} else {
			pvs = append(pvs, pv)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": pvs,
	})
}

// Get PV godoc
// @Summary Show detail PV
// @Description get PV Details
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Param name path string true "name of the PV"
// @Param workspace query string true "name of the Workspace"
// @Param cluster query string true "name of the Cluster"
// @Param project query string true "name of the Project"
// @Success 200 {object} model.PV
// @Router /pvs/{name} [get]
// @Tags Kubernetes
func GetPV(c echo.Context) error {
	params := model.PARAMS{
		Kind:      "persistentvolumes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		User:      c.QueryParam("user"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}

	getData, err := common.DataRequest(params)
	// if err != nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, err)
	// 	return nil
	// }
	if err != nil || common.InterfaceToString(common.FindData(getData, "status", "")) == "Failure" {
		msg := common.ErrorMsg2(http.StatusNotFound, common.ErrNotFound)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": msg,
		})
	}
	pv := model.PV{
		Name:          common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Capacity:      common.InterfaceToString(common.FindData(getData, "spec", "capacity.storage")),
		AccessMode:    common.InterfaceToArray(common.FindData(getData, "spec", "accessModes")),
		ReclaimPolicy: common.InterfaceToString(common.FindData(getData, "spec", "persistentVolumeReclaimPolicy")),
		Status:        common.InterfaceToString(common.FindData(getData, "status", "phase")),
		Claim:         common.FindData(getData, "spec", "claimRef"),
		StorageClass:  common.InterfaceToString(common.FindData(getData, "spec", "storageClassName")),
		VolumeMode:    common.InterfaceToString(common.FindData(getData, "spec", "volumeMode")),
		Cluster:       c.QueryParam("cluster"),
		CreateAt:      common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		Lable:         common.FindData(getData, "metadata", "labels"),
		Annotations:   common.FindData(getData, "metadata", "annotations"),
		Events:        getCallEvent(params),
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": pv,
	})
}

// Create PV godoc
// @Summary Create PV
// @Description Create PV
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Param json body string true "PV Info Body"
// @Param cluster query string true "name of the Cluster"
// @Param workspace query string true "name of the Workspace"
// @Param project query string true "name of the Project"
// @Success 200 {object} model.Error
// @Router /pvs [post]
// @Tags Kubernetes
func CreatePV(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:    "persistentvolumes",
		Cluster: c.QueryParam("cluster"),
		Project: c.QueryParam("project"),
		Method:  c.Request().Method,
		Body:    responseBody(c.Request().Body),
	}
	params.Project = ""
	postData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "Created",
		"code":   http.StatusCreated,
		"data":   postData,
	})
}

// Delete PV godoc
// @Summary Delete PV
// @Description Delete PV
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Param name path string true "name of the PV"
// @Param workspace query string true "name of the Workspace"
// @Param cluster query string true "name of the Cluster"
// @Param project query string true "name of the Project"
// @Success 200 {object} model.Error
// @Router /pvs/{name} [delete]
// @Tags Kubernetes
func DeletePV(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:    "persistentvolumes",
		Name:    c.Param("name"),
		Cluster: c.QueryParam("cluster"),
		Project: c.QueryParam("project"),
		Method:  c.Request().Method,
		Body:    responseBody(c.Request().Body),
	}
	params.Project = ""
	postData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "Deleted",
		"code":   http.StatusOK,
		"data":   postData,
	})
}
