package models

import "time"

type Post struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author_id  int       `json:"author_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CreatePost struct {
	Title   string `json:"title" binding:"required,min=3"`
	Content string `json:"content" binding:"required,min=10"`
}

type UpdatePost struct {
	Title   string `json:"title,omitempty" binding:"required,min=3"`
	Content string `json:"content,omitempty" binding:"required,min=10"`
}

type PatchPost struct {
	Title   *string `json:"title,omitempty" binding:"omitempty,min=3"`
	Content *string `json:"content,omitempty" binding:"omitempty,min=10"`
}

type PostResponse struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author_id  int       `json:"author_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`

	LikesCount int  `json:"likes_count"`
	LikedByMe  bool `json:"liked_by_me"`
}

type PostListResponse struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author_id  int       `json:"author_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`

	LikesCount int  `json:"likes_count"`
	LikedByMe  bool `json:"liked_by_me"`
}
