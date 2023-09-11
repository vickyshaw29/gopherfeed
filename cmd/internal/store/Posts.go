package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"userId"`
	Tags      []string  `json:"tags"`
	Comments  []Comment `json:"comments"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRowContext(ctx,
		query, post.Content, post.Title, post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetById(ctx context.Context, postID int) (*Post, error) {
	query := `SELECT id, title, content, user_id, tags, created_at, updated_at
	FROM posts
	WHERE id = $1
	`
	row := s.db.QueryRowContext(ctx, query, postID)
	var post Post
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorNotFound
		default:
			return nil, err
		}
	}
	return &post, nil
}

func (s *PostStore) DeleteById(ctx context.Context, postID int) error {
	query := `DELETE FROM posts WHERE id = $1`
	result, err := s.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrorNotFound
	}
	return nil

}

func (s *PostStore) UpdateById(ctx context.Context, post *Post) error {
	query := `UPDATE posts
	SET title = $1, content=$2
	WHERE id = $3
	`
	_, err := s.db.ExecContext(ctx, query, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}
	return nil
}
