package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Type      string     `json:"type"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:current_timestamp()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
	UserID    uuid.UUID  `json:"-"`
	User      User       `gorm:"foreignKey:UserID"`
}

func NewPost(userId uuid.UUID, postType string, title string, content string) *Post {
	postId, _ := uuid.NewRandom()
	newPost := Post{
		ID:        postId,
		UserID:    userId,
		Type:      postType,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
	return &newPost
}
