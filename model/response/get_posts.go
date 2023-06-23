package response

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/betarobin/poster/entity"
	typepost "github.com/betarobin/poster/enum/type_post"
	"github.com/betarobin/poster/model/content"
)

type GetPostsResponse struct {
	PostId    string      `json:"post_id"`
	Type      string      `json:"type"`
	Title     string      `json:"title"`
	Content   interface{} `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
}

func Posts(posts *[]entity.Post) map[string][]*GetPostsResponse {
	var response = []*GetPostsResponse{}

	if posts == nil {
		return map[string][]*GetPostsResponse{"posts": response}
	}

	for _, post := range *posts {
		postResponse := &GetPostsResponse{
			PostId:    post.ID.String(),
			Type:      post.Type,
			Title:     post.Title,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		contentJsonString := post.Content

		switch post.Type {
		case typepost.Text:
			text := &content.Text{}
			err := json.Unmarshal([]byte(contentJsonString), text)
			if err == nil {
				postResponse.Content = text
			} else {
				continue
			}
		case typepost.Image:
			image := &content.Image{}
			err := json.Unmarshal([]byte(contentJsonString), image)
			if err == nil {
				postResponse.Content = image
			} else {
				continue
			}
		case typepost.Checklist:
			checklist := &content.Checklist{}
			err := json.Unmarshal([]byte(contentJsonString), checklist)
			if err == nil {
				postResponse.Content = checklist
			} else {
				continue
			}
		default:
			continue
		}

		fmt.Println(contentJsonString)

		response = append(response, postResponse)
	}
	return map[string][]*GetPostsResponse{"posts": response}
}
