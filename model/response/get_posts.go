package response

import (
	"time"

	"github.com/betarobin/poster/entity"
	contenthelper "github.com/betarobin/poster/helper/content"
)

type GetPostsResponse struct {
	PostId    string      `json:"post_id"`
	Type      string      `json:"type"`
	Title     string      `json:"title"`
	Content   interface{} `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
}

func Posts(posts *[]entity.Post) []*GetPostsResponse {
	var response = []*GetPostsResponse{}

	if posts == nil {
		return response
	}

	for _, post := range *posts {
		postResponse := &GetPostsResponse{
			PostId:    post.ID.String(),
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		content, err := contenthelper.ParseContent(post.Type, post.Content)

		if err != nil {
			continue
		} else {
			postResponse.Content = content
		}

		response = append(response, postResponse)
	}
	return response
}
