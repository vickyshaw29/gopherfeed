package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrorNotFound = errors.New("record not found")
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
		GetById(ctx context.Context, postID int) (*Post, error)
	}
	Users interface {
		Create(ctx context.Context, user *User) error
	}
	Comments interface {
		GetByPostID(ctx context.Context, postID int64) (*[]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
