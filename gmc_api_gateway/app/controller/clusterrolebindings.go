package controller

import (
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetClusterrolebinding godoc
// @Summary Show detail clusterrolebinding
// @Description get clusterrolebinding Details
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Param name path string true "name of the Clusterrolebinding"
// @Param cluster query string true "name of the Cluster"
// @Param workspace query string true "name of the Workspace"
// @Success 200 {object} model.CLUSTERROLEBINDING
// @Router /clusterrolebindings/{name} [get]
// @Tags Kubernetes
func GetClusterrolebinding(c echo.Context) error {
	params := model.PARAMS{
		Kind:      "clusterrolebindings",
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
	clusterrolebinding := model.CLUSTERROLEBINDING{
		Name:        common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Labels:      common.FindData(getData, "metadata", "labels"),
		Subjects:    common.FindData(getData, "subjects", ""),
		RoleRef:     common.FindData(getData, "roleRef", ""),
		Annotations: common.FindData(getData, "data", "annotations"),
		CreateAt:    common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		Cluster:     params.Cluster,
		Workspace:   params.Workspace,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": clusterrolebinding,
	})
}

// GetClusterrolebinding godoc
// @Summary Show List clusterrolebinding
// @Description get clusterrolebinding List
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Param clusterName query string false "name of the Cluster"
// @Param workspaceName query string false "name of the Workspace"
// @Success 200 {object} model.CLUSTERROLEBINDING
// @Router /clusterrolebindings [get]
// @Tags Kubernetes
func GetAllClusterrolebindings(c echo.Context) error {
	var clusterrolebindings []model.CLUSTERROLEBINDING
	// fmt.Printf("## clusterrolebings", clusterrolebindings)
	params := model.PARAMS{
		Kind:      "clusterrolebindings",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		User:      c.QueryParam("user"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	// getData, err := common.DataRequest(params)
	// if err != nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, err)
	// 	return nil
	// }

	data := GetModelList(params)

	for i, _ := range data {
		clusterrolebinding := model.CLUSTERROLEBINDING{
			Name:        common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Subjects:    common.FindData(data[i], "subjects", ""),
			RoleRef:     common.FindData(data[i], "roleRef", ""),
			Labels:      common.FindData(data[i], "metadata", "labels"),
			Annotations: common.FindData(data[i], "metadata", "annotations"),
			CreateAt:    common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
			Cluster:     common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			Workspace:   common.InterfaceToString(common.FindData(data[i], "workspaceName", "")),
			UserName:    common.InterfaceToString(common.FindData(data[i], "userName", "")),
		}
		if params.User != "" {
			if params.User == clusterrolebinding.UserName {
				clusterrolebindings = append(clusterrolebindings, clusterrolebinding)
			}
		} else {
			clusterrolebindings = append(clusterrolebindings, clusterrolebinding)
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": clusterrolebindings,
	})
}
func CreateClusterRolebinding(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:    "clusterrolebindings",
		Cluster: c.QueryParam("cluster"),
		Project: c.QueryParam("project"),
		Method:  c.Request().Method,
		Body:    responseBody(c.Request().Body),
	}

	postData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"info": common.StringToInterface(postData),
	})
}

func DeleteClusterRolebinding(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:    "clusterrolebindings",
		Name:    c.Param("name"),
		Cluster: c.QueryParam("cluster"),
		Project: c.QueryParam("project"),
		Method:  c.Request().Method,
		Body:    responseBody(c.Request().Body),
	}

	postData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{
		"info": common.StringToInterface(postData),
	})
}
