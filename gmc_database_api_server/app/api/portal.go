package api

import (
	"fmt"
	"gmc_database_api_server/app/common"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTest(c echo.Context) (err error) {

	data, err := common.GetModel(c, "pods")
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	fmt.Printf("data type is %s\n", common.Typeof(data))

	data1, _ := common.FilterStr(data, "metadata.ownerReferences")
	fmt.Printf("data1 type is %s\n", common.Typeof(data1))

	data2 := common.Filter(data, "metadata.ownerReferences")
	// fmt.Printf("data2 type is %s\n", common.Typeof(data2))

	log.Println("[#2] data is", data2)
	log.Println("[#2-1] data is", data2["name"])
	data3, _ := common.Finding(data, "metadata.ownerReferences", "blockOwnerDeletion")
	fmt.Printf("data3 type is %s\n", common.Typeof(data3))

	log.Println("[#3] data is", data3)
	data4, _ := common.Finding(data, "spec.volumes", "name")
	fmt.Printf("data4 type is %s\n", common.Typeof(data4))

	log.Println("[#4] data is", data4)

	return c.JSON(http.StatusOK, echo.Map{
		"testData":  data,
		"testData1": data1,
		"testData2": data2,
		"testData3": data3,
		"testData4": data4,
	})
}
