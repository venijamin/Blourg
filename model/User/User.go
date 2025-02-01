package User

type User struct {
	Password    string `json:"password"`
	Username    string `json:"username" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Country     string `json:"country"`
	Created     string `json:"created"`
	Birth       string `json:"birth"`
	DisplayName string `json:"display_name"`
	Id          string `json:"id" gorm:"primaryKey"`
}
