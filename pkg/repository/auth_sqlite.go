package repository

import (
	"chat"
	"fmt"
)

type AuthorizerSqlite struct {
	db *DB
}

func NewAuthorizerSqlite(db *DB) *AuthorizerSqlite {
	return &AuthorizerSqlite{db: db}
}

func (a *AuthorizerSqlite) CreateUser(user chat.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3)",
		usersTable,
	)

	queryId := fmt.Sprintf(
		"SELECT id FROM %s WHERE name=$1 AND username=$2 AND password_hash=$3",
		usersTable,
	)
	_, err := a.db.Exec(query, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	row := a.db.QueryRow(queryId, user.Name, user.Username, user.Password)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthorizerSqlite) GetUser(username string, password string) (chat.User, error) {
	var user chat.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := a.db.Get(&user, query, username, password)

	return user, err
}

func (a *AuthorizerSqlite) CanCreateNewUser(username string) bool {
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)
	var id int
	err := a.db.Get(&id, query, username)
	if err != nil {
		return true
	}
	return false
}