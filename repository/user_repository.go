package repository

import (
	"log"

	"github.com/betarobin/poster/database"
	"github.com/betarobin/poster/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func VerifyUser(username string, password string) (*entity.User, *gorm.DB) {
	var user entity.User
	db := database.Connect()
	result := db.Where(&entity.User{Username: username, Password: password}).First(&user)

	if result.Error != nil {
		log.Println("[VerifyUser] Failed to verify User")
	}

	return &user, result
}

func InsertUser(username string, password string) (*entity.User, *gorm.DB) {
	user := entity.NewUser(username, password)
	db := database.Connect()
	result := db.Create(&user)

	if result.Error != nil {
		log.Println("[InsertUser] Failed to insert User")
	}

	return user, result
}

func FindUserByUsername(username string) (*entity.User, *gorm.DB) {
	var user entity.User
	db := database.Connect()
	result := db.Where(&entity.User{Username: username}).First(&user)

	if result.Error != nil {
		log.Printf("[FindUserByUsername] Failed to query User with username %s\n", username)
	}

	return &user, result
}

func FindUserById(userUUID uuid.UUID) (*entity.User, *gorm.DB) {
	var user entity.User
	db := database.Connect()
	result := db.Where(&entity.User{ID: userUUID}).First(&user)

	if result.Error != nil {
		log.Printf("[FindUserById] Failed to query User with id %s\n", userUUID.String())
	}

	return &user, result
}
