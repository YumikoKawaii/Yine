package message

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/config"
	"gorm.io/gorm"
)

const (
	MessageSent     = "sent"
	MessageReceiver = "receiver"
	MessageRead     = "read"
)

var (
	db *gorm.DB
)

type Message struct {
	MID      uint32    `json:"mess_id"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Type     string    `json:"type"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (m Message) NewMessage(sender string, receiver string, t string, content string) (mess Message) {

	db.Exec("insert into messages (sender, receiver, type, content, time) values (?, ?, ?, ?, ?)", sender, receiver, t, content, time.Now())

	return Message{
		Sender:   sender,
		Receiver: receiver,
		Type:     t,
		Content:  content,
		Time:     time.Now(),
	}

}

func (m Message) FetchAllPersonalConversation(sender string, receiver string) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where (sender = ? and receiver = ?) or (sender = ? and receiver = ?)", sender, receiver, receiver, sender).Scan(&result)
	return result

}

func (m Message) FetchPersonalConversation(sender string, receiver string, mess_id uint32) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where (sender = ? and receiver = ?) or (sender = ? and receiver = ?) and mess_id > ?", sender, receiver, receiver, sender, mess_id).Scan(&result)
	return result

}

func (m Message) FetchAllGroupConversation(receiver string) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where receiver = ?", receiver).Scan(&result)
	return result

}

func (m Message) FetchGroupConversation(receiver string, mess_id uint32) []Message {

	result := make([]Message, 0)
	db.Raw("select * from messages where receiver = ? and mess_id > ?", receiver, mess_id).Scan(&result)
	return result
}

func (m Message) LastestPersonalMessage(sender string, receiver string) Message {

	result := Message{}

	db.Raw("select * from messages where (sender = ? and receiver = ?) or (sender = ? and receiver = ?) order by mess_id DESC limit 1", sender, receiver, receiver, sender).Scan(&result)

	return result

}

func (m Message) LastestGroupMessage(receiver string) Message {

	result := Message{}

	db.Raw("select * from messages where receiver = ? order by mess_id DESC limit 1", receiver).Scan(&result)

	return result

}
