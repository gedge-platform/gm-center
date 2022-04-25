package controller

import (
	"context"
	"errors"
	"fmt"
	"gmc_api_gateway/app/common"
	db "gmc_api_gateway/app/database"
	"gmc_api_gateway/app/model"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCredentialDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)

	return cdb
}

func ListCredential(c echo.Context) (err error) {
	var results []model.Credential
	cdb := GetCredentialDB("credentials")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Credential

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

func FindCredential(c echo.Context) (err error) {
	var credential model.Credential
	// var keyvalueinfolist []model.KeyValueInfoList
	cdb := GetCredentialDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("name")

	if err := cdb.FindOne(ctx, bson.M{"name": search_val}).Decode(&credential); err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Credential not found."))
		return nil
	} else {
		fmt.Println()
		return c.JSON(http.StatusOK, &credential)
	}
}

func InsertCredential(c echo.Context) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	cdb := GetClusterDB("credentials")

	param := model.CPARAM{
		Name:         c.QueryParam("name"),
		Type:         "OPENSTACK",
		Domain:       c.QueryParam("domain"),
		Region:       c.QueryParam("region"),
		Url:          c.QueryParam("url"),
		Tenant:       c.QueryParam("tenant"),
		Access_id:    c.QueryParam("access_id"),
		Access_token: c.QueryParam("access_token"),
		Project:      c.QueryParam("project")}
	// Method:       c.Request().Method}
	// c.Bind(param)
	// c.Validate(param)
	cdb.InsertOne(ctx, param)
	return c.JSON(http.StatusCreated, echo.Map{"data": param})
}

func UpdateCredential(c echo.Context) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	cdb := GetClusterDB("credentials")
	search_val := c.Param("name")

	param := model.CPARAM{
		Name:         search_val,
		Type:         "OPENSTACK",
		Domain:       c.QueryParam("domain"),
		Region:       c.QueryParam("region"),
		Url:          c.QueryParam("url"),
		Tenant:       c.QueryParam("tenant"),
		Access_id:    c.QueryParam("access_id"),
		Access_token: c.QueryParam("access_token"),
		Project:      c.QueryParam("project"),
	}

	cdb.UpdateOne(ctx, bson.M{"name": search_val}, bson.M{"$set": param})
	return c.JSON(http.StatusOK, search_val+"Updated Complete")

}

func DeleteCredential(c echo.Context) (err error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	cdb := GetClusterDB("credentials")
	// var credential model.Credential

	search_val := c.Param("name")
	result, err := cdb.DeleteOne(ctx, bson.M{"name": search_val})

	if result.DeletedCount == 0 {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Member not found."))
		return
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"status": http.StatusOK,
			"data":   search_val + " Member Deleted Complete",
		})
	}
}
