package User

type UserDisplayDTO struct {
	Country     string `json:"country"`
	Created     string `json:"created"`
	Birth       string `json:"birth"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
}
