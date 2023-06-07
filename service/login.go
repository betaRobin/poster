package service

import (
	"github.com/betaRobin/poster/database"
	"github.com/betaRobin/poster/entity"
	"github.com/betaRobin/poster/models/request"
)

func Login(request request.Login) bool {
	db := database.Connect()

	var user entity.User

	result := db.Where(&entity.User{Username: request.Username, Password: request.Password}).First(&user)

	return result.Error == nil
}
