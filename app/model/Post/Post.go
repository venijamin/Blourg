package Post

type Post struct {
	PostId   string `json:"post_id"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	UpVote   int    `json:"up_vote"`
	DownVote int    `json:"down_vote"`
}
