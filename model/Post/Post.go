package Post

type Post struct {
	PostUUID     string `json:"post_uuid" gorm:"primaryKey"`
	UserUUID     string `json:"user_uuid"`
	Title        string `json:"title" gorm:"unique"`
	Body         string `json:"body"`
	UpVote       int    `json:"up_vote"`
	DownVote     int    `json:"down_vote"`
	CreationDate string `json:"creation_date"`
}
