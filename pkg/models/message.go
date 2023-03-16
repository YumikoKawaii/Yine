package models

import (
	"time"
)

const (
	MessageSent     = "sent"
	MessageReceiver = "receiver"
	MessageRead     = "read"
)

type Message struct {
	MID      uint32 `json:"mid" gorm:"primaryKey;autoIncrement:true"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Time     time.Time
}

func init() {
	db.AutoMigrate(&Message{})
}

func (m Message) NewMessage(sender string, receiver string, t string, content string) {

	db.Exec("insert into messages (sender, receiver, type, content, time) values (?, ?, ?, ?, ?)", sender, receiver, t, content, time.Now())

}

func (m Message) FetchAllPersonalConversation(sender string, receiver string) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where (sender = ? and receiver = ?) or (sender = ? and receiver = ?)", sender, receiver, receiver, sender).Scan(&result)
	return result

}

func (m Message) FetchPersonalConversation(sender string, receiver string, mid uint32) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where (sender = ? and receiver = ?) or (sender = ? and receiver = ?) and m_id > ?", sender, receiver, receiver, sender, mid).Scan(&result)
	return result

}

func (m Message) FetchAllGroupConversation(receiver string) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where receiver = ?", receiver).Scan(&result)
	return result

}

func (m Message) FetchGroupConversation(receiver string, mid uint32) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where receiver = ? and m_id > ?", receiver, mid).Scan(&result)
	return result
}
