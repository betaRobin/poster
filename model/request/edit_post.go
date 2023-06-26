package request

type EditPostRequest struct {
	Type    string      `json:"type"`
	Title   string      `json:"title"`
	PostID  string      `json:"post_id"`
	Content interface{} `json:"content"`
}
