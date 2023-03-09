package models

import "github.com/YumikoKawaii/Yine/pkg/config"

type Chat struct {
	CID   string `json:"c_id" gorm:"primarykey"`
	ID    string `json:"id"`
	Guest string `json:"guest"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Chat{})
}

func (c Chat) NewChat(id string, guest string, cid string) {

	newRecord := Chat{
		CID:   cid,
		ID:    id,
		Guest: guest,
	}

	db.Create(&newRecord)

}

func (c Chat) RemoveChat(cid string) {

	db.Exec("delete from chats where c_id = ?", cid)

}
