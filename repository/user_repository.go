package repository

import (
	"log"

	"github.com/betarobin/poster/database"
	"github.com/betarobin/poster/entity"
	"gorm.io/gorm"
)

func VerifyUser(username string, password string) (entity.User, *gorm.DB) {
	var user entity.User
	db := database.Connect()
	result := db.Where(&entity.User{Username: username, Password: password}).First(&user)

	if result.Error != nil {
		log.Println("[VerifyUser] Failed to verify User")
	}
	log.Println(user)
	return user, result
}

func InsertUser(username string, password string) (entity.User, *gorm.DB) {
	user := entity.User{Username: username, Password: password}
	db := database.Connect()
	result := db.Create(&user)

	if result.Error != nil {
		log.Println("[InsertUser] Failed to insert User")
	}

	return user, result
}

func FindUserByUsername(username string) (entity.User, *gorm.DB) {
	var user entity.User
	db := database.Connect()
	result := db.Where(&entity.User{Username: username}).First(&user)

	if result.Error != nil {
		log.Printf("[FindUserByUsername] Failed to query User with username %s\n", username)
	}

	return user, result
}
