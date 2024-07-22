package service

import (
	"backend/model"
	"backend/repository"
	"encoding/json"
	"net/http"
)

var userRepository repository.UserRepository

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(userRepository.GetAllUsers())
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	// Add new user to database
	userRepository.RegisterUser(newUser.Email, newUser.Username, newUser.Password)
	json.NewEncoder(w).Encode(newUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin model.UserLoginDTO
	_ = json.NewDecoder(r.Body).Decode(&userLogin)

	// Create a session token to keep the user logged in

	userRepository.LoginUser(userLogin.Email, userLogin.Username, userLogin.Password)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var username string
	_ = json.NewDecoder(r.Body).Decode(&username)
	userRepository.DeleteUser(username)
}
