package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"gmc_api_gateway/app/common"
	db "gmc_api_gateway/app/database"
	"gmc_api_gateway/app/model"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Namespace struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name   string `json:"name"`
		Labels struct {
			IstioCheck string `json:"istio-injection"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
	} `json:"spec"`
	Status struct {
	} `json:"status"`
}

func GetProjectDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)
	return cdb
}

func CreateProject(c echo.Context) (err error) {
	cdb := GetProjectDB("project")
	cdb2 := GetProjectDB("member")
	cdb3 := GetProjectDB("workspace")
	cdb4 := GetProjectDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	models := new(model.Project)
	validate := validator.New()

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
	}

	memberObjectId, err := cdb2.Find(ctx, bson.M{"memberName": models.MemberName})
	workspaceObjectId, err := cdb3.Find(ctx, bson.M{"workspaceName": models.WorkspaceName})

	var clusterObjectId2 []bson.D
	var clusterObjectId3 *mongo.Cursor
	var memberObjectId2 []bson.D
	var workspaceObjectId2 []bson.D
	var slice []primitive.ObjectID

	for i := 0; i < len(models.ClusterName); i++ {
		clusterObjectId3, _ = cdb4.Find(ctx, bson.M{"clusterName": models.ClusterName[i]})
		clusterObjectId3.All(ctx, &clusterObjectId2)
		slice = append(slice, clusterObjectId2[0][0].Value.(primitive.ObjectID))
	}

	if err = memberObjectId.All(ctx, &memberObjectId2); err != nil {
		log.Fatal(err)
	}
	if err = workspaceObjectId.All(ctx, &workspaceObjectId2); err != nil {
		log.Fatal(err)
	}

	if err = validate.Struct(models); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	newProject := model.NewProject{
		Name:          models.Name,
		Description:   models.Description,
		Type:          models.Type,
		Owner:         memberObjectId2[0][0].Value.(primitive.ObjectID),
		Creator:       memberObjectId2[0][0].Value.(primitive.ObjectID),
		Created_at:    models.Created_at,
		Workspace:     workspaceObjectId2[0][0].Value.(primitive.ObjectID),
		Selectcluster: slice,
		IstioCheck:    models.IstioCheck,
	}

	result, err := cdb.InsertOne(ctx, newProject)
	if err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}
	for _, cluster := range models.ClusterName {
		clusterInfo := FindClusterDB(cluster)
		namespace := Namespace{}
		namespace.APIVersion = "v1"
		namespace.Kind = "Namespace"
		namespace.Metadata.Name = models.Name
		namespace.Metadata.Labels.IstioCheck = models.IstioCheck
		url := "https://" + clusterInfo.Endpoint + ":6443/api/v1/namespaces/"
		Token := clusterInfo.Token
		// fmt.Println("clusterInfo.Endpoint: ", clusterInfo.Endpoint)
		// fmt.Println("clusterInfo.Token: ", clusterInfo.Token)
		data, err := json.Marshal(namespace)
		fmt.Printf("// %s", data)
		if err != nil {
			common.ErrorMsg(c, http.StatusBadRequest, err)
			return err
		}
		var jsonStr = []byte(fmt.Sprint(string(data)))
		code := RequsetKube(url, "POST", jsonStr, Token)
		fmt.Println("code", code)
		switch code {
		case 200:
		case 201:
		case 202:
		// case 409:
		// cdb.DeleteOne(ctx, bson.M{"_id": result.InsertedID})
		// fmt.Println("result : ", result.InsertedID)
		// common.ErrorMsg(c, http.StatusBadRequest, err)
		// return err
		default:
			cdb.DeleteOne(ctx, bson.M{"_id": result.InsertedID})
			common.ErrorMsg(c, http.StatusBadRequest, err)
			return err
		}
	}
	return c.JSON(http.StatusOK, result)
}

func ListUserProject(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		User:      c.QueryParam("user"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	var showsProject []bson.M
	// var userProject model.NewProject
	var userProjects []model.USERPROJECT

	if params.User == "" {
		cdb := GetProjectDB("project")

		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

		findOptions := options.Find()

		cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}
		if err = cur.All(ctx, &showsProject); err != nil {
			panic(err)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.TODO())

	} else {
		userObj := FindMemberDB(params)
		showsProject = GetDBList(params, "project", userObj.ObjectId, "projectOwner")
	}
	for _, project := range showsProject {
		params.Project = common.InterfaceToString(project["projectName"])
		temp_project := GetDBProject(params)
		var UserProject model.USERPROJECT
		UserProject.DBProject = temp_project
		userProjects = append(userProjects, UserProject)
	}
	// var userProjects []model.USERPROJECT

	// findOptions := options.Find()

	// cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for cur.Next(context.TODO()) {
	// 	lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "selectCluster"}, {"foreignField", "_id"}, {"as", "selectCluster"}}}}
	// 	lookupWorkspace := bson.D{{"$lookup", bson.D{{"from", "workspace"}, {"localField", "workspace"}, {"foreignField", "_id"}, {"as", "workspace"}}}}

	// 	showProjectCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, lookupWorkspace})

	// 	if err = showProjectCursor.All(ctx, &showsProject); err != nil {
	// 		panic(err)
	// 	}
	// }
	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// cur.Close(context.TODO())

	return c.JSON(http.StatusOK, userProjects)
}

func ListSystemProject(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("cluster"),
		User:      c.QueryParam("user"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	Projects := ListDB("project")
	params.Project = ""
	var projects []model.SYSTEMPROJECT
	getData := GetModelList(params)
	fmt.Println("getData : ", getData)
	for k, _ := range getData {
		project := model.SYSTEMPROJECT{
			Name:        common.InterfaceToString(common.FindData(getData[k], "metadata", "name")),
			Status:      common.InterfaceToString(common.FindData(getData[k], "status", "phase")),
			ClusterName: common.InterfaceToString(common.FindData(getData[k], "clusterName", "")),
			CreateAt:    common.InterfaceToTime(common.FindData(getData[k], "metadata", "creationTimestamp")),
		}
		projects = append(projects, project)
	}
	projectList := difference(Projects, projects)

	return c.JSON(http.StatusOK, echo.Map{
		"data": projectList,
	})

}
func difference(slice1 []primitive.M, slice2 []model.SYSTEMPROJECT) []model.SYSTEMPROJECT {
	var diff []model.SYSTEMPROJECT

	for _, s1 := range slice2 {
		found := false
		for _, s2 := range slice1 {
			if s1.Name == common.InterfaceToString(s2["projectName"]) {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, s1)
		}
	}
	return diff
}

func GetUserProject(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("cluster"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	params.Project = params.Name
	project := GetDBProject(params)
	if project.Name == "" {
		msg := common.ErrorMsg2(http.StatusNotFound, common.ErrNotFound)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": msg,
		})
	}
	clusters := project.Selectcluster
	resourceCnt, resourceUsage, eventList := GetUserProjectResource(params, clusters)

	var detailList []model.PROJECT_DETAIL
	for _, cluster := range clusters {
		params.Cluster = cluster.Name
		// params.Project = ""
		getData, err := common.DataRequest(params)
		if err != nil || common.InterfaceToString(common.FindData(getData, "status", "")) == "Failure" {
			msg := common.ErrorMsg2(http.StatusNotFound, common.ErrNotFound)
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": msg,
			})
		}
		// tempMetric := []string{"namespace_cpu", "namespace_memory", "namespace_pod_count"}
		// tempresult := NowMonit("namespace", params.Cluster, params.Name, tempMetric)
		fmt.Println("getData : ", getData)
		projectDetail := model.PROJECT_DETAIL{
			Status:        common.InterfaceToString(common.FindData(getData, "status", "phase")),
			ClusterName:   cluster.Name,
			Resource:      resourceCnt,
			Label:         common.FindData(getData, "metadata", "labels"),
			Annotation:    common.FindData(getData, "metadata", "annotations"),
			ResourceUsage: resourceUsage,
			CreateAt:      common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		}
		detailList = append(detailList, projectDetail)

	}
	userProject := model.USERPROJECT{
		DBProject: project,
		Events:    eventList,
		Detail:    detailList,
	}

	// var showsProject []bson.M
	// cdb := GetWorkspaceDB("project")
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// search_val := c.Param("projectName")

	// findOptions := options.Find()

	// cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for cur.Next(context.TODO()) {
	// 	lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "selectCluster"}, {"foreignField", "_id"}, {"as", "selectCluster"}}}}
	// 	lookupWorkspace := bson.D{{"$lookup", bson.D{{"from", "workspace"}, {"localField", "workspace"}, {"foreignField", "_id"}, {"as", "workspace"}}}}
	// 	matchCluster := bson.D{
	// 		{Key: "$match", Value: bson.D{
	// 			{Key: "projectName", Value: search_val},
	// 		}},
	// 	}

	// 	showLoadedCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, lookupWorkspace, matchCluster})

	// 	if err = showLoadedCursor.All(ctx, &showsProject); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(showsProject)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// cur.Close(context.TODO())

	// if showsProject == nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, errors.New("Project not found."))
	// 	return
	// } else {
	// 	return c.JSON(http.StatusOK, showsProject)
	// }
	return c.JSON(http.StatusOK, userProject)

}
func GetSystemProject(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("cluster"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	if params.Cluster == "" {
		msg := common.ErrorMsg2(http.StatusNotFound, common.ErrClusterNotFound)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": msg,
		})
	}
	params.Project = params.Name
	getData, err := common.DataRequest(params)
	if err != nil || common.InterfaceToString(common.FindData(getData, "status", "")) == "Failure" {
		msg := common.ErrorMsg2(http.StatusNotFound, common.ErrNotFound)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": msg,
		})
	}
	fmt.Println("getData : ", getData)
	// for k, _ := range getData {
	project := model.SYSTEMPROJECT{
		Name:        common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Status:      common.InterfaceToString(common.FindData(getData, "status", "phase")),
		ClusterName: common.InterfaceToString(common.FindData(getData, "clusterName", "")),
		CreateAt:    common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
	}
	project.Detail.ClusterName = params.Cluster
	project.Detail.CreateAt = common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp"))
	project.Detail.Label = common.FindData(getData, "metadata", "labels")
	project.Detail.Annotation = common.FindData(getData, "metadata", "Annotations")
	var Clusters []model.Cluster
	var tsCluster model.Cluster
	cluster := FindClusterDB(params.Cluster)
	common.Transcode(cluster, &tsCluster)
	Clusters = append(Clusters, tsCluster)
	resourceCnt, resourceUsage, eventList := GetUserProjectResource(params, Clusters)
	project.Events = eventList
	project.Detail.ResourceUsage = resourceUsage
	project.Detail.Resource = resourceCnt

	return c.JSON(http.StatusOK, project)
}
func DeleteProject(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("cluster"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	cdb := GetProjectDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("name")
	params.Project = c.Param("name")
	project := GetDBProject(params)
	fmt.Println("project : ", project)
	for _, cluster := range project.Selectcluster {
		// fmt.Printf("########clusterName : %s", slice[i])
		// clusters := GetClusterDB(common.InterfaceToString(slice[i]))

		url := "https://" + cluster.Endpoint + ":6443/api/v1/namespaces/" + params.Name
		Token := cluster.Token

		if err != nil {
			common.ErrorMsg(c, http.StatusBadRequest, err)
			return err
		}

		code := RequsetKube(url, "DELETE", nil, Token)

		switch code {
		case 200:
		case 202:
		default:
			common.ErrorMsg(c, http.StatusBadRequest, err)
			return err
		}
	}
	result, err := cdb.DeleteOne(ctx, bson.M{"projectName": search_val})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
		return
	}
	if result.DeletedCount == 0 {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Project not found."))
		return
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"status": http.StatusOK,
			"data":   search_val + " Project Deleted Complete",
		})
	}
}

func UpdateProject(c echo.Context) (err error) {
	cdb := GetRequestDB("project")
	cdb2 := GetProjectDB("cluster")
	cdb3 := GetProjectDB("member")
	cdb4 := GetProjectDB("workspace")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("projectName")

	models := new(model.RequestProject)
	validate := validator.New()

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
	}
	memberObjectId, err := cdb3.Find(ctx, bson.M{"memberName": models.MemberName})
	workspaceObjectId, err := cdb4.Find(ctx, bson.M{"workspaceName": models.WorkspaceName})

	var clusterObjectId2 []bson.D
	var clusterObjectId3 *mongo.Cursor
	var memberObjectId2 []bson.D
	var workspaceObjectId2 []bson.D
	var slice []primitive.ObjectID

	for i := 0; i < len(models.ClusterName); i++ {
		clusterObjectId3, _ = cdb2.Find(ctx, bson.M{"clusterName": models.ClusterName[i]})
		clusterObjectId3.All(ctx, &clusterObjectId2)
		slice = append(slice, clusterObjectId2[0][0].Value.(primitive.ObjectID))
	}

	if err = memberObjectId.All(ctx, &memberObjectId2); err != nil {
		log.Fatal(err)
	}
	if err = workspaceObjectId.All(ctx, &workspaceObjectId2); err != nil {
		log.Fatal(err)
	}

	if err = validate.Struct(models); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	var update primitive.M
	// switch models.조건{
	// case nil :
	// update = bson.M{"workspace" : workspaceObjectId2[0][0].Value.(primitive.ObjectID),"projectOwner": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectCreator": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectDescription": models.Description, "selectCluster":slice}
	// default :
	update = bson.M{"workspace": workspaceObjectId2[0][0].Value.(primitive.ObjectID), "projectOwner": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectCreator": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectDescription": models.Description, "selectCluster": slice}
	// }

	result, err := cdb.UpdateOne(ctx, bson.M{"projectName": search_val}, bson.M{"$set": update})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to update."))
		return
	}

	if result.MatchedCount == 1 {
		if err := cdb.FindOne(ctx, bson.M{"projectName": search_val}).Decode(&cdb); err != nil {
			common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to match Project."))
			return nil
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   search_val + " Updated Complete",
	})
}

func GetDBProject(params model.PARAMS) model.DBProject {
	var project model.NewProject
	var showsProject model.DBProject
	var results bson.M
	var workspace model.Workspace
	var user model.Member
	var clusterList []model.Cluster
	cdb := GetClusterDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := params.Project

	// query := &bson.M{
	// 	"projectOwner": params.User,
	// 	"projectName":  search_val,
	// }

	if err := cdb.FindOne(ctx, bson.M{"projectName": search_val}).Decode(&project); err != nil {

	}
	if err := cdb.FindOne(ctx, bson.M{"projectName": search_val}).Decode(&results); err != nil {

	}
	if err := cdb.FindOne(ctx, bson.M{"projectName": search_val}).Decode(&showsProject); err != nil {

	}
	user_objectId := project.Owner
	userList := GetClusterDB("member")
	users, _ := context.WithTimeout(context.Background(), time.Second*10)
	if err := userList.FindOne(users, bson.M{"_id": user_objectId}).Decode(&user); err != nil {
	}
	workspace_objectId := project.Workspace
	workspaceList := GetClusterDB("workspace")
	workspaces, _ := context.WithTimeout(context.Background(), time.Second*10)
	if err := workspaceList.FindOne(workspaces, bson.M{"_id": workspace_objectId}).Decode(&workspace); err != nil {
	}
	tempList := GetClusterDB("cluster")

	clusters, _ := context.WithTimeout(context.Background(), time.Second*10)

	cluster_objectId := project.Selectcluster
	for i := range cluster_objectId {
		var cluster model.Cluster
		if err := tempList.FindOne(clusters, bson.M{"_id": cluster_objectId[i]}).Decode(&cluster); err != nil {
		}
		clusterList = append(clusterList, cluster)
	}
	showsProject.Workspace = workspace
	showsProject.Selectcluster = clusterList
	showsProject.MemberName = user.Id
	return showsProject
}
