package Post

type PostCreationDTO struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}
