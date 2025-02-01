package User

type UserRegisterDTO struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
