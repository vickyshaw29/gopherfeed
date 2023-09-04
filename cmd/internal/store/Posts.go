package store

import (
	"context"
	"database/sql"
)

type Post struct {
	ID        int64
	Content   string
	Title     string
	UserID    int64
	CreatedAt string
	UpdatedAt string
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context) error {
	return nil
}
