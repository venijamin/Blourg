package repository

import (
	"backend/model/Comment"
	"backend/model/Post"
	"backend/security"
	"github.com/google/uuid"
)

type PostRepository struct{}

func (repo PostRepository) GetAllPosts() []Post.Post {
	var posts []Post.Post
	security.GetMainDB().Find(&posts)
	return posts
}

func (repo PostRepository) GetPostById(id string) Post.Post {
	var post Post.Post
	security.GetMainDB().Where("post_id", id).First(&post)
	return post
}

func (repo PostRepository) CreatePost(dto Post.PostCreationDTO) Post.Post {
	post := Post.Post{
		PostId:   uuid.New().String(),
		Username: dto.Username,
		Title:    dto.Title,
		Body:     dto.Body,
		UpVote:   0,
		DownVote: 0,
	}
	security.GetMainDB().Create(&post)

	return post
}

func (repo CommentRepository) GetAllCommentsForPost(postId string) []Comment.Comment {
	var comments []Comment.Comment
	security.GetMainDB().Where("post_id = ?", postId).Find(&comments)

	return comments
}
