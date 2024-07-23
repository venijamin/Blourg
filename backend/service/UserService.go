package service

import (
	"backend/model/User"
	"backend/repository"
	"encoding/json"
	"log"
	"net/http"
)

var userRepository repository.UserRepository

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(userRepository.GetAllUsers())
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser User.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	// Add new user to database
	userRepository.RegisterUser(newUser.Email, newUser.Username, newUser.Password)
	json.NewEncoder(w).Encode(newUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin User.UserLoginDTO
	_ = json.NewDecoder(r.Body).Decode(&userLogin)

	sessionToken, err := userRepository.LoginUser(userLogin.Email, userLogin.Username, userLogin.Password)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(sessionToken)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userDelete User.UserDeleteDTO
	_ = json.NewDecoder(r.Body).Decode(&userDelete)
	json.NewEncoder(w).Encode(userRepository.DeleteUser(userDelete))
}
