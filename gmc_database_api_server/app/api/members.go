package api

import (
	"net/http"
	"strings"

	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"
	
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetAllMembers(c echo.Context) (err error) {
	var msgError messageFormat
	db := db.DbManager()
	models := []model.Member{}
	db.Find(&models)

	if db.Find(&models).RowsAffected == 0 {
		msgError.StatusCode = http.StatusOK
		msgError.Message = "No Data"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func GetMember(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("id")
	models := FindDB(db, "Id", search_val)

	if models == nil {
		var msgError messageFormat
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func CreateMember(c echo.Context) (err error) {
	db := db.DbManager()
	models := new(model.MemberWithPassword)
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

func UpdateMember(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("id")
	models := model.MemberWithPassword{}
	var msgError messageFormat

	if err := c.Bind(&models); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := FindDBwithPW(db, "Id", search_val); err == nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	} else {
		models.Id = search_val
	}

	models2 := FindDBwithPW(db, "Id", search_val)

	if models.Name != "" { models2.Name = models.Name	} 
	if models.Email != "" { models2.Email = models.Email	} 
	if models.Contact != "" { models2.Contact = models.Contact } 
	if models.Description != "" { models2.Description = models.Description } 
	if models.Password != "" { models2.Password = models.Password	}

	if err := db.Save(&models2).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusExpectationFailed, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models2})
}


func DeleteMember(c echo.Context) (err error) {
	db := db.DbManager()
	search_val := c.Param("id")
	// models := model.Member{}
	var msgError messageFormat

	if err := FindDB(db, "Id", search_val); err == nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}
	
	models := FindDB(db, "Id", search_val)

	if err := db.Delete(&models).Error; err != nil {
		msgError.StatusCode = http.StatusInternalServerError
		msgError.Message = "Internal Server Error"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusInternalServerError, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": models})
}

func FindDB(db *gorm.DB, select_val string, search_val string) *model.Member {
	models := model.Member{}
	if strings.Compare(select_val, "Id") == 0 {
		if err := db.First(&models, model.Member{Id: search_val}).Error; err != nil {
			return nil
		}
	} else if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.Member{Name: search_val}).Error; err != nil {
			return nil
		}
	}
	return &models
}

func FindDBwithPW(db *gorm.DB, select_val string, search_val string) *model.MemberWithPassword {
	models := model.MemberWithPassword{}
	if strings.Compare(select_val, "Id") == 0 {
		if err := db.First(&models, model.MemberWithPassword{Member: model.Member{Id: search_val}}).Error; err != nil {
			return nil
		}
	} else if strings.Compare(select_val, "Name") == 0 {
		if err := db.First(&models, model.MemberWithPassword{Member: model.Member{Name: search_val}}).Error; err != nil {
			return nil
		}
	}
	return &models
}
