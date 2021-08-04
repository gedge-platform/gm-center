package api

import (
	"gmc_database_api_server/app/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Cluster(c echo.Context) (err error) {
	// model := new(model.JOB)

	// return c.JSON(200, echo.Map{"items": model})

	url := "http://192.168.150.197:5555/kube/v1/cluster1/jobs"
	if data := getData(c, url, false); data != "nf" {
		n := gjson.Parse(data)
		k := n.Get("items.#.metadata").Array()
		log.Println("k is ", k)
	}

	var jobModel model.JOB
	jobModel.Workspace = "test"
	// jobModel.Image = "dddfgfgg"

	return c.JSON(http.StatusOK, echo.Map{
		"items": echo.Map{
			"items":  jobModel,
			"items2": new(model.POD),
			// "OwnerReference": new(model.OwnerReference),
			// "Jobdetails":     new(model.JOBDETAIL),
		},
		"test": echo.Map{
			"items": new(model.CRONJOBDETAIL),
			// "Jobdetails": new(model.JOBDETAIL),
		},
	})

}
