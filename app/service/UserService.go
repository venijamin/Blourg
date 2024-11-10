package service

import (
	"backend/model/User"
	"backend/repository"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var userRepository repository.UserRepository

var userTemplate = template.Must(template.ParseFiles("src/template/user-list.html"))

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := userRepository.GetAllUsers()

	w.Header().Set("Content-Type", "text/html")  // Set content type to text/html
	w.Header().Set("HX-Trigger", "postsUpdated") // Optional: Trigger an HTMX event

	if err := userTemplate.Execute(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Create the PostCreationDTO from the form data
	newUser := User.UserLoginDTO{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	// Add new user to database
	userRepository.RegisterUser(newUser.Email, newUser.Username, newUser.Password)
}
func GetRegisterPage(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("src/template/register-form.html"))
	w.Header().Set("Content-Type", "text/html") // Set content type to text/html

	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("src/template/login-form.html"))
	w.Header().Set("Content-Type", "text/html") // Set content type to text/html

	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Create the PostCreationDTO from the form data
	userLogin := User.UserLoginDTO{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

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
