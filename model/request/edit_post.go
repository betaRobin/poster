package request

type EditPostRequest struct {
	PostID      string  `json:"post_id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
