package controller

import (
	"context"
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
	models2 := new(model.ProjectClusters)
	validate := validator.New()

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
	}

	memberObjectId,err:= cdb2.Find(ctx, bson.M{"memberName": models.MemberName})
	workspaceObjectId,err:= cdb3.Find(ctx, bson.M{"workspaceName": models.WorkspaceName})
	clusterObjectId,err:= cdb4.Find(ctx, bson.M{"clusterName": models2.ClusterName})

	var memberObjectId2 []bson.D
	var workspaceObjectId2 []bson.D
	var clusterObjectId2 []bson.A


	if err = memberObjectId.All(ctx, &memberObjectId2); err != nil{
		log.Fatal(err)
	}
	if err = workspaceObjectId.All(ctx, &workspaceObjectId2); err != nil{
		log.Fatal(err)
	}
	if err = clusterObjectId.All(ctx, &clusterObjectId2); err != nil{
		log.Fatal(err)
	}
	// fmt.Println(memberObjectId2[0][0].Value)
	// fmt.Println(workspaceObjectId2[0][0].Value)
	// fmt.Println(clusterObjectId2)
	// fmt.Println("ㅅㄴㄷㅅㄴㄴㅅㄴㅅ")
	// fmt.Println(models.MemberName)
	// fmt.Println(models.Selectcluster)
	// fmt.Println(models2.ClusterName)
	// fmt.Println("ㅅㄴㄷㅅㄴㄴㅅㄴㅅ")

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

	// tstst := model.ProjectClusters{
	// 	Cluster: clusterObjectId2[0].(primitive.ObjectID),
	// }
	// fmt.Println(tstst)

	newRequest := model.NewProject{
		Name : models.Name,
		Description : models.Description,
		Type : models.Type,
		Owner : memberObjectId2[0][0].Value.(primitive.ObjectID),
		Creator : memberObjectId2[0][0].Value.(primitive.ObjectID),
		Created_at : models.Created_at,
		Workspace: workspaceObjectId2[0][0].Value.(primitive.ObjectID),
		// Selectcluster : tstst,
		// .clusterObjectId2[0][0].Value.(primitive.ObjectID)
	}

	// result, err := cdb.InsertOne(ctx, models)
	result, err := cdb.InsertOne(ctx, newRequest)
	if err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, result)
}

func ListProject(c echo.Context) (err error) {
	var showsProject []bson.M
	cdb := GetProjectDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "selectCluster.cluster"}, {"foreignField", "_id"}, {"as", "selectCluster"}}}}
		lookupWorkspace := bson.D{{"$lookup", bson.D{{"from", "workspace"}, {"localField", "workspace"}, {"foreignField", "_id"}, {"as", "workspace"}}}}

		showProjectCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, lookupWorkspace})

		if err = showProjectCursor.All(ctx, &showsProject); err != nil {
			panic(err)
		}
		// fmt.Println(showsProject)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, showsProject)
}

func FindProject(c echo.Context) (err error) {
	var showsProject []bson.M
	cdb := GetWorkspaceDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("projectName")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "selectCluster.cluster"}, {"foreignField", "_id"}, {"as", "selectCluster"}}}}
		lookupWorkspace := bson.D{{"$lookup", bson.D{{"from", "workspace"}, {"localField", "workspace"}, {"foreignField", "_id"}, {"as", "workspace"}}}}
		matchCluster := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "projectName", Value: search_val},
			}},
		}

		showLoadedCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, lookupWorkspace, matchCluster})

		if err = showLoadedCursor.All(ctx, &showsProject); err != nil {
			panic(err)
		}
		fmt.Println(showsProject)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	if showsProject == nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Project not found."))
		return
	} else {
		return c.JSON(http.StatusOK, showsProject)
	}
}

func DeleteProject(c echo.Context) (err error) {
	cdb := GetProjectDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("projectName")

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

// func UpdateProject(c echo.Context) (err error) {
// 	cdb := GetProjectDB("project")
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

// 	models := new(model.Project)

// 	if err = c.Bind(models); err != nil {
// 		common.ErrorMsg(c, http.StatusBadRequest, err)
// 		return nil
// 	}

// 	if err = c.Validate(models); err != nil {
// 		common.ErrorMsg(c, http.StatusUnprocessableEntity, err)
// 		return nil
// 	}

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	search_val := c.Param("id")

// 	result, err := cdb.UpdateOne(ctx, bson.M{"projectId": search_val})
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
// 		return
// 	}
// 	if result.DeletedCount == 0 {
// 		common.ErrorMsg(c, http.StatusNotFound, errors.New("Project not found."))
// 		return
// 	} else {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"status": http.StatusOK,
// 			"data":   search_val + " Deleted",
// 		})
// 	}
// }
