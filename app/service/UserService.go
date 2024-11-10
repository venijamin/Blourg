package service

import (
	"backend/model/User"
	"backend/repository"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
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

// CheckSessionToken checks if a session token exists and is valid
func CheckSessionToken(r *http.Request) bool {
	// Get the session token from cookies
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		return false // No session token found
	}
	userCookie, err := r.Cookie("username")
	if err != nil || userCookie.Value == "" {
		return false // No session token found
	}
	//security.GetUserSessionsDB().Where("username", userCookie).First(&)

	// You can add additional validation for the token here, like checking expiration
	// For now, we just check if it's set
	return true
}

func RenderLoginLinks(w http.ResponseWriter, r *http.Request) {
	// Check if user is logged in by looking for the session token
	if CheckSessionToken(r) {
		// Render the "Log Out" link if the user is logged in
		w.Write([]byte(`
			<a href="#" hx-get="/logout" hx-trigger="click" hx-target="#modal-content">Log Out</a>
		`))
	} else {
		// Render the "Log In" and "Register" links if the user is not logged in
		w.Write([]byte(`
			<a href="#" hx-get="/login" hx-target="#modal-content" hx-trigger="click">Log In</a> |
			<a href="#" hx-get="/register" hx-target="#modal-content" hx-trigger="click">Register</a>
		`))
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
	sessionCookie := &http.Cookie{
		Name:     "session_token",                // Name of the cookie
		Value:    sessionToken,                   // The value of the session token
		Expires:  time.Now().Add(24 * time.Hour), // Set expiration time (e.g., 24 hours)
		HttpOnly: true,                           // Prevent JavaScript access (security feature)
		Secure:   false,                          // Ensure the cookie is sent over HTTPS
		SameSite: http.SameSiteStrictMode,        // Prevent CSRF attacks (Strict mode)
		Path:     "/",                            // Cookie is available across the entire site
	}
	usernameCookie := &http.Cookie{
		Name:     "username_token",               // Name of the cookie
		Value:    userLogin.Username,             // The value of the session token
		Expires:  time.Now().Add(24 * time.Hour), // Set expiration time (e.g., 24 hours)
		HttpOnly: true,                           // Prevent JavaScript access (security feature)
		Secure:   false,                          // Ensure the cookie is sent over HTTPS
		SameSite: http.SameSiteStrictMode,        // Prevent CSRF attacks (Strict mode)
		Path:     "/",                            // Cookie is available across the entire site
	}
	// Set the cookie in the response
	http.SetCookie(w, sessionCookie)
	http.SetCookie(w, usernameCookie)

}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userDelete User.UserDeleteDTO
	_ = json.NewDecoder(r.Body).Decode(&userDelete)
	json.NewEncoder(w).Encode(userRepository.DeleteUser(userDelete))
}
