package models

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/config"
)

const (
	MessageSent     = "sent"
	MessageReceiver = "receiver"
	MessageRead     = "read"
)

type Message struct {
	CoID    string `json:"co_id" gorm:"primarykey"`
	MID     int    `json:"m_id" gorm:"autoincrement"`
	Sender  string `json:"sender"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Time    time.Time
	Status  string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Profile{})
}

func (m Message) NewMessage(coid string, sender string, t string, content string) {

	db.Create(Message{
		CoID:    coid,
		Sender:  sender,
		Type:    t,
		Content: content,
		Time:    time.Now(),
		Status:  MessageSent,
	})

}

func (m Message) GetMessage(coid string, limit int) []Message {

	result := make([]Message, limit)

	db.Raw("select * from messages where co_id = ? order by time desc limit ?", coid, limit).Scan(&result)
	return result

}

func (m Message) ChangeMessageStatus(coid string, mid int, status string) {

	db.Exec("update messages set status = ? where co_id = ? and m_id = ?", status, coid, mid)

}
