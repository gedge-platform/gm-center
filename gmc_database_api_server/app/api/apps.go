package api

import (
	"net/http"
	"strings"

	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"
	
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetAllApps(c echo.Context) (err error) {
	var msgError messageFormat
	db := db.DbManager()
	models := []model.App{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		msgError.StatusCode = http.StatusOK
		msgError.Message = "No Data"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func GetApp(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := FindAppDB(db, "Name", search_val)

	if models == nil {
		var msgError messageFormat
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func CreateApp(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.App)
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

func UpdateApp(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	models := model.App{}
	var msgError messageFormat

	if err := c.Bind(&models); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := FindAppDB(db, "Name", search_val); err == nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	} else {
		models.Name = search_val
	}

	models2 := FindAppDB(db, "Name", search_val)

	if models.Name != "" { models2.Name = models.Name	} 
	if models.Description != "" { models2.Description = models.Description	} 
	if models.Category != "" { models2.Category = models.Category } 
	if models.Installed != models2.Installed { models2.Installed = models.Installed } 

	if err := db.Save(&models2).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusExpectationFailed, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models2})
}


func DeleteApp(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("name")
	// models := model.App{}
	var msgError messageFormat

	if err := FindAppDB(db, "Name", search_val); err == nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}
	
	models := FindAppDB(db, "Name", search_val)

	if err := db.Delete(&models).Error; err != nil {
		msgError.StatusCode = http.StatusInternalServerError
		msgError.Message = "Internal Server Error"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusInternalServerError, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func FindAppDB(db *gorm.DB, select_val string, search_val string) *model.App {
	models := model.App{}
	if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.App{Name: search_val}).Error; err != nil {
			return nil
		}
	}
	return &models
}