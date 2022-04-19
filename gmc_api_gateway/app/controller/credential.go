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
