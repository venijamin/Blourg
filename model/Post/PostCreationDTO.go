package Post

type PostCreationDTO struct {
	UserUUID     string `json:"user_uuid"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	CreationDate string `json:"creation_date"`
}
