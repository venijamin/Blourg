package service

import (
	"blourg/model/Post"
	"blourg/model/User"
	"blourg/repository"
	"blourg/utils/security"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var postRepository repository.PostRepository

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var postListTemplate = template.Must(template.ParseFiles("src/template/post-list.html"))
	posts := postRepository.GetAllPosts()

	w.Header().Set("Content-Type", "text/html")

	if err := postListTemplate.Execute(w, posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	var postTemplate = template.Must(template.ParseFiles("src/template/post.html"))
	vars := mux.Vars(r)
	postId, _ := vars["postId"]
	_ = json.NewDecoder(r.Body).Decode(&postId)

	w.Header().Set("Content-Type", "text/html")

	post := postRepository.GetPostById(postId)
	if err := postTemplate.Execute(w, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

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
	user, _ := userRepository.GetUserByUsername(claims.Username)

	// Create the PostCreationDTO from the form data
	postCreation := Post.PostCreationDTO{
		UserUUID: user.Id,
		Title:    r.FormValue("title"),
		Body:     r.FormValue("body"),
	}

	// Call your repository function to create the post
	postRepository.CreatePost(postCreation)

	// Optionally send a response back to HTMX
	// Returning an HTML response (success message or the created post)
	w.Header().Set("Content-Type", "text/html") // Set the correct content type for HTML response
	w.Header().Set("HX-Refresh", "true")

	w.WriteHeader(http.StatusCreated)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	postId, _ := vars["postId"]
	_ = json.NewDecoder(r.Body).Decode(&postId)

	// Create the PostCreationDTO from the form data
	postCreation := Post.PostCreationDTO{
		UserUUID: r.FormValue("user_uuid"),
		Title:    r.FormValue("title"),
		Body:     r.FormValue("body"),
	}

	// Call your repository function to create the post
	postRepository.UpdatePost(postCreation, postId)

	// Optionally send a response back to HTMX
	// Returning an HTML response (success message or the created post)
	w.Header().Set("Content-Type", "text/html") // Set the correct content type for HTML response
	w.Header().Set("HX-Refresh", "true")

	w.WriteHeader(http.StatusNoContent)
}

func DeletePostById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	postId, _ := vars["postId"]
	_ = json.NewDecoder(r.Body).Decode(&postId)

	w.Header().Set("HX-Refresh", "true")

	postRepository.DeletePostById(postId)
}
