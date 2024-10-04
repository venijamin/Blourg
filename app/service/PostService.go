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

var postTemplate = template.Must(template.ParseFiles("template/post-list.html"))

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts := postRepository.GetAllPosts() // Fetch the posts from your repository

	w.Header().Set("Content-Type", "text/html")  // Set content type to text/html
	w.Header().Set("HX-Trigger", "postsUpdated") // Optional: Trigger an HTMX event

	if err := postTemplate.Execute(w, posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, _ := vars["postId"]
	_ = json.NewDecoder(r.Body).Decode(&postId)

	json.NewEncoder(w).Encode(postRepository.GetPostById(postId))
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var postCreation Post.PostCreationDTO
	_ = json.NewDecoder(r.Body).Decode(&postCreation)

	postRepository.CreatePost(postCreation)
}

// DeletePost TODO: make it
func DeletePost(w http.ResponseWriter, r *http.Request) {

	var postId string
	_ = json.NewDecoder(r.Body).Decode(&postId)

}

func GetAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, _ := vars["postId"]
	_ = json.NewDecoder(r.Body).Decode(&postId)

	json.NewEncoder(w).Encode(postRepository.GetAllCommentsForPost(postId))
}
