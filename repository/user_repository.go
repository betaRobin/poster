package repository

import (
	"github.com/betarobin/poster/database"
	"github.com/betarobin/poster/entity"
	"gorm.io/gorm"
)

func VerifyUser(username string, password string) *gorm.DB {
	db := database.Connect()

	var user entity.User

	return db.Where(&entity.User{Username: username, Password: password}).First(&user)
}
