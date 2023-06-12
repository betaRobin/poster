package repository

import (
	"log"

	"github.com/betarobin/poster/database"
	"github.com/betarobin/poster/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InsertPost(userId string, title string, description string) (*entity.Post, *gorm.DB) {
	post := entity.Post{Title: title, Description: description, UserId: uuid.MustParse(userId)}
	db := database.Connect()
	result := db.Create(&post)

	if result.Error != nil {
		log.Println("[InsertPost] Failed to insert post")
	}
	log.Println(post)

	return &post, result
}
