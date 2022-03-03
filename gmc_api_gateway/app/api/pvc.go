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
// @Summary Show app PVCs
// @Description get pvc List
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PVC
// @Header 200 {string} Token "qwerty"
// @Router /pvcs [get]
func GetAllPVCs(c echo.Context) error {
	var pvcs []model.PVC
	fmt.Printf("## PVCs", pvcs)
	params := model.PARAMS{
		Kind:      "persistentvolumeclaims",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	data := GetModelList(params)
	fmt.Printf("####Pod data confirm : %s", data)
	// Name        string           `json:"name"`
	// Capacity   string           `json:"capacity"`
	// AccessMode      []string `json:"accessMode"`
	// Status    interface{}      `json:"status"`
	// Volume   interface{}        `json:"volume"`
	// StorageClass       string           `json:"storageClass"`
	// // Reason        []EVENT          `json:"events"`
	// CreateAt time.Time          `json:"createAt"`
	// Events  []EVENT          `json:"events"`
	for i, _ := range data {
		pvc := model.PVC{
			Name:              common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Capacity:         common.InterfaceToString(common.FindData(data[i], "spec.resources", "requests.storage")),
			AccessMode:           common.InterfaceToArray(common.FindData(data[i], "spec", "accessModes")),
			Status:            common.InterfaceToString(common.FindData(data[i], "status", "phase")),
			StorageClass:             common.InterfaceToString(common.FindData(data[i], "spec", "storageClassName")),
			Volume :            common.InterfaceToString(common.FindData(data[i], "spec", "volumeName")),
			Cluster: common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			CreateAt : common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
		}
		pvcs = append(pvcs, pvc)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": pvcs,
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
func GetPVC(c echo.Context) error {
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
				Events: getCallEvent(params),
			}

	return c.JSON(http.StatusOK, echo.Map{
		"data": pv,
	})
}

