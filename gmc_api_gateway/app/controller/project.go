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

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
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
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

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
		log.Fatal(err)
	}

	result, err := cdb.InsertOne(ctx, models)
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
	// return c.JSON(http.StatusOK, showsProject)
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
