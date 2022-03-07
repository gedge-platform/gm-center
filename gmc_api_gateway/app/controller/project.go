package controller

import (
	"context"
	"errors"
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
	var results []model.Project
	cdb := GetProjectDB("project")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Project
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, results)
}

func FindProject(c echo.Context) (err error) {
	var project model.Project
	cdb := GetProjectDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("projectName")

	if err := cdb.FindOne(ctx, bson.M{"projectName": search_val}).Decode(&project); err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	} else {
		return c.JSON(http.StatusOK, &project)
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
			"data":   search_val + " Deleted",
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
