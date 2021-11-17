package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable    = "users"
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

//postgres-

type PostgresConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func NewPostgresDB(config PostgresConfig) (*DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
	))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

//-postgres
