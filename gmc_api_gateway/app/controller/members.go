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

func GetDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)

	return cdb
}

func CreateMember(c echo.Context) (err error) {
	cdb := GetDB("member")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	models := new(model.Member)
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

func ListMember(c echo.Context) (err error) {
	var results []model.Member
	cdb := GetDB("member")

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Member
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		elem.Password = "******"
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, results)
}

func FindMember(c echo.Context) (err error) {
	var member model.Member
	cdb := GetDB("member")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("memberId")

	if err := cdb.FindOne(ctx, bson.M{"memberId": search_val}).Decode(&member); err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Member not found."))
		return nil
	} else {
		member.Password = "******"
		return c.JSON(http.StatusOK, &member)
	}
}

func DeleteMember(c echo.Context) (err error) {
	cdb := GetDB("member")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("memberId")

	result, err := cdb.DeleteOne(ctx, bson.M{"memberId": search_val})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
		return
	}
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

func UpdateMember(c echo.Context) (err error) {
	cdb := GetDB("member")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := c.Param("memberId")

	models := new(model.RequestMember)
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
	// update = bson.M{"memberName": models.Name, "email": models.Email, "password":models.Password, "contact": models.Contact, "memberRole": models.RoleName}
	// default :
	update = bson.M{"memberName": models.Name, "email": models.Email, "password": models.Password, "contact": models.Contact, "memberRole": models.RoleName}
	// }

	result, err := cdb.UpdateOne(ctx, bson.M{"memberId": search_val}, bson.M{"$set": update})
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to update."))
		return
	}

	if result.MatchedCount == 1 {
		if err := cdb.FindOne(ctx, bson.M{"memberId": search_val}).Decode(&cdb); err != nil {
			common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to match Member."))
			return nil
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   search_val + " Updated Complete",
	})
}

func FindDBwithPW(select_val string, search_val string) *model.MemberWithPassword {
	var models model.MemberWithPassword
	cdb := GetDB("member")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	if err := cdb.FindOne(ctx, bson.M{"memberId": search_val}).Decode(&models); err != nil {
		return nil
	} else if err := cdb.FindOne(ctx, bson.M{"memberName": search_val}).Decode(&models); err != nil {
		return nil
	}
	return &models
}

func FindMemberDB(params model.PARAMS) model.Member {
	var member model.Member
	cdb := GetClusterDB("member")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := params.User

	if err := cdb.FindOne(ctx, bson.M{"memberId": search_val}).Decode(&member); err != nil {
	}
	return member
}
