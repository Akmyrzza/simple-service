package sqlite

import (
	"database/sql"

	"github.com/Akmyrzza/simple-service/internal/adapters/logger"
	_ "github.com/mattn/go-sqlite3"
)

type St struct {
	db *sql.DB
	lg logger.Lite
}

func New(lg logger.Lite) *St {

	database, _ := sql.Open("sqlite3", "./main.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, name TEXT, phone TEXT)")
	statement.Exec()

	return &St{
		db: database,
		lg: lg,
	}
}
