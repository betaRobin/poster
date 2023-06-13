package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:current_timestamp()"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"default:null"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"default:null"`
	UserID      uuid.UUID  `json:"user_id"`
	User        User       `gorm:"foreignKey:UserID"`
}

func NewPost(title string, description string, userId uuid.UUID) *Post {
	postId, _ := uuid.NewRandom()
	newPost := Post{
		ID:          postId,
		Title:       title,
		Description: description,
		UserID:      userId,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		DeletedAt:   nil,
	}
	return &newPost
}
