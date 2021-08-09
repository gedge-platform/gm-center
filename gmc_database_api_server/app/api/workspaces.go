package api

import (
	"log"
	"net/http"
	"strings"

	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/db"
	"gmc_database_api_server/app/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllWorkspaces(c echo.Context) (err error) {
	db := db.DbManager()
	models := []model.Workspace{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		common.ErrorMsg(c, http.StatusOK, common.ErrNoData)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func GetWorkspace(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := FindWorkspaceDB(db, "Name", search_val)

	if models == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func CreateWorkspace(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.Workspace)

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
	}
	if err = c.Validate(models); err != nil {
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
	}

	if err != nil {
		panic(err)
	}

	if err := db.Create(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": models})
}

func UpdateWorkspace(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := model.Workspace{}

	if err := c.Bind(&models); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := FindWorkspaceDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
	} else {
		models.Name = search_val
	}

	models2 := FindWorkspaceDB(db, "Name", search_val)

	if models.Description != "" {
		models2.Description = models.Description
	}
	if models.SelectCluster != "" {
		models2.SelectCluster = models.SelectCluster
	}
	if models.Owner != "" {
		models2.Owner = models.Owner
	}
	if models.Creator != "" {
		models2.Creator = models.Creator
	}

	if err := db.Save(&models2).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models2})
}

func DeleteWorkspace(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")

	if err := FindWorkspaceDB(db, "Name", search_val); err == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
	}

	models := FindWorkspaceDB(db, "Name", search_val)

	if err := db.Delete(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func FindWorkspaceDB(db *gorm.DB, select_val string, search_val string) *model.Workspace {
	models := model.Workspace{}
	if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.Workspace{Name: search_val}).Error; err != nil {
			return nil
		}
	}
	return &models
}

func WorkspaceData(c echo.Context, kind string) (string, string) {
	workspaceName := c.QueryParam("workspace")

	if err := WorkspaceValidate(workspaceName); err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return "", ""
	}

	Workspace, err := GetWorkspaceModel(c)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return "", ""
	}

	data, err := common.GetModel(c, kind)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return "", ""
	}

	return data, Workspace.SelectCluster
}

func WorkspaceValidate(name string) error {
	log.Println("name is ", name)
	if strings.Compare(name, "") == 0 {
		return common.ErrWorkspaceNotFound
	}
	return nil
}

func GetWorkspaceModel(c echo.Context) (*model.Workspace, error) {
	db := db.DbManager()
	workspaceName := c.QueryParam("workspace")
	models := FindWorkspaceDB(db, "Name", workspaceName)

	if models == nil {
		return models, common.ErrWorkspaceNotFound
	}

	return models, nil
}
