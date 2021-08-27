package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/db"
	"gmc_database_api_server/app/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
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

	if models.SelectCluster != "" {
		models2.SelectCluster = models.SelectCluster
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
func Get_Project(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:    "namespaces",
		Name:    c.Param("name"),
		Cluster: c.QueryParam("cluster"),
		Project: c.QueryParam("project"),
		Method:  c.Request().Method,
		Body:    c.Request().Body,
	}
	params.Workspace = c.Param("name")
	getData, err := common.GetModel(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}
	if common.FindData(getData, "status", "") == "Failure" {
		fmt.Printf("error code : %s\n", getData)
		err_data := common.InterfaceToString(getData)
		// errReturn := KubernetesNS.Array()
		errJson := make(map[string]string)
		err_ns := json.Unmarshal([]byte(err_data), &errJson)
		if err_ns != nil {
			fmt.Printf("err_ns : %s\n", err_ns)
		}
		// 	common.ErrorMsg(c, http.StatusNotFound, err)
		// return c.JSON(http.StatusNotFound, errJson)
		return c.JSON(http.StatusNotFound, errJson)
	}
	projectModel := GetProject2(c)
	getData0 := common.FindData(getData, "", "") // 빈칸으로 둘 시, 전체 조회
	var Namespace model.Namespace
	common.Transcode(getData0, &Namespace)
	projectModel.Name = common.InterfaceToString(common.FindData(getData, "metadata", "name"))
	projectModel.CreateAt = common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp"))
	projectModel.Label = common.FindData(getData, "metadata", "labels")
	projectModel.Annotation = common.FindData(getData, "metadata", "annotations")
	params.Workspace = projectModel.WorkspaceName
	projectModel.Status = common.InterfaceToString(common.FindData(getData, "status", "phase"))
	projectModel.ClusterName = c.QueryParam("cluster")
	projectModel.Events = getCallEvent(params)
	ResourceCnt := model.PROJECT_RESOURCE{
		DeploymentCount: ResourceCnt(params, "deployments"),
		PodCount:        ResourceCnt(params, "pods"),
		ServiceCount:    ResourceCnt(params, "services"),
		CronjobCount:    ResourceCnt(params, "cronjobs"),
		JobCount:        ResourceCnt(params, "jobs"),
		// VolumeCount:     ResourceCnt(params, "deployments"),
	}

	projectModel.Resource = ResourceCnt
	return c.JSON(http.StatusOK, echo.Map{
		"project": projectModel,
	})
}
func Get_Projects(c echo.Context) (err error) {
	var Projects model.Projects
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("cluster"),
		Method:    c.Request().Method,
		Body:      c.Request().Body,
	}
	if c.QueryParam("workspace") == "" {
		params.Workspace = c.QueryParam("cluster")
		params.Project = c.QueryParam("cluster")
		getData, err := common.GetModel(params)
		if err != nil {
			common.ErrorMsg(c, http.StatusNotFound, err)
			return nil
		}
		getData0 := common.FindingArray(common.Finding(getData, "items"))
		for k, _ := range getData0 {
			params.Name = (gjson.Get(getData0[k].String(), "metadata.name")).String()
			ProjectModel := GetProject3(params)
			var Project model.Project
			common.Transcode(ProjectModel, &Project)
			Project.Name = params.Name
			Project.Status = (gjson.Get(getData0[k].String(), "status.phase")).String()
			Project.CreateAt = (gjson.Get(getData0[k].String(), "metadata.creationTimestamp")).Time()
			Project.ClusterName = params.Cluster
			tempMetric := []string{"namespace_cpu", "namespace_memory", "namespace_pod_count"}
			tempresult := NowMonit("namespace", params.Cluster, params.Name, tempMetric)
			Project.ResourceUsage = tempresult
			Projects = append(Projects, Project)
		}

	} else {
		params.Workspace = c.QueryParam("workspace")
		params.Project = c.QueryParam("workspace")
		workspace := GetWorkspace2(params)
		if workspace == nil {
			common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
			return
		}
		selectCluster := workspace.SelectCluster
		slice := strings.Split(selectCluster, ",")
		for i, _ := range slice {
			params.Cluster = slice[i]
			params.Name = ""
			getData, err := common.GetModel(params)
			if err != nil {
				common.ErrorMsg(c, http.StatusNotFound, err)
				return nil
			}
			getData0 := common.FindingArray(common.Finding(getData, "items"))
			for k, _ := range getData0 {
				params.Name = (gjson.Get(getData0[k].String(), "metadata.name")).String()
				ProjectModel := GetProject3(params)
				var Project model.Project
				common.Transcode(ProjectModel, &Project)
				Project.Name = params.Name
				Project.Status = (gjson.Get(getData0[k].String(), "status.phase")).String()
				Project.CreateAt = (gjson.Get(getData0[k].String(), "metadata.creationTimestamp")).Time()
				Project.ClusterName = params.Cluster
				tempMetric := []string{"namespace_cpu", "namespace_memory", "namespace_pod_count"}
				tempresult := NowMonit("namespace", params.Cluster, params.Name, tempMetric)
				Project.ResourceUsage = tempresult
				Projects = append(Projects, Project)
			}
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": Projects,
	})
}
func ResourceCnt(params model.PARAMS, kind string) int {
	params.Kind = kind
	params.Name = ""
	deployments, _ := common.GetModel(params)
	deployment := common.FindDataStr(deployments, "items", "")
	deployment_cnt := common.FindingLen(deployment)
	fmt.Printf("deployment_cnt : %d\n", deployment_cnt)
	return deployment_cnt
}
