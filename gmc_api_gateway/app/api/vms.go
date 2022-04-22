package api

import (
	"context"
	db "gmc_api_gateway/app/database"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCredentialsDB(name string) *mongo.Collection {
	db := db.DbManager()
	cdb := db.Collection(name)
	return cdb
}
func GetVms(c echo.Context) (err error) {
	var showVm []bson.M
	cdb := GetCredentialsDB("credentials")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	findOptions := options.Find()

	cur, err := cdb.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		//lookupCluster := bson.D{{"$lookup", bson.D{{"from", "cluster"}, {"localField", "selectCluster"}, {"foreignField", "_id"}, {"as", "selectCluster"}}}}
		lookupProject := bson.D{{"$lookup", bson.D{{"from", "project"}, {"localField", "project"}, {"foreignField", "_id"}, {"as", "project"}}}}

		showProjectCursor, err := cdb.Aggregate(ctx, mongo.Pipeline{lookupProject})

		if err = showProjectCursor.All(ctx, &showVm); err != nil {
			panic(err)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return c.JSON(http.StatusOK, echo.Map{"data": showVm})
}
