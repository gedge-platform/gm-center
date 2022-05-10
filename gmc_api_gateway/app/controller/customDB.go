package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDBProjectList(params model.PARAMS, obj primitive.ObjectID, search_type string) []bson.M {
	cdb := GetClusterDB("project")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := obj

	cursor, err := cdb.Find(context.TODO(), bson.D{{search_type, search_val}})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	return results

}

func DuplicateCheckDB(c echo.Context) (err error) {
	search_val := c.Param("name")
	// var models interface{}
	if c.QueryParam("type") == "cluster" {
		cdb := GetWorkspaceDB("cluster")
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		models := model.Cluster{}
		if err := cdb.FindOne(ctx, bson.M{"clusterName": search_val}).Decode(&models); err != nil {
			common.ErrorMsg(c, http.StatusOK, common.ErrDuplicatedCheckOK)
		} else {
			common.ErrorMsg(c, http.StatusBadRequest, common.ErrDuplicated)
		}
	} else if c.QueryParam("type") == "project" {
		models := model.Project{}
		cdb := GetWorkspaceDB("project")
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		if err := cdb.FindOne(ctx, bson.M{"projectName": search_val}).Decode(&models); err != nil {
			common.ErrorMsg(c, http.StatusOK, common.ErrDuplicatedCheckOK)
		} else {
			common.ErrorMsg(c, http.StatusBadRequest, common.ErrDuplicated)
		}
	} else if c.QueryParam("type") == "workspace" {
		models := model.Workspace{}
		cdb := GetWorkspaceDB("workspace")
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		if err := cdb.FindOne(ctx, bson.M{"workspaceName": search_val}).Decode(&models); err != nil {
			common.ErrorMsg(c, http.StatusOK, common.ErrDuplicatedCheckOK)
		} else {
			common.ErrorMsg(c, http.StatusBadRequest, common.ErrDuplicated)
		}
	} else if c.QueryParam("type") == "member" {
		models := model.Workspace{}
		cdb := GetWorkspaceDB("member")
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		if err := cdb.FindOne(ctx, bson.M{"memberName": search_val}).Decode(&models); err != nil {
			common.ErrorMsg(c, http.StatusOK, common.ErrDuplicatedCheckOK)
		} else {
			common.ErrorMsg(c, http.StatusBadRequest, common.ErrDuplicated)
		}
	} else {
		common.ErrorMsg(c, http.StatusBadRequest, common.ErrTypeNotFound)
	}
	return nil
}
