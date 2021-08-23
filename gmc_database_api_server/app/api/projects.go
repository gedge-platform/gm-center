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

func GetAllProjects(c echo.Context) (err error) {
	db := db.DbManager()
	models := []model.Project{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
		return
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}
func GetAllProjects2(c echo.Context) []model.Project {
	db := db.DbManager()
	models := []model.Project{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		common.ErrorMsg(c, http.StatusOK, common.ErrNoData)

	}

	return models
}
func GetProject2(c echo.Context) *model.Project {
	db := db.DbManager()
	search_val := c.Param("name")
	models := FindProjectDB(db, "Name", search_val)

	if models == nil {
		// common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		var model model.Project
		model.Type = "system"
		return &model
	}

	return models
}
func GetProject3(params model.PARAMS) *model.Project {
	db := db.DbManager()
	search_val := params.Name
	models := FindProjectDB(db, "Name", search_val)

	if models == nil {
		// common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		var model model.Project
		model.Type = "system"
		model.WorkspaceName = "system"
		model.SelectCluster = params.Cluster
		return &model
	}

	return models
}

func GetProject(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := FindProjectDB(db, "Name", search_val)

	if models == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func CreateProject(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.Project)

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

func UpdateProject(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := model.Project{}

	if err := c.Bind(&models); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := FindProjectDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	} else {
		models.Name = search_val
	}

	models2 := FindProjectDB(db, "Name", search_val)

	if models.Postfix != "" {
		models2.Postfix = models.Postfix
	}
	if models.Description != "" {
		models2.Description = models.Description
	}
	if models.Type != "" {
		models2.Type = models.Type
	}
	if models.Owner != "" {
		models2.Owner = models.Owner
	}
	if models.Creator != "" {
		models2.Creator = models.Creator
	}
	if models.WorkspaceName != "" {
		models2.WorkspaceName = models.WorkspaceName
	}

	if err := db.Save(&models2).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models2})
}

func DeleteProject(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")

	if err := FindProjectDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}

	models := FindProjectDB(db, "Name", search_val)

	if err := db.Delete(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func FindProjectDB(db *gorm.DB, select_val string, search_val string) *model.Project {
	models := model.Project{}
	if check := strings.Compare(search_val, "") == 0; check {
		return nil
	}
	if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.Project{Name: search_val}).Error; err != nil {
			return nil
		}
	}
	return &models
}
