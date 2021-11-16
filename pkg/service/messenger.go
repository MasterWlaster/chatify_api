package service

import (
	"chat"
	"chat/pkg/repository"
)

type MessageService struct {
	repository repository.IMessenger
}

func NewMessageService(repository repository.IMessenger) *MessageService {
	return &MessageService{repository: repository}
}

func (s *MessageService) Send(m chat.Message) error {
	return s.repository.Send(m)
}

func (s *MessageService) GetMessage(id int) (chat.Message, error) {
	return s.repository.WaitForMessage(id)
}

func (s *MessageService) GetDialog(id1 int, id2 int) []repository.MessagesOutput {
	return s.repository.GetDialog(id1, id2)
}

func (s *MessageService) GetAllDialogs(id int) []repository.DialogsOutput {
	return s.repository.GetAllDialogs(id)
}

func (s *MessageService) GetUserStatus(id int) bool {
	return s.repository.GetUserStatus(id)
}