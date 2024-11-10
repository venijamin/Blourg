package service

import (
	"backend/model/Post"
	"backend/repository"
	"encoding/json"
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
	var postTemplate = template.Must(template.ParseFiles("src/template/post-expanded.html"))
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

	// Create the PostCreationDTO from the form data
	postCreation := Post.PostCreationDTO{
		Username: r.FormValue("username"),
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
		Username: r.FormValue("username"),
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

func GetAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, _ := vars["postId"]
	_ = json.NewDecoder(r.Body).Decode(&postId)

	json.NewEncoder(w).Encode(postRepository.GetAllCommentsForPost(postId))
}
