package repository

import (
	"backend/model/Comment"
	"backend/security"
	"github.com/google/uuid"
)

type CommentRepository struct{}

func (repo CommentRepository) CreateComment(dto Comment.CommentCreationDTO) Comment.Comment {
	comment := Comment.Comment{
		CommentId: uuid.New().String(),
		PostId:    dto.PostId,
		Username:  dto.Username,
		Body:      dto.Body,
		UpVote:    0,
		DownVote:  0,
	}
	security.GetMainDB().Create(&comment)

	return comment
}

// DeleteComment TODO: add functionality to delete comments
func (repo CommentRepository) DeleteComment(commentId string) {
	return
}
