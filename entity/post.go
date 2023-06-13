package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:current_timestamp()"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"default:null"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"default:null"`
	UserId      uuid.UUID  `json:"user_id"`
	User        User
}

func NewPost(title string, description string, userId uuid.UUID) *Post {
	postId, _ := uuid.NewRandom()
	newPost := Post{
		Id:          postId,
		Title:       title,
		Description: description,
		UserId:      userId,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		DeletedAt:   nil,
	}
	return &newPost
}
