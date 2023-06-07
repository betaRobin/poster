package request

import "github.com/google/uuid"

type DeletePostRequest struct {
	Username string    `json:"username"`
	PostId   uuid.UUID `json:"post_id"`
}
