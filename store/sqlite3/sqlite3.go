package sqlite3

import (
	"database/sql"
	"wargame-bot/store"
)

// An sqlite 3 database
type Sqlite3Store struct {
	db *sql.DB
}

// Ceates a new sqlite 3 database
func NewSqlite3Store(db *sql.DB) store.Store {
	return &Sqlite3Store{
		db: db,
	}
}
