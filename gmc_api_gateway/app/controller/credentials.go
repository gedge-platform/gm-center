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
	cdb2 := GetCredentialsDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	models := new(model.Credentials)
	validate := validator.New()

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)

		return nil
	}

	projectObjectId, err := cdb2.Find(ctx, bson.M{"objectId": models.Project})
	var projectObjectId2 []bson.D

	if err = projectObjectId.All(ctx, &projectObjectId2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("1", projectObjectId)
	fmt.Println("2", projectObjectId2)

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
		//Project:      projectObjectId2[0][0].Value.(primitive.ObjectID),
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
	var showsCredentials []bson.M
	cdb := GetProjectDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupProject := bson.D{{"$lookup", bson.D{{"from", "project"}, {"localField", "selectProject"}, {"foreignField", "_id"}, {"as", "sProject"}}}}

		showCredentialsCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupProject})

		fmt.Println(showCredentialsCursor)
		if err = showCredentialsCursor.All(ctx, &showsCredentials); err != nil {
			panic(err)
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, showsCredentials)
}

func FindCredentials(c echo.Context) (err error) {
	var showsCredentials []bson.M
	cdb := GetCredentialsDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("credentialsName")

	fmt.Println("--", search_val)
	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		lookupProject := bson.D{{"$lookup", bson.D{{"from", "project"}, {"localField", "project"}, {"foreignField", "_id"}, {"as", "sproject"}}}}
		matchProject := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "name", Value: search_val},
			}},
		}
		fmt.Println("1", matchProject)
		fmt.Println("1", bson.D{})
		showLoadedCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupProject, matchProject})

		if err = showLoadedCursor.All(ctx, &showsCredentials); err != nil {
			panic(err)
		}
		fmt.Println(showsCredentials)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	if showsCredentials == nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Credentials not found."))
		return
	} else {
		return c.JSON(http.StatusOK, showsCredentials)
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
	/* 	cdb2 := GetCredentialsDB("cluster")
	   	cdb3 := GetCredentialsDB("member")
	   	cdb4 := GetCredentialsDB("workspace") */
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("credentialsName")

	models := new(model.RequestCredentials)
	validate := validator.New()

	if err = c.Bind(models); err != nil {
		common.ErrorMsg(c, http.StatusBadRequest, err)
		return nil
	}
	/* 	memberObjectId, err := cdb3.Find(ctx, bson.M{"memberName": models.MemberName})
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
	   	} */

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

	fmt.Println(models)

	update = bson.M{"name": models.Name, "type": models.Type, "domain": models.Domain}
	fmt.Println(update)
	result, err := cdb.UpdateOne(ctx, bson.M{"name": search_val}, bson.M{"$set": update})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to update."))
		return
	}

	if result.MatchedCount == 1 {
		if err := cdb.FindOne(ctx, bson.M{"credentialsName": search_val}).Decode(&cdb); err != nil {
			common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to match Credentials."))
			return nil
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   search_val + " Updated Complete",
	})
}
