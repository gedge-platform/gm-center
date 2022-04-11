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

func GetClusterDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)

	return cdb
}

func CreateCluster(c echo.Context) (err error) {
	cdb := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	models := new(model.Cluster)
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

func ListCluster(c echo.Context) (err error) {
	var results []model.Cluster
	cdb := GetClusterDB("cluster")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Cluster
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

func FindCluster(c echo.Context) (err error) {
	var cluster model.Cluster
	cdb := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("clusterName")

	if err := cdb.FindOne(ctx, bson.M{"clusterName": search_val}).Decode(&cluster); err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Cluster not found."))
		return nil
	} else {
		return c.JSON(http.StatusOK, &cluster)
	}
}

func DeleteCluster(c echo.Context) (err error) {
	cdb := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("clusterName")

	result, err := cdb.DeleteOne(ctx, bson.M{"clusterName": search_val})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
		return
	}
	if result.DeletedCount == 0 {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Cluster not found."))
		return
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"status": http.StatusOK,
			"data":   search_val + " Cluster Deleted Complete",
		})
	}
}

func UpdateCluster(c echo.Context) (err error) {
	cdb := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("clusterName")

	models := new(model.RequestCluster)
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

	var update primitive.M
	// switch models.조건{
	// case nil :
	// update = bson.M{"clusterEndpoint": models.Endpoint, "clusterType":models.Type, "token": models.Token}
	// default :
	update = bson.M{"clusterEndpoint": models.Endpoint, "clusterType": models.Type, "token": models.Token}
	// }

	result, err := cdb.UpdateOne(ctx, bson.M{"clusterName": search_val}, bson.M{"$set": update})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to update."))
		return
	}

	if result.MatchedCount == 1 {
		if err := cdb.FindOne(ctx, bson.M{"clusterName": search_val}).Decode(&cdb); err != nil {
			common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to match Cluster."))
			return nil
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   search_val + " Updated Complete",
	})
}
