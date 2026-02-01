package repository

import (
	"blog-platform/database"
	"blog-platform/models"
	"errors"
)

func GetCommentByPostId(postId int) ([]models.Comment, error) {

	query := `
		select id, content, post_id, user_id, created_at
		from comments 
		where post_id = $1
		order by created_at asc
	`

	rows, err := database.DB.Query(query, postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var com models.Comment
		err := rows.Scan(
			&com.Id,
			&com.Content,
			&com.Post_id,
			&com.User_id,
			&com.Created_at,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, com)
	}
	return comments, nil

}

func CreateComment(postId, userId int, content string) (models.Comment, error) {

	query := `
		insert into comments (content, post_id, user_id)
		values ($1, $2, $3)
		returning id, content, post_id, user_id, created_at
	`

	var com models.Comment
	err := database.DB.QueryRow(query, content, postId, userId).Scan(
		&com.Id,
		&com.Content,
		&com.Post_id,
		&com.User_id,
		&com.Created_at,
	)
	if err != nil {
		return models.Comment{}, err
	}

	return com, nil
}

var ErrNotFound = errors.New("Comment not found")

func DeleteComment(commentId, userId int) error {

	query := `
		delete from comments 
		where id = $1 and user_id = $2
	`

	result, err := database.DB.Exec(query, commentId, userId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil

}
