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
	if models.Type == "edge" {
		point := GeoCoder(models.Address)
		Point := make(map[string]string)
		Point["x"] = common.InterfaceToString(common.FindData(point, "response.result.point", "x"))
		Point["y"] = common.InterfaceToString(common.FindData(point, "response.result.point", "y"))
		models.Point = Point
	}
	result, err := cdb.InsertOne(ctx, models)
	if err != nil {
		common.ErrorMsg(c, http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, result)
}

func ListCluster(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		User:      c.QueryParam("user"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	var results []model.CLUSTER
	var results2 []model.CLUSTER
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
		var cluster model.CLUSTER
		cluster.Cluster = elem
		results = append(results, cluster)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	for _, cluster := range results {
		params.Name = cluster.Name
		gpuList, check := GpuCheck(cluster.Name)
		fmt.Println("gpuList : ", gpuList)
		fmt.Println("check : ", check)
		if check == true {
			cluster.GpuCnt = len(gpuList)
		} else {
			cluster.GpuCnt = 0
		}
		params.Cluster = cluster.Name
		params.Name = ""
		getData, err := common.DataRequest(params)
		if err != nil {
			common.ErrorMsg(c, http.StatusNotFound, err)
			return nil
		}
		cluster.NodeCnt = common.InterfaceOfLen(common.FindData(getData, "items", ""))
		tempMetric := []string{"cpu_usage", "memory_usage", "pod_running"}
		tempresult := NowMonit("cluster", params.Cluster, params.Name, tempMetric)
		cluster.ResourceUsage = tempresult
		results2 = append(results2, cluster)
	}
	return c.JSON(http.StatusOK, results2)
}

// GetCluster godoc
// @Summary Show detail cluster
// @Description get cluster Details
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CLUSTER
// @Security   Bearer
// @Param name path string true "name of the Cluster"
// @Router /cluster/{name} [get]
func FindCluster(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "nodes",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		User:      c.QueryParam("user"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	var cluster model.Cluster
	cdb := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := params.Name

	if err := cdb.FindOne(ctx, bson.M{"clusterName": search_val}).Decode(&cluster); err != nil {
		common.ErrorMsg(c, http.StatusNotFound, errors.New("Cluster not found."))
		return nil
	} else {
		Cluster := model.CLUSTER_DETAIL{
			Cluster: cluster,
		}
		gpuList, check := GpuCheck(params.Name)
		if check != false {
			Cluster.Gpu = gpuList
		} else {
			Cluster.Gpu = nil
		}
		params.Cluster = params.Name
		params.Name = ""
		getData, err := common.DataRequest(params)
		if err != nil {
			common.ErrorMsg(c, http.StatusNotFound, err)
			return nil
		}
		Nodes, _ := common.FindDataLabelKey(getData, "items", "labels", "node-role.kubernetes.io/master")
		var NodeList []model.NODE
		for n, _ := range Nodes {
			Node := model.NODE{
				Name:     common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
				NodeType: common.InterfaceToString(common.FindData(Nodes[n], "nodeType", "")),

				CreateAt:    common.InterfaceToTime(common.FindData(Nodes[n], "metadata", "creationTimestamp")),
				Version:     common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "kubeletVersion")),
				Label:       common.FindData(Nodes[n], "metadata", "labels"),
				Annotation:  common.FindData(Nodes[n], "metadata", "annotations"),
				Allocatable: common.FindData(Nodes[n], "status", "allocatable"),
				IP:          common.InterfaceToString(common.FindData(Nodes[n], "status", "addresses.0.address")),
				// Status                   : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
				// Network                 : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
				Os:                      common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "operatingSystem")) + " / " + common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "osImage")),
				Kernel:                  common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "kernelVersion")),
				ContainerRuntimeVersion: common.InterfaceToString(common.FindData(Nodes[n], "status.nodeInfo", "containerRuntimeVersion")),
				Capacity:                common.FindData(Nodes[n], "status", "capacity"),
				// Addresses  : common.InterfaceToString(common.FindData(Nodes[n], "metadata", "name")),
			}
			NodeList = append(NodeList, Node)
		}
		ResourceCnt := model.Resource_cnt{
			DeploymentCount:  ResourceCnt(params, "deployments"),
			DaemonsetCount:   ResourceCnt(params, "daemonsets"),
			StatefulsetCount: ResourceCnt(params, "statefulsets"),
			PodCount:         ResourceCnt(params, "deployments"),
			ServiceCount:     ResourceCnt(params, "pods"),
			CronjobCount:     ResourceCnt(params, "cronjobs"),
			JobCount:         ResourceCnt(params, "jobs"),
			VolumeCount:      ResourceCnt(params, "persistentvolumeclaims"),
		}

		Cluster.Resource = ResourceCnt
		Cluster.Nodes = NodeList
		Cluster.Events = getCallEvent(params)
		return c.JSON(http.StatusOK, Cluster)
	}
}
func ListClusterDB(name string) []model.Cluster {
	var results []model.Cluster
	cdb := GetClusterDB(name)

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

	return results
}
func FindClusterDB(value string) *model.Cluster {
	var cluster model.Cluster
	cdb := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	search_val := value

	if err := cdb.FindOne(ctx, bson.M{"clusterName": search_val}).Decode(&cluster); err != nil {
		// common.ErrorMsg(c, http.StatusNotFound, errors.New("Cluster not found."))
		return nil
	} else {
		fmt.Println("[####cluster] : ", &cluster)
		return &cluster
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
