package service

import (
	"chat"
	"chat/pkg/repository"
)

type IAuthorizer interface {
	CreateUser(user chat.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type IMessenger interface {
	Send(m chat.Message) error
	GetMessage(id int) (chat.Message, error)
	GetDialog(id1 int, id2 int) []repository.MessagesOutput
	GetAllDialogs(id int) []repository.DialogsOutput
	GetUserStatus(id int) bool
}

type Service struct {
	IAuthorizer
	IMessenger
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		NewAuthService(repository.IAuthorizer),
		NewMessageService(repository.IMessenger),
	}
}

