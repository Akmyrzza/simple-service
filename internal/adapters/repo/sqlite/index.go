package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type St struct {
	db *sql.DB
}

func New() *St {

	database, _ := sql.Open("sqlite3", "./main.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, name TEXT, phone TEXT)")
	statement.Exec()

	c := &St{
		db: database,
	}

	return c
}
