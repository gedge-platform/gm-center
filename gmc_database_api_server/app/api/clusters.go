package api

import (
	"net/http"
	"strings"

	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"
	
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetAllClusters(c echo.Context) (err error) {
	var msgError messageFormat
	db := db.DbManager()
	models := []model.Cluster{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		msgError.StatusCode = http.StatusOK
		msgError.Message = "No Data"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func GetCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := FindClusterDB(db, "Name", search_val)

	if models == nil {
		var msgError messageFormat
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func CreateCluster(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.Cluster)
	var msgError messageFormat

	if err = c.Bind(models); err != nil {
		msgError.StatusCode = http.StatusBadRequest
		msgError.Message = "Bad Request"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}
	if err = c.Validate(models); err != nil {
		msgError.StatusCode = http.StatusUnprocessableEntity
		msgError.Message = "The given data was invalid."
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	if err != nil {
		panic(err)
	}

	if err := db.Create(&models).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": models})
}

func UpdateCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := model.Cluster{}
	var msgError messageFormat

	if err := c.Bind(&models); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := FindClusterDB(db, "Name", search_val); err == nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}  else {
		models.Name = search_val
	}


	models2 := FindClusterDB(db, "Name", search_val)

	if models.Ip != "" { models2.Ip = models.Ip	} 
	// if models.extIp != "" { models2.extIp = models.extIp }
	if models.Role != "" { models2.Role = models.Role	} 
	if models.Type != "" { models2.Type = models.Type	} 
	if models.Endpoint != "" { models2.Endpoint = models.Endpoint } 
	if models.Creator != "" { models2.Creator = models.Creator } 
	if models.State != "" { models2.State = models.State } 
	if models.Version != "" { models2.Version = models.Version }


	if err := db.Save(&models2).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusExpectationFailed, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models2})
}


func DeleteCluster(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	// models := model.Cluster{}
	var msgError messageFormat

	if err := FindClusterDB(db, "Name", search_val); err == nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}
	
	models := FindClusterDB(db, "Name", search_val)

	if err := db.Delete(&models).Error; err != nil {
		msgError.StatusCode = http.StatusInternalServerError
		msgError.Message = "Internal Server Error"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusInternalServerError, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func FindClusterDB(db *gorm.DB, select_val string, search_val string) *model.Cluster {
	models := model.Cluster{}
	if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.Cluster{Name: search_val}).Error; err != nil {
			return nil
		}
	}
	return &models
}