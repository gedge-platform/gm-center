package controller

import (
	"context"
	"fmt"
	"gmc_api_gateway/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func TotalDashboard(c echo.Context) (err error) {
	clusters := GetClusterDB("cluster")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	ClusterCount, err := clusters.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	workspaces := GetClusterDB("workspace")
	workspaceCount, err := workspaces.CountDocuments(context.Background(), bson.D{})
	projects := GetClusterDB("project")
	projectCount, err := projects.CountDocuments(context.Background(), bson.D{})
	cursor, err := clusters.Find(context.TODO(), bson.D{{"clusterType", "core"}})
	if err != nil {
		log.Fatal(err)
	}

	var coreClouds []bson.M

	if err = cursor.All(ctx, &coreClouds); err != nil {
		log.Fatal(err)
	}
	cursor2, err := clusters.Find(context.TODO(), bson.D{{"clusterType", "edge"}})
	if err != nil {
		log.Fatal(err)
	}
	var edgeClouds []bson.M
	if err = cursor2.All(ctx, &edgeClouds); err != nil {
		log.Fatal(err)
	}
	dashbaordData := model.TOTAL_DASHBOARD{
		ClusterCnt:     ClusterCount,
		CoreClusterCnt: len(coreClouds),
		EdgeClusterCnt: len(edgeClouds),
		ProjectCnt:     projectCount,
		WorkspaceCnt:   workspaceCount,
		PodCpuTop5:     dashboard_pod_monit("all", dashboardMetric["dashboard_cpu_used_podList"]),
		PodMemTop5:     dashboard_pod_monit("all", dashboardMetric["dashboard_mem_used_podList"]),
		ClusterCpuTop5: dashboard_cluster_monit("all", dashboardMetric["dashboard_cpu_used_clusterList"]),
		ClusterMemTop5: dashboard_cluster_monit("all", clusterMetric["memory_usage"]),
		EdgeCloud:      edgeClouds,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": dashbaordData,
	})
}
func TotalDashboard(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "namespaces",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		User:      c.QueryParam("user"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": dashbaordData,
	})
}

func GeoCoder(add string) (result string) {

	key := os.Getenv("GEO_KEY")
	baseUrl := "http://api.vworld.kr/req/address?service=address&request=getcoord&version=2.0&crs=epsg:4326&refine=true&simple=true&format=json&type=road&errorFormat=json"

	adderss := url.QueryEscape(add)
	httpUrl := baseUrl + "&address=" + adderss + "&key=" + key
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, httpUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	body_str := string(body)
	return body_str
}
