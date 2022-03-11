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
	fmt.Println("models : ", models)

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
	// findOptions.SetLimit(5)

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
	search_val := c.Param("id")

	if err := cdb.FindOne(ctx, bson.M{"memberId": search_val}).Decode(&member); err != nil {
		// common.ErrorMsg(c, http.StatusNotFound, err)
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
	search_val := c.Param("id")

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

// func UpdateMember(c echo.Context) (err error) {
// 	cdb := GetDB("member")
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

// 	models := new(model.Member)

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

// 	result, err := cdb.UpdateOne(ctx, bson.M{"memberId": search_val})
// 	if err != nil {
// 		common.ErrorMsg(c, http.StatusNotFound, errors.New("failed to delete."))
// 		return
// 	}
// 	if result.DeletedCount == 0 {
// 		common.ErrorMsg(c, http.StatusNotFound, errors.New("Member not found."))
// 		return
// 	} else {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"status": http.StatusOK,
// 			"data":   search_val + " Deleted",
// 		})
// 	}
// }
