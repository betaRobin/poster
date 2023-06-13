package response

import (
	"time"

	"github.com/betarobin/poster/entity"
)

type GetPostsRequest struct {
	PostId      string     `json:"post_id"`
	Username    string     `json:"username"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func Posts(posts []entity.Post) *[]GetPostsRequest {
	var response = []GetPostsRequest{}

	if posts == nil {
		return &response
	}

	for _, post := range posts {
		response = append(response, GetPostsRequest{
			PostId:      post.Id.String(),
			Username:    post.User.Username,
			Title:       post.Title,
			Description: post.Description,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		})
	}
	return &response
}
