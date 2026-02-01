package models

type User struct {
	Id            int    `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	Password_hash string `json:"password_hash,omitempty"`
	Created_at    string `json:"created_at,omitempty"`
}

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
