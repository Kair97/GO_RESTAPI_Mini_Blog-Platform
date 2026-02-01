package repository

import (
	"blog-platform/database"
	"blog-platform/models"
	"database/sql"
	"errors"
)

func GetPostsWithLikes(userId *int) ([]models.PostListResponse, error) {
	var uid sql.NullInt64
	if userId != nil {
		uid = sql.NullInt64{Int64: int64(*userId), Valid: true}
	} else {
		uid = sql.NullInt64{Valid: false}
	}
	query := `
		SELECT
			p.id,
			p.title,
			p.content,
			p.author_id,
			p.created_at,
			p.updated_at,
			COUNT(l.user_id) AS likes_count,
			BOOL_OR(ul.user_id IS NOT NULL) AS liked_by_me
		FROM posts p
		LEFT JOIN likes l
			ON l.post_id = p.id
		LEFT JOIN likes ul
			ON ul.post_id = p.id AND ul.user_id = $1
		GROUP BY p.id
		ORDER BY p.created_at DESC
	`

	rows, err := database.DB.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostListResponse

	for rows.Next() {

		var p models.PostListResponse

		if err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Content,
			&p.Author_id,
			&p.Created_at,
			&p.Updated_at,
			&p.LikesCount,
			&p.LikedByMe,
		); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil

}

func GetPosts() ([]models.Post, error) {
	query := `
		Select id, title, content, author_id, created_at, updated_at
		from posts 
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post

	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Author_id,
			&post.Created_at,
			&post.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

var Notfound = errors.New("Post not found")

func GetPostById(id int) (models.Post, error) {

	query := `
		select id, title, content, author_id, created_at, updated_at 
		from posts
		where id = $1
	`

	var post models.Post

	err := database.DB.QueryRow(query, id).Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.Author_id,
		&post.Created_at,
		&post.Updated_at,
	)
	if err == sql.ErrNoRows {
		return models.Post{}, Notfound
	}

	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func CreatePost(title string, content string, author_id int) (models.Post, error) {

	query := `
		insert into posts (title, content, author_id)
		values($1, $2, $3)
		returning id, title, content, author_id, created_at, updated_at 
	`

	var post models.Post

	err := database.DB.QueryRow(
		query,
		title,
		content, author_id,
	).Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.Author_id,
		&post.Created_at,
		&post.Updated_at,
	)

	if err != nil {
		return models.Post{}, err
	}
	return post, nil

}

func UpdatePost(id int, title string, content string, author_id int) (models.Post, error) {

	query := `
	update posts 
	set title = $1, content = $2, updated_at = NOW()
	where id = $3 and author_id = $4 
	RETURNING id, title, content, author_id, created_at, updated_at
	`

	var post models.Post
	err := database.DB.QueryRow(
		query,
		title,
		content,
		id,
		author_id,
	).Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.Author_id,
		&post.Created_at,
		&post.Updated_at,
	)
	if err == sql.ErrNoRows {
		return models.Post{}, Notfound
	}
	if err != nil {
		return models.Post{}, err
	}

	return post, nil

}

var ErrNothingToUpdate = errors.New("Nothing to update")

func PatchPost(id int, title *string, content *string, author_id int) (models.Post, error) {

	query := `
		update posts 
		set title = coalesce($1, title), content = coalesce($2, content), updated_at = NOW()
		where id = $3 and author_id = $4
		returning id, title, content, author_id, created_at, updated_at
	`

	var post models.Post

	err := database.DB.QueryRow(
		query,
		title,
		content,
		id,
		author_id,
	).Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.Author_id,
		&post.Created_at,
		&post.Updated_at,
	)
	if err == sql.ErrNoRows {
		return models.Post{}, Notfound
	}
	if err != nil {
		return models.Post{}, err
	}

	return post, nil

}

func DeletePost(id int, author_id int) error {
	query := `
		delete from posts 
		where id = $1 and author_id = $2
	`

	result, err := database.DB.Exec(query, id, author_id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return Notfound
	}

	return nil

}
