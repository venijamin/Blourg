package model

type UserSession struct {
	SessionToken string `json:"sessionToken"`
	Username     string `json:"username"`
}
