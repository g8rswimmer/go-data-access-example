package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // placing the database import within the package that it initialized
)

func Open(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory")
	if err != nil {
		return nil, fmt.Errorf("sqlite database open error %w", err)
	}

	if _, err := db.ExecContext(ctx, user); err != nil {
		db.Close()
		return nil, fmt.Errorf("sqlite database user table error %w", err)
	}
	return db, nil
}
