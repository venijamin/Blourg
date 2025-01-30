package repository

import (
	"blourg/model/Post"
	"blourg/utils/security"
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
	security.GetMainDB().Where("post_uuid", id).First(&post)
	return post
}

func (repo PostRepository) CreatePost(dto Post.PostCreationDTO) Post.Post {
	post := Post.Post{
		PostUUID: uuid.New().String(),
		UserUUID: dto.UserUUID,
		Title:    dto.Title,
		Body:     dto.Body,
		UpVote:   0,
		DownVote: 0,
	}
	security.GetMainDB().Create(post)

	return post
}
func (repo PostRepository) DeletePostById(id string) {
	var post Post.Post
	security.GetMainDB().Where("post_uuid", id).First(&post).Delete(&post)

}
func (repo PostRepository) UpdatePost(dto Post.PostCreationDTO, id string) Post.Post {
	var post Post.Post
	security.GetMainDB().Where("post_uuid", id).First(&post)
	post.UserUUID = dto.UserUUID
	post.Title = dto.Title
	post.Body = dto.Body
	security.GetMainDB().Save(&post)
	return post
}
