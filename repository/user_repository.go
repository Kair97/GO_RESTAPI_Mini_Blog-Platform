package repository

import (
	"blog-platform/database"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

func CreateUser(username, email, passwordHash string) error {

	query := `
		insert into users 
		(username, email, password_hash)
		values($1, $2, $3)
	`
	_, err := database.DB.Exec(query, username, email, passwordHash)
	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return ErrEmailAlreadyExists
			}
		}

		return err
	}
	return nil

}

var ErrInvalidCredentials = errors.New("Invalid credential")

func GetUserByEmail(email string) (int, string, error) {
	query := `
		select id, password_hash from users where email = $1
	`
	var id int
	var hash string
	err := database.DB.QueryRow(query, email).Scan(&id, &hash)
	if err == sql.ErrNoRows {
		return 0, "", ErrInvalidCredentials
	}

	if err != nil {
		return 0, "", err
	}

	return id, hash, nil

}
