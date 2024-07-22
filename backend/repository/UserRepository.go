package repository

import (
	"backend/model"
	"backend/security"
	"errors"
)

type UserRepository struct {
}

func (repo UserRepository) GetAllUsers() []model.User {
	var users []model.User
	security.GetDatabase().Find(&users)

	return users
}

func (repo UserRepository) GetUserByUsername(username string) model.User {
	var user model.User
	security.GetDatabase().Where("name = ?", username).Find(&user)
	return user
}

func (repo UserRepository) RegisterUser(email string, username string, password string) model.User {
	var newUser model.User
	newUser.Email = email
	newUser.Username = username
	newUser.Password = security.HashPassword(password)
	security.GetDatabase().Create(&newUser)
	return newUser
}

func (repo UserRepository) LoginUser(email string, username string, password string) (bool, error) {
	user := repo.GetUserByUsername(username)
	if security.VerifyPassword(password, user.Password) {
		return true, nil
	}
	return false, errors.New("username or password incorrect")
}

func (repo UserRepository) LogoutUser(email string, username string) {
	return
}

func (repo UserRepository) ChangePassword(email string, username string, password string) {
	return
}

func (repo UserRepository) DeleteUser(username string) {
	security.GetDatabase().Delete("name = ?", username)
}
