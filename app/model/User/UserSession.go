package User

type UserSession struct {
	SessionToken string `json:"session_token"`
	Username     string `json:"username"`
}
