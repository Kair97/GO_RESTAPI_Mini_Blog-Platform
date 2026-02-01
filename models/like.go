package models

type Like struct {
	User_id    int    `json:"user_id,omitempty"`
	Post_id    int    `json:"post_id,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}
