package repository

import (
	"chat"
	"fmt"
)

type AuthorizerPostgres struct {
	db *DB
}

func NewAuthorizerPostgres(db *DB) *AuthorizerPostgres {
	return &AuthorizerPostgres{db: db}
}

func (a *AuthorizerPostgres) CreateUser(user chat.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id",
		usersTable,
	)

	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthorizerPostgres) GetUser(username string, password string) (chat.User, error) {
	var user chat.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := a.db.Get(&user, query, username, password)

	return user, err
}

func (a *AuthorizerPostgres) CanCreateNewUser(username string) bool {
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)
	var id int
	err := a.db.Get(&id, query, username)
	if err != nil {
		return true
	}
	return false
}