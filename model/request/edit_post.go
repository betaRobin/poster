package request

type EditPostRequest struct {
	Type    string  `json:"type"`
	PostID  string  `json:"post_id"`
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}
