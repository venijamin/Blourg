package model

type UserDeleteDTO struct {
	Username     string `json:"username"`
	SessionToken string `json:"sessionToken"`
	Password     string `json:"password"`
}
