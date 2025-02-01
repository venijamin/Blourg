package repository

import (
	"blourg/model/User"
	"blourg/utils/security"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct{}

func (repo UserRepository) CreateUser(creds User.UserRegisterDTO) error {
	user := User.User{
		Password:    creds.Password,
		Username:    creds.Username,
		Email:       creds.Email,
		Country:     "",
		Created:     time.DateTime,
		Birth:       "",
		DisplayName: creds.Username,
		Id:          uuid.NewString(),
	}

	if err := security.GetMainDB().Create(user).Error; err != nil {
		return gorm.ErrDuplicatedKey
	}
	return nil
}

func (repo UserRepository) GetUserByUsername(username string) (*User.User, error) {
	user := User.User{}
	if err := security.GetMainDB().Model(&user).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) GetUserByEmail(email string) (*User.User, error) {
	user := User.User{}
	if err := security.GetMainDB().Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) GetUserByDisplayname(displayname string) (*User.User, error) {
	user := User.User{}
	if err := security.GetMainDB().Model(&user).Where("displayname = ?", displayname).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) GetPasswordByUsername(username string) (string, error) {
	var user User.User
	if err := security.GetMainDB().Model(&user).Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}

func (repo UserRepository) GetUserByUUID(uuid string) (*User.User, error) {
	user := User.User{}
	if err := security.GetMainDB().Model(&user).Where("id = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
