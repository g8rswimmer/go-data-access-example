package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // placing the database import within the package that it initialized
)

// Open will open a database and execute a series of statments
func Open(ctx context.Context, stmts []string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory")
	if err != nil {
		return nil, fmt.Errorf("sqlite database open error %w", err)
	}

	for _, stmt := range stmts {
		if _, err := db.ExecContext(ctx, stmt); err != nil {
			db.Close()
			return nil, fmt.Errorf("sqlite database statment (%s) error %w", stmt, err)
		}
	}
	return db, nil
}
