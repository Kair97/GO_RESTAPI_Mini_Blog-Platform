package models

type Comment struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	Post_id    int    `json:"post_id"`
	User_id    int    `json:"user_id"`
	Created_at string `json:"created_at"`
}

type CreateComment struct {
	Content string `json:"content" binding:"required,min=1"`
}
