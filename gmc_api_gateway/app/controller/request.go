package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"encoding/json"
	"gmc_api_gateway/app/common"
	db "gmc_api_gateway/app/database"
	"gmc_api_gateway/app/model"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRequestDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)

	return cdb
}

func CreateRequest(c echo.Context) (err error) {
	cdb := GetRequestDB("request")
	cdb2 := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	models := new(model.Request)
	// clustermodels := new(model.Cluster)
	validate := validator.New()

	fmt.Println("1")
	fmt.Printf("test: %s",c.Get("date"))
	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
	}
	fmt.Println("2")
	fmt.Println(models)
	fmt.Println(models.ClusterName)
	fmt.Println("3")
	// var result2 map[string]interface{}
	test,err:= cdb2.Find(ctx, bson.M{"clusterName": models.ClusterName})
	// test3 := cdb2.FindOne(ctx, bson.D{{Key:"clusterName", Value:models.ClusterName}}).Decode(&cdb2)
	// if err := cdb2.FindOne(ctx, bson.M{"clusterName": models.ClusterName}).Decode(&cdb2); err != nil {
	// 	fmt.Printf("testeste: %s" ,err )
	// 	// common.ErrorMsg(c, http.StatusInternalServerError, err)
	// 	return nil
	// }
	fmt.Println(test)
	var test2 []bson.M
	
	if err = test.All(ctx, &test2); err != nil{
		log.Fatal(err)
	}
	fmt.Println("4")
	fmt.Println(test2)
	fmt.Println("5")
	fmt.Println(reflect.ValueOf(test2))
	test3 := fmt.Sprint(test2)
	fmt.Println(test3)
	fmt.Println(strings.Split(test3, "clusterEndpoint"))
	fmt.Println("6")
	fmt.Println(StringToInterface(test3))
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

	newRequest := model.Request{
		Id : models.Id,
		Status : models.Status,
		Message : models.Message,
		Name : models.Name,
		Reason : models.Reason,
		Type : models.Type,
		Date : models.Date,
		// Cluster: cdb2.FindOne(ctx, bson.M{"clusterName": models.ClusterName}),		
	}

	// result, err := cdb.InsertOne(ctx, models)
	result, err := cdb.InsertOne(ctx, newRequest)
	if err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, result)
}

func ListRequest(c echo.Context) (err error) {
	var showsRequest []bson.M
	cdb := GetRequestDB("request")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "cluster"}, {"foreignField", "_id"}, {"as", "cluster"}}}}
		lookupWorkspace := bson.D{{"$lookup", bson.D{{"from", "workspace"}, {"localField", "workspace"}, {"foreignField", "_id"}, {"as", "workspace"}}}}
		lookupProject := bson.D{{"$lookup", bson.D{{"from", "project"}, {"localField", "project"}, {"foreignField", "_id"}, {"as", "project"}}}}

		showProjectCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, lookupWorkspace, lookupProject})

		if err = showProjectCursor.All(ctx, &showsRequest); err != nil {
			panic(err)
		}
		// fmt.Println(showsRequest)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, showsRequest)
}

func FindRequest(c echo.Context) (err error) {
	var showsRequest []bson.M
	cdb := GetRequestDB("request")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("requestId")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "cluster"}, {"foreignField", "_id"}, {"as", "cluster"}}}}
		lookupWorkspace := bson.D{{"$lookup", bson.D{{"from", "workspace"}, {"localField", "workspace"}, {"foreignField", "_id"}, {"as", "workspace"}}}}
		lookupProject := bson.D{{"$lookup", bson.D{{"from", "project"}, {"localField", "project"}, {"foreignField", "_id"}, {"as", "project"}}}}
		matchCluster := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "requestId", Value: search_val},
			}},
		}

		showLoadedCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupCluster, lookupWorkspace, lookupProject, matchCluster})

		if err = showLoadedCursor.All(ctx, &showsRequest); err != nil {
			panic(err)
		}
		fmt.Println(showsRequest)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	if showsRequest == nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Request not found."))
		return
	} else {
		return c.JSON(http.StatusOK, showsRequest)
	}
}

func DeleteRequest(c echo.Context) (err error) {
	cdb := GetRequestDB("request")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("requestId")

	result, err := cdb.DeleteOne(ctx, bson.M{"requestId": search_val})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
		return
	}
	if result.DeletedCount == 0 {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Request not found."))
		return
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"status": http.StatusOK,
			"data":   search_val + " Deleted Complete",
		})
	}
}

func UpdateRequest(c echo.Context) (err error) {
	cdb := GetRequestDB("request")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("requestId")

	models := new(model.RequestUpdate)
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

	update := bson.M{"status": models.Status, "name": models.Name, "reason": models.Reason, "message": models.Message, "cluster": models.Cluster, "date": models.Date}

	result, err := cdb.UpdateOne(ctx, bson.M{"requestId": search_val}, bson.M{"$set": update})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to update."))
		return
	}
	
	if result.MatchedCount == 1 {
		if err := cdb.FindOne(ctx, bson.M{"requestId": search_val}).Decode(&cdb); err != nil {
			common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to match request."))
			return nil
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   search_val + " Updated Complete",
	})	
}
func StringToInterface(i string) interface{} {
	var x interface{}
	if err := json.Unmarshal([]byte(i), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	return x
}