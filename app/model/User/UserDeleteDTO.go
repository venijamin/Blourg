package User

type UserDeleteDTO struct {
	Username     string `json:"username"`
	SessionToken string `json:"session_token"`
	Password     string `json:"password"`
}
