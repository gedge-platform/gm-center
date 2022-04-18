package api

import (
	"fmt"
	"net/http"
	"strings"

	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"

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
	params := model.PARAMS{
		// Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	db := db.DbManager()
	search_val := params.Name
	models := FindWorkspaceDB(db, "Name", search_val)

	if models == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}
	var workspaceDetail model.Workspace_detail
	var projectList []model.Workspace_project
	var tsWorkspace model.Workspace
	common.Transcode(models, &tsWorkspace)
	workspaceDetail.Workspace = tsWorkspace
	params.Workspace = params.Name
	projects := GetWorkspaceFilterProjects(params)
	for p := range projects {
		fmt.Println("[####projects] : ", projects[p])
		project := model.Workspace_project{
			Name:          projects[p].Name,
			SelectCluster: projects[p].SelectCluster,
			CreateAt:      projects[p].CreateAt,
			Creator:       projects[p].Creator,
		}
		projectList = append(projectList, project)
	}
	Resource, Usage, Events := GetUserProjectResource(c, projects)
	workspaceDetail.ProjectList = projectList
	workspaceDetail.Resource = Resource
	workspaceDetail.ResourceUsage = Usage
	workspaceDetail.Events = Events

	fmt.Printf("selectClusters : %+v", workspaceDetail.Workspace.SelectCluster)
	return c.JSON(http.StatusOK, echo.Map{
		"data": workspaceDetail,
	})
}

func GetDBWorkspace(params model.PARAMS) *model.Workspace {
	db := db.DbManager()
	search_val := params.Workspace
	models := FindWorkspaceDB(db, "Name", search_val)
	if models == nil {
		// common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
		return nil
	}

	return models
}

func CreateWorkspace(c echo.Context) (err error) {
	db := db.DbManager()
	// models := new(model.Workspace)

	params := echo.Map{}
	if err := c.Bind(&params); err != nil {
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return err
	}

	selectClusterArr := common.InterfaceToArray(params["selectCluster"])
	var str string
	for i := range selectClusterArr {
		str += selectClusterArr[i] + ","
	}
	str = strings.TrimRight(str, ",")
	params["selectCluster"] = str
	var models model.Workspace
	common.Transcode(params, &models)
	// if err = c.Bind(models); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// common.ErrorMsg(c, http.StatusBadRequest, err)

	// return nil
	// }
	if err = c.Validate(models); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// if err != nil {
	// 	panic(err)
	// }

	if err := db.Create(&models).Error; err != nil {
		common.ErrorMsg(c, http.StatusExpectationFailed, err)
		return nil
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
