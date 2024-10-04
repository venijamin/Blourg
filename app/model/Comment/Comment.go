package Comment

type Comment struct {
	CommentId string `json:"comment_id"`
	PostId    string `json:"post_id"`
	Username  string `json:"username"`
	Body      string `json:"body"`
	UpVote    int    `json:"up_vote"`
	DownVote  int    `json:"down_vote"`
}
