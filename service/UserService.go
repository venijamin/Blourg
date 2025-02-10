package service

import (
	"blourg/model/User"
	"blourg/repository"
	"blourg/utils/security"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var userRepository repository.UserRepository

func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code until this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &User.UserJWT{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return security.GetJWTKey(), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(security.GetJWTKey())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
func Welcome(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &User.UserJWT{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return security.GetJWTKey(), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
func Signout(w http.ResponseWriter, r *http.Request) {
	// immediately clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}
func Signup(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	creds := User.UserRegisterDTO{
		Password: password,
		Username: username,
		Email:    email,
	}
	// Parse and decode the request body into a new `Credentials` instance
	//creds := &User.UserRegisterDTO{}
	//err := json.NewDecoder(r.Body).Decode(creds)
	//if err != nil {
	//	// If there is something wrong with the request body, return a 400 status
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// Next, insert the username, along with the hashed password into the database
	creds.Password = string(hashedPassword)

	if err := userRepository.CreateUser(creds); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	creds := User.UserLoginDTO{
		Password: password,
		Username: username,
	}
	// Get the JSON body and decode into credentials
	//err := json.NewDecoder(r.Body).Decode(&creds)
	//if err != nil {
	//	// If the structure of the body is wrong, return an HTTP error
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	// Get the expected password from our in memory map
	hashedPassword, err := userRepository.GetPasswordByUsername(creds.Username)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(30 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &User.UserJWT{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(security.GetJWTKey())
	if err != nil {
		log.Println(err)
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func ServeSignup(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")
	if err == nil {
		http.RedirectHandler("/", 302).ServeHTTP(w, r)
		return
	}
	t, _ := template.ParseFiles("src/sign-up.html")
	t.Execute(w, nil)
}

func ServeSignin(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")
	if err == nil {
		http.RedirectHandler("/", 302).ServeHTTP(w, r)
		return
	}
	t, _ := template.ParseFiles("src/sign-in.html")
	t.Execute(w, nil)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &User.UserJWT{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return security.GetJWTKey(), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create the PostCreationDTO from the form data
	w.Header().Set("Content-Type", "text/html")

	var userTemplate = template.Must(template.ParseFiles("src/template/user-profile.html"))
	user, _ := userRepository.GetUserByUsername(claims.Username)
	userDisplayDTO := User.UserDisplayDTO{
		Country:     user.Country,
		Created:     user.Created,
		Birth:       user.Birth,
		DisplayName: user.DisplayName,
		Username:    user.Username,
	}
	if err := userTemplate.Execute(w, userDisplayDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ServeAboutMe(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/about-me.html")
	t.Execute(w, nil)
}
