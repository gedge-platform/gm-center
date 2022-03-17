package api

import (
	"fmt"
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Router /pvs/:name [get]
func GetStorageclass(c echo.Context) error {
	var storageclasses []model.STORAGECLASS
	fmt.Printf("## STORAGECLASS", storageclasses)
	params := model.PARAMS{
		Kind:      "storageclasses",
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

	fmt.Println("[###########storageclass]", getData)
	storageclass := model.STORAGECLASS{
		Name:                 common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Cluster:  common.InterfaceToString(common.FindData(getData, "clusterName", "")),
		ReclaimPolicy:        common.InterfaceToString(common.FindData(getData, "reclaimPolicy", "")),
		Provisioner:          common.InterfaceToString(common.FindData(getData, "provisioner", "")),
		VolumeBindingMode:    common.InterfaceToString(common.FindData(getData, "volumeBindingMode", "")),
		AllowVolumeExpansion: common.InterfaceToString(common.FindData(getData, "allowVolumeExpansion", "")),
		Parameters : common.FindData(getData, "parameters", ""),
		CreateAt:             common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		Labels :    common.FindData(getData, "metadata", "labels"),
		Annotations:          common.FindData(getData, "metadata", "annotations"),
		//Age:                  common.InterfaceToString(common.FindData(getData, "age", "")),
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": storageclass,
	})
}

func GetStorageclasses(c echo.Context) (err error) {
	var storageclasses []model.STORAGECLASS
	fmt.Printf("## storageclasses", storageclasses)
	params := model.PARAMS{
		Kind:      "storageclasses",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	data := GetModelList(params)
	fmt.Printf("####storageclass data confirm : %s", data)

	for i, _ := range data {
		var allowVolumeExpansion string
		if common.InterfaceToString(common.FindData(data[i], "allowVolumeExpansion", "")) != "" {
			allowVolumeExpansion = common.InterfaceToString(common.FindData(data[i], "allowVolumeExpansion", ""))
		}else {
			allowVolumeExpansion = "false"
		}
		storageclass := model.STORAGECLASS{
			Name:                 common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Cluster:  common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			ReclaimPolicy:        common.InterfaceToString(common.FindData(data[i], "reclaimPolicy", "")),
			Provisioner:          common.InterfaceToString(common.FindData(data[i], "provisioner", "")),
			VolumeBindingMode:    common.InterfaceToString(common.FindData(data[i], "volumeBindingMode", "")),
			AllowVolumeExpansion: allowVolumeExpansion,
			CreateAt:             common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
		}
		storageclasses = append(storageclasses, storageclass)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": storageclasses,
	})
}
