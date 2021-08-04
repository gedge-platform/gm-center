package api

import (
	"gmc_database_api_server/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Cluster(c echo.Context) (err error) {
	// model := new(model.JOB)

	// return c.JSON(200, echo.Map{"items": model})

	// url := "http://192.168.150.197:5555/kube/v1/cluster1/jobs"
	// if data := getData(c, url, false); data != "nf" {
	// 	n := gjson.Parse(data)
	// 	k := n.Get("items.#.metadata").Array()
	// 	log.Println("k is ", k)
	// }
	// replica := [new(model.REPLICA)]
	// var replica = {0,0,0,0,0}
	test := new(model.SERVICE_DETAIL)

	return c.JSON(http.StatusOK, echo.Map{
		"test": echo.Map{
			"items": test,
			// "items2": replica,
		},
	})

}
