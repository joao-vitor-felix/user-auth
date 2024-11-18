package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/goschool/crud/types"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
}

type SQLiteUserStore struct {
	db *sql.DB
}

func NewSQLiteUserStore(db *sql.DB) *SQLiteUserStore {
	return &SQLiteUserStore{
		db: db,
	}
}

func (u *SQLiteUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	query := `INSERT INTO users (email, password_hash)
	VALUES (?, ?)
	RETURNING id`

	var userID string
	err := u.db.QueryRowContext(ctx, query, user.Email, user.PasswordHash).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("CreateUser :%w", err)
	}

	user.ID = userID
	return user, nil
}
