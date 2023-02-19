package models

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

type Session struct {
	ID      string `json:"id" gorm:"primarykey"`
	Session string `json:"session"`
	Expired string `json:"expired"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Session{})
}

func CreateSession(ID string) {

	newSession := Session{}
	newSession.ID = ID
	newSession.Session = utils.Hashing(ID + string(rune(time.Now().UnixNano())))

	newSession.Expired = time.Now().AddDate(0, 0, 10).Format(utils.TimeFormat)
	db.Create(&newSession)

}

func VerifySession(ID string, key string) bool {

	var e string = ""
	db.Raw("select expired from sessions where id = ? and session = ?", ID, key).Scan(&e)

	if e == "" {
		return false
	}

	date, _ := time.Parse(utils.TimeFormat, e)

	return time.Now().Before(date)

}
