package request

import "github.com/google/uuid"

type EditPostRequest struct {
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PostId      uuid.UUID `json:"post_id"`
}
