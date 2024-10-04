package Comment

type CommentCreationDTO struct {
	PostId   string `json:"post_id"`
	Username string `json:"username"`
	Body     string `json:"body"`
}
