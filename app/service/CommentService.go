package service

import (
	"backend/model/Comment"
	"backend/repository"
	"encoding/json"
	"net/http"
)

var commentRepository repository.CommentRepository

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var commentCreation Comment.CommentCreationDTO
	_ = json.NewDecoder(r.Body).Decode(&commentCreation)

	commentRepository.CreateComment(commentCreation)

}
