package service

import (
	"github.com/betarobin/poster/database"
	"github.com/betarobin/poster/entity"
	"github.com/betarobin/poster/models/request"
)

func Login(request request.Login) bool {
	db := database.Connect()

	var user entity.User

	result := db.Where(&entity.User{Username: request.Username, Password: request.Password}).First(&user)

	return result.Error == nil
}
