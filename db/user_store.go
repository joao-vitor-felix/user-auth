package db

import (
	"context"
	"database/sql"

	"github.com/goschool/crud/types"
)

type UserStore interface {
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
	return nil, nil
}
