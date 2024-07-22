package repository

import (
	"backend/model"
	"backend/security"
	"errors"
)

// TODO: change all the params to dto objects

type UserRepository struct {
}

func (repo UserRepository) GetAllUsers() []model.User {
	var users []model.User
	security.GetMainDB().Find(&users)

	return users
}

func (repo UserRepository) GetUserByUsername(username string) model.User {
	var user model.User
	security.GetMainDB().Where("username = ?", username).Find(&user)
	return user
}

func (repo UserRepository) RegisterUser(email string, username string, password string) model.User {
	var newUser model.User
	newUser.Email = email
	newUser.Username = username
	newUser.Password = security.HashPassword(password)
	security.GetMainDB().Create(&newUser)
	return newUser
}

// LoginUser TODO: Support login with email
func (repo UserRepository) LoginUser(email string, username string, password string) (string, error) {
	user := repo.GetUserByUsername(username)

	// If the user is verified meaning the password matches the one that
	// is hashed and stored in the database then create a session token
	if security.VerifyPassword(password, user.Password) {

		sessionToken, hashedSessionToken := security.GenerateSessionToken()
		security.GetUserSessionsDB().Create(&model.UserSession{
			Username:     username,
			SessionToken: hashedSessionToken,
		})

		return sessionToken, nil
	}

	return "", errors.New("username or password incorrect")

}

// LogoutUser TODO: logout functionality
func (repo UserRepository) LogoutUser(email string, username string) {
	return
}

// ChangePassword TODO: change password functionality
func (repo UserRepository) ChangePassword(email string, username string, password string) {
	return
}

// DeleteUser checks if the user sessionToken and the passwords are correct to validate the delete process
// TODO: fix it doesn't work well
func (repo UserRepository) DeleteUser(userDelete model.UserDeleteDTO) bool {
	username := userDelete.Username
	password := userDelete.Password
	sessionToken := userDelete.SessionToken

	user := repo.GetUserByUsername(username)
	var userSessions []model.UserSession
	security.GetUserSessionsDB().Where("username = ?", username).Find(&userSessions)

	for _, userSession := range userSessions {
		if security.VerifyPassword(password, user.Password) && security.VerifySession(sessionToken, userSession.SessionToken) {
			security.GetMainDB().Delete(&user)
			security.GetUserSessionsDB().Delete(&userSessions)
			return true
		}
	}

	return false
}
