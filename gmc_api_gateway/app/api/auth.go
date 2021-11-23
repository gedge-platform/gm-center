package api

import (
	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"
	"strings"
)

func AuthenticateUser(id, password string) bool {
	db := db.DbManager()
	var user model.MemberWithPassword

	idCheck := strings.Compare(id, "") != 0
	passCheck := strings.Compare(password, "") != 0

	if idCheck && passCheck {

		if err := db.First(&user, model.MemberWithPassword{Member: model.Member{Id: id}, Password: password}).Error; err == nil {
			return true
		}

	}

	return false
}
