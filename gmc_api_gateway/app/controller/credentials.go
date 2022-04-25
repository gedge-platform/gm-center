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

func GetCredentialsDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)
	return cdb
}

func CreateCredentials(c echo.Context) (err error) {
	cdb := GetCredentialsDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	models := new(model.Credentials)
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

	newCredentials := model.NewCredentials{
		Name:         models.Name,
		Type:         models.Type,
		Domain:       models.Domain,
		Region:       models.Region,
		Url:          models.Url,
		Tenant:       models.Tenant,
		Access_id:    models.Access_id,
		Access_token: models.Access_token,
		Project:      models.Project,
		//Created_at: models.Created_at,
	}

	result, err := cdb.InsertOne(ctx, newCredentials)
	if err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)

		return nil
	}

	return c.JSON(http.StatusOK, result)
}

func ListCredentials(c echo.Context) (err error) {
	var results []model.Credentials
	cdb := GetCredentialsDB("credentials")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Credentials
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

func FindCredentials(c echo.Context) (err error) {
	var credentials model.Credentials
	//var credentials bson.M
	cdb := GetCredentialsDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("credentialsName")

	if err := cdb.FindOne(ctx, bson.M{"name": search_val}).Decode(&credentials); err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("credentials not found."))
		return nil
	} else {

		return c.JSON(http.StatusOK, &credentials)
	}

}

func DeleteCredentials(c echo.Context) (err error) {
	cdb := GetCredentialsDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("credentialsName")

	result, err := cdb.DeleteOne(ctx, bson.M{"name": search_val})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
		return
	}
	if result.DeletedCount == 0 {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Credentials not found."))
		return
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"status": http.StatusOK,
			"data":   search_val + " Credentials Deleted Complete",
		})
	}
}

func UpdateCredentials(c echo.Context) (err error) {
	cdb := GetRequestDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("credentialsName")

	models := new(model.RequestCredentials)
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
	// update = bson.M{"workspace" : workspaceObjectId2[0][0].Value.(primitive.ObjectID),"projectOwner": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectCreator": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectDescription": models.Description, "selectCluster":slice}
	// default :
	//update = bson.M{"workspace": workspaceObjectId2[0][0].Value.(primitive.ObjectID), "projectOwner": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectCreator": memberObjectId2[0][0].Value.(primitive.ObjectID), "projectDescription": models.Description, "selectCluster": slice}

	update = bson.M{"name": models.Name, "type": models.Type, "region": models.Region, "domain": models.Domain, "url": models.Url, "tenant": models.Tenant, "access_id": models.Access_id, "access_token": models.Access_token, "project": models.Project}
	fmt.Println(update)
	result, err := cdb.UpdateOne(ctx, bson.M{"name": search_val}, bson.M{"$set": update})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to update."))
		return
	}

	if result.MatchedCount == 1 {
		if err := cdb.FindOne(ctx, bson.M{"name": search_val}).Decode(&cdb); err != nil {
			common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to match Credentials."))
			return nil
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   search_val + " Updated Complete",
	})
}
