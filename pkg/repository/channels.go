package repository

import (
	"chat"
)

var channels map[int]chan chat.Message

func notifyReceiver(m chat.Message) {
	if _, ok := channels[m.ReceiverId]; !ok {
		return
	}

	channels[m.ReceiverId] <- m
}

func justReceivedMessage(toUser int) (chat.Message, error) {
	if _, ok := channels[toUser]; !ok {
		channels[toUser] = make(chan chat.Message)
	}

	c := channels[toUser]

	return <- c, nil
}
