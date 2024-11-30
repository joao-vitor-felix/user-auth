package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	sqlFile, err := os.ReadFile("sql/users.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return nil, err
	}

	return db, nil
}
