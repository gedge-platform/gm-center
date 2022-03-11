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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetWorkspaceDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)

	return cdb
}

func CreateWorkspace(c echo.Context) (err error) {
	cdb := GetWorkspaceDB("workspace")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	models := new(model.Workspace)
	validate := validator.New()

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
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

	result, err := cdb.InsertOne(ctx, models)
	if err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, result)
}

func ListWorkspace(c echo.Context) (err error) {
	var showsWorkspace []bson.M
	cdb := GetWorkspaceDB("workspace")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "selectCluster.cluster"}, {"foreignField", "_id"}, {"as", "selectCluster"}}}}
		// unwindCluster := bson.D{{"$unwind", bson.D{{"path", "$selectCluster2"}, {"preserveNullAndEmptyArrays", false}}}}
		// matchCluster := bson.D{
		// 	{Key: "$match", Value: bson.D{
		// 		{Key: "cluster", Value: "$ID"},
		// 	}},
		// }
		fmt.Println("ttt : ", mongo.Pipeline{lookupCluster})
		showWorkspaceCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster})
		// showWorkspaceCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, unwindCluster})

		if err = showWorkspaceCursor.All(ctx, &showsWorkspace); err != nil {
			panic(err)
		}
		// fmt.Println(showsWorkspace)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, showsWorkspace)
}

func FindWorkspace(c echo.Context) (err error) {
	var showsWorkspace []bson.M
	cdb := GetWorkspaceDB("workspace")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("workspaceName")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "selectCluster.cluster"}, {"foreignField", "_id"}, {"as", "selectCluster"}}}}
		matchCluster := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "workspaceName", Value: search_val},
			}},
		}

		showLoadedCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, matchCluster})

		if err = showLoadedCursor.All(ctx, &showsWorkspace); err != nil {
			panic(err)
		}
		fmt.Println(showsWorkspace)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	if showsWorkspace == nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Workspace not found."))
		return
	} else {
		return c.JSON(http.StatusOK, showsWorkspace)
	}
	// return c.JSON(http.StatusOK, showsWorkspace)
}

func DeleteWorkspace(c echo.Context) (err error) {
	cdb := GetWorkspaceDB("workspace")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("workspaceName")

	result, err := cdb.DeleteOne(ctx, bson.M{"workspaceName": search_val})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
		return
	}
	if result.DeletedCount == 0 {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Workspace not found."))
		return
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"status": http.StatusOK,
			"data":   search_val + " Workspace Deleted Complete",
		})
	}
}

// func UpdateWorkspace(c echo.Context) (err error) {
// 	cdb := GetWorkspaceDB("workspace")
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

// 	models := new(model.Workspace)

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

// 	result, err := cdb.UpdateOne(ctx, bson.M{"workspaceId": search_val})
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
// 		return
// 	}
// 	if result.DeletedCount == 0 {
// 		common.ErrorMsg(c, http.StatusNotFound, errors.New("Workspace not found."))
// 		return
// 	} else {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"status": http.StatusOK,
// 			"data":   search_val + " Deleted",
// 		})
// 	}
// }
