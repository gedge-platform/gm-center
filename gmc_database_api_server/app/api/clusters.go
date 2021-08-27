package api

import (
	"fmt"
	"net/http"
	"strings"

	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/db"
	"gmc_database_api_server/app/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllClusters(c echo.Context) (err error) {
	db := db.DbManager()
	models := []model.Cluster{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
		return
	}
	fmt.Printf("[3##]models : %+v\n", models)
	return c.JSON(http.StatusOK, echo.Map{"data": models})
}
func GetAllClusters2(params model.PARAMS) []model.Cluster {
	db := db.DbManager()
	models := []model.Cluster{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		// common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
		return nil
	}
	fmt.Printf("[3##]models : %+v\n", models)
	return models
}
func GetCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := FindClusterDB(db, "Name", search_val)

	if models == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func GetCluster2(params model.PARAMS) *model.Cluster {
	search_val := params.Name
	db := db.DbManager()
	// search_val := c.Param("name")
	models := FindClusterDB(db, "Name", search_val)

	if models == nil {
		// common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}

	return models
}

func CreateCluster(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.Cluster)

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
	}
	if err = c.Validate(models); err != nil {
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return nil
	}

	if err != nil {
		panic(err)
	}

	if err := db.Create(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return nil
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": models})
}

func UpdateCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := model.Cluster{}

	if err := c.Bind(&models); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := FindClusterDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	} else {
		models.Name = search_val
	}

	models2 := FindClusterDB(db, "Name", search_val)

	if models.Ip != "" {
		models2.Ip = models.Ip
	}
	// if models.extIp != "" { models2.extIp = models.extIp }
	// if models.Role != "" {
	// 	models2.Role = models.Role
	// }
	if models.Type != "" {
		models2.Type = models.Type
	}
	if models.Endpoint != "" {
		models2.Endpoint = models.Endpoint
	}
	if models.Creator != "" {
		models2.Creator = models.Creator
	}
	// if models.Version != "" {
	// 	models2.Version = models.Version
	// }

	if err := db.Save(&models2).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models2})
}

func DeleteCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")

	if err := FindClusterDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, common.ErrNotFound)
		return nil
	}

	models := FindClusterDB(db, "Name", search_val)

	if err := db.Delete(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func FindClusterDB(db *gorm.DB, select_val string, search_val string) *model.Cluster {
	models := model.Cluster{}

	if check := strings.Compare(search_val, "") == 0; check {
		return nil
	}

	if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.Cluster{Name: search_val}).Error; err != nil {
			return nil
		}
	}
	return &models
}

func Get_Cluster(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.Param("name"),
		Workspace: c.Param("name"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
	var Cluster model.Node

	common.Transcode(getData0, &Cluster)
	clusterModel := GetCluster2(params)
	if clusterModel == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}
	gpuList, check := GpuCheck(params.Name)
	if check != false {
		clusterModel.Gpu = gpuList
	} else {
		clusterModel.Gpu = nil
	}
	clusterModel.Label = common.FindData(getData, "metadata", "labels")
	clusterModel.Annotation = common.FindData(getData, "metadata", "annotations")
	clusterModel.CreateAt = common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp"))
	clusterModel.Version = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kubeletVersion"))
	clusterModel.Os = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "operatingSystem"))
	clusterModel.Kernel = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kernelVersion"))
	clusterModel.Events = getCallEvent(params)
	// common.Transcode(getData0, &clusterModel)

	return c.JSON(http.StatusOK, echo.Map{
		"cluster": clusterModel,
		// "getData":  getData98,
	})

	// return nil
}

func Get_Clusters(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind: "nodes",
		// Name:      clusterModel[k].Name,
		// Cluster:   clusterModel[k].Name,
		// Workspace: clusterModel[k].Name,
		// Project:   clusterModel[k].Name,
		Method: c.Request().Method,
		Body:   c.Request().Body,
	}
	if c.QueryParam("workspace") == "" {
		clusterModel := GetAllClusters2(params)
		for k, _ := range clusterModel {
			fmt.Printf("value : %+v\n", clusterModel[k].Name)
			params.Name = clusterModel[k].Name
			params.Cluster = clusterModel[k].Name
			params.Workspace = clusterModel[k].Name
			params.Project = clusterModel[k].Name
			// params.Name = value.Name
			getData, err := common.GetModel(params)
			if err != nil {
				common.ErrorMsg(c, http.StatusNotFound, err)
				return nil
			}
			getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
			var Cluster model.Node
			common.Transcode(getData0, &Cluster)
			gpuList, check := GpuCheck(params.Name)
			if check != false {
				clusterModel[k].Gpu = gpuList
			} else {
				clusterModel[k].Gpu = nil
			}
			clusterModel[k].Label = common.FindData(getData, "metadata", "labels")
			clusterModel[k].Annotation = common.FindData(getData, "metadata", "annotations")
			clusterModel[k].CreateAt = common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp"))
			clusterModel[k].Version = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kubeletVersion"))
			clusterModel[k].Os = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "operatingSystem"))
			clusterModel[k].Kernel = common.InterfaceToString(common.FindData(getData, "status.nodeInfo", "kernelVersion"))
			tempMetric := []string{"cpu_usage", "memory_usage", "pod_running"}
			tempresult := NowMonit("cluster", params.Cluster, params.Name, tempMetric)
			clusterModel[k].ResourceUsage = tempresult
			// clusterModel[k].Kernel = "123"
		}
		return c.JSON(http.StatusOK, echo.Map{
			"clusters": clusterModel,
		})
	} else {
		var clusterModel []model.Cluster
		params.Workspace = c.QueryParam("workspace")
		workspace := GetWorkspace2(params)
		if workspace == nil {
			common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
			return
		}
		selectCluster := workspace.SelectCluster
		slice := strings.Split(selectCluster, ",")
		for i, _ := range slice {

			params.Name = slice[i]
			params.Cluster = slice[i]
			params.Project = slice[i]
			cluster := GetCluster2(params)
			var Cluster model.Cluster
			common.Transcode(cluster, &Cluster)
			getData, err := common.GetModel(params)
			if err != nil {
				common.ErrorMsg(c, http.StatusNotFound, err)
				return nil
			}
			fmt.Printf("[###data] : %s\n", getData)
			clusterModel = append(clusterModel, Cluster)

		}
		return c.JSON(http.StatusOK, echo.Map{
			"clusters": clusterModel,
		})
	}

}
