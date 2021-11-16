package repository

import "chat"

type IAuthorizer interface {
	CreateUser(user chat.User) (int, error)
	GetUser(username string, password string) (chat.User, error)
	CanCreateNewUser(username string) bool
}

type IMessenger interface {
	Send(m chat.Message) error
	WaitForMessage(id int) (chat.Message, error)
	GetDialog(id1 int, id2 int) []MessagesOutput
	GetAllDialogs(id int) []DialogsOutput
	GetUserStatus(id int) bool
}

type Repository struct {
	IAuthorizer
	IMessenger
}

func NewRepository(db *DB) *Repository {
	return &Repository{
		NewAuthorizerSqlite(db),
		NewMessengerSqlite(db),
	}
}
