package models

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID      string `json:"id"`
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

func ValidSession(ID string, Password string, key string) bool {

	var e string = ""
	db.Raw("select expired from sessions where id = ? and session = ?", ID, key).Scan(&e)

	if e == "" {
		return false
	}

	date, _ := time.Parse(utils.TimeFormat, e)

	if time.Now().After(date) {
		return false
	}

	p := utils.Hashing(Password)
	var data string = ""
	db.Raw("select id from accounts where id = ? and password = ?", ID, p).Scan(&data)

	return data != ""

}
