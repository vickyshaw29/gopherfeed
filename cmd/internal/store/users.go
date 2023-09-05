package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (username, email, password)
	VALUES($1, $2, $3) RETURNING id, created_at
	`
	err := s.db.QueryRowContext(ctx,
		query, user.Username, user.Email, user.Password,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
