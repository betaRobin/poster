package request

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
