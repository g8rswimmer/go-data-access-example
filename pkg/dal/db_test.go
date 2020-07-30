package dal

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func setupDB(stmts []string) *sql.DB {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory")
	if err != nil {
		panic(err)
	}

	for _, stmt := range stmts {
		if _, err := db.Exec(stmt); err != nil {
			db.Close()
			panic(err)
		}
	}
	return db

}
