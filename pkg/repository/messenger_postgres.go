package repository

import (
	"chat"
	"fmt"
)

type MessengerPostgres struct {
	db *DB
}

func NewMessengerPostgres(db *DB) *MessengerPostgres {
	return &MessengerPostgres{db: db}
}

func (s *MessengerPostgres) Send(m chat.Message) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (sender_id, receiver_id, time_sent, text) VALUES ($1, $2, $3, $4)",
		messagesTable,
	)

	_, err := s.db.Exec(query, m.SenderId, m.ReceiverId, m.TimeSent, m.Text)
	if err != nil {
		return err
	}

	notifyReceiver(m)

	return nil
}

func (s *MessengerPostgres) WaitForMessage(id int) (chat.Message, error) {
	return justReceivedMessage(id)
}

func (s *MessengerPostgres) GetDialog(id1 int, id2 int) []MessagesOutput {
	query := fmt.Sprintf(
		"SELECT %[1]s.username, %[2]s.time_sent, %[2]s.text FROM %[1]s INNER JOIN %[2]s ON %[2]s.sender_id = %[1]s.id WHERE %[2]s.sender_id=$1 AND %[2]s.receiver_id=$2 OR %[2]s.sender_id=$2 AND %[2]s.receiver_id=$1",
		usersTable,
		messagesTable,
	)

	rows, err := s.db.Query(query, id1, id2)
	if err != nil {
		//todo
	}

	messages := make([]MessagesOutput, 0)

	for rows.Next() {
		var m MessagesOutput

		err = rows.Scan(&m.SenderUsername, &m.TimeSent, &m.Text)
		if err != nil {
			//todo
		}

		messages = append(messages, m)
	}

	return messages
}

func (s *MessengerPostgres) GetAllDialogs(id int) []DialogsOutput {
	query := fmt.Sprintf(
		"SELECT id, username, status_online FROM %s WHERE id!=$1",
		usersTable,
	)

	rows, err := s.db.Query(query, id)
	if err != nil {
		//todo
	}

	dialogs := make([]DialogsOutput, 0)

	for rows.Next() {
		var d DialogsOutput

		err = rows.Scan(&d.Id, &d.Name, &d.StatusOnline)
		if err != nil {
			//todo
		}

		dialogs = append(dialogs, d)
	}

	return dialogs
}

func (s *MessengerPostgres) GetUserStatus(id int) bool {
	query := fmt.Sprintf(
		"SELECT status_online FROM %s WHERE id=$1",
		usersTable,
	)

	var status bool

	row := s.db.QueryRow(query, id)
	if err := row.Scan(&status); err != nil {
		//todo
	}

	return status
}