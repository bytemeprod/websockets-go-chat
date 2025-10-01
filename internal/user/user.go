package user

import "time"

type Message struct {
	Id       string    `json:"id"`
	Text     string    `json:"text"`
	Date     time.Time `json:"date"`
	Sender   string    `json:"from"`
	Chatroom string    `json:"chatroom"`
}
