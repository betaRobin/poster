package entity

import "github.com/google/uuid"

type Post struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Username    string    `json:"username"`
}
