package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable = "users"
	messagesTable = "messages"
)

type DB struct {
	*sqlx.DB
}

//sqlite-

type SqliteConfig struct {
	Path string
}

func NewSqliteDB(config SqliteConfig) (*DB, error) {
	db, err := sqlx.Open("sqlite3", config.Path)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

//-sqlite