package entity

import (
	"github.com/google/uuid"
)

type Post struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      uuid.UUID `json:"user-id"`
	User        User
}
