package repository

import (
	"log"

	"github.com/betarobin/poster/database"
	"github.com/betarobin/poster/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InsertPost(userId uuid.UUID, postType string, title string, content string) (*entity.Post, *gorm.DB) {
	post := entity.NewPost(userId, postType, title, content)
	db := database.Connect()
	result := db.Create(post)

	if result.Error != nil {
		log.Println("[InsertPost] Failed to insert post")
	}

	return post, result
}

func GetPostsByUserId(userId uuid.UUID) (*[]entity.Post, *gorm.DB) {
	var posts []entity.Post
	db := database.Connect()
	result := db.Model(&entity.Post{}).Preload("User").Where("user_id = ?", userId).Find(&posts)

	if result.Error != nil {
		log.Println("[GetPostsByUserId] Failed to get posts")
	}

	return &posts, result
}

func GetPostById(postId uuid.UUID) (*entity.Post, *gorm.DB) {
	var post entity.Post
	db := database.Connect()
	result := db.Model(&entity.Post{}).Preload("User").Where("id = ?", postId).Find(&post)

	if result.Error != nil {
		log.Println("[GetPostsByUserId] Failed to get posts")
	}

	return &post, result
}

func EditPostContent(post entity.Post) *gorm.DB {
	db := database.Connect()
	result := db.Save(&post)

	if result.Error != nil {
		log.Println("[EditPostContent] Failed to edit post")
	}

	return result
}

func DeletePostById(postId uuid.UUID) *gorm.DB {
	db := database.Connect()
	result := db.Delete(&entity.Post{}, postId)

	if result.Error != nil {
		log.Println("[DeletePostById] Failed to delete post")
	}

	return result
}
