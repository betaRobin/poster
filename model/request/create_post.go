package request

type CreatePostRequest struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
