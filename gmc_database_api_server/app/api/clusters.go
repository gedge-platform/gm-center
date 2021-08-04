package api

import (
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

	return c.JSON(http.StatusOK, echo.Map{"data": models})
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
	if models.Role != "" {
		models2.Role = models.Role
	}
	if models.Type != "" {
		models2.Type = models.Type
	}
	if models.Endpoint != "" {
		models2.Endpoint = models.Endpoint
	}
	if models.Creator != "" {
		models2.Creator = models.Creator
	}
	if models.State != "" {
		models2.State = models.State
	}
	if models.Version != "" {
		models2.Version = models.Version
	}

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
