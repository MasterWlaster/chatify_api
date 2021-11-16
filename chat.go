package chat

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	StatusOnline bool   `json:"status_online"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

type Message struct {
	Id         int    `json:"id"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id" binding:"required"`
	TimeSent   int64  `json:"time_sent"`
	Text       string `json:"text" binding:"required"`
}
