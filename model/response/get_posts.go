package response

import (
	"github.com/betarobin/poster/entity"
)

type EditPostRequest struct {
	PostId      string `json:"post_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Posts(posts []entity.Post) *[]EditPostRequest {
	var response []EditPostRequest
	for _, post := range posts {
		response = append(response, EditPostRequest{
			PostId:      post.Id.String(),
			Title:       post.Title,
			Description: post.Description,
		})
	}
	return &response
}
