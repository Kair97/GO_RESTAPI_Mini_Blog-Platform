package repository

import (
	"blog-platform/database"
	"errors"

	"github.com/lib/pq"
)

var ErrAlreadyLiked = errors.New("post already liked")

func LikePost(postId, userId int) error {

	query := `
		insert into likes (user_id, post_id)
		values ($1, $2)
	`

	_, err := database.DB.Exec(query, userId, postId)
	if err != nil {

		if pqerr, ok := err.(*pq.Error); ok {
			if pqerr.Code == "23505" {
				return ErrAlreadyLiked
			}
		}
		return err
	}
	return nil

}

var ErrLikeNotFound = errors.New("like not found")

func UnlikePost(postId, userId int) error {

	query := `
		delete from likes 
		where post_id = $1 and user_id = $2
		`

	res, err := database.DB.Exec(query, postId, userId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrLikeNotFound
	}

	return nil

}

func GetLikesCount(postId int) (int, error) {

	query := `
		select count(*) from likes where post_id = $1
	`

	var count int

	err := database.DB.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil

}

func DidUserLikePost(postId, userId int) (bool, error) {

	query := `
	select exists(
		select 1 from likes where post_id = $1 and user_id = $2
	)
	`
	var liked bool
	err := database.DB.QueryRow(query, postId, userId).Scan(&liked)
	if err != nil {
		return false, err
	}
	return liked, err

}
