package service

import (
	"backend/model/Post"
	"backend/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var postRepository repository.PostRepository

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(postRepository.GetAllPosts())
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
