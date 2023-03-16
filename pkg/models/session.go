package models

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/utils"
)

type Session struct {
	User    string `json:"id"`
	Session string `json:"session"`
	Expired string `json:"expired"`
}

func init() {
	db.AutoMigrate(&Session{})
}

func (s Session) CreateSession(user string) {

	db.Create(Session{
		User:    user,
		Session: utils.Hashing(user + string(rune(time.Now().UnixNano()))),
		Expired: time.Now().AddDate(0, 0, 10).Format(utils.TimeFormat),
	})

}

func (s Session) GetSession(user string) string {

	result := ""
	db.Raw("select session from sessions where user = ?", user).Scan(&result)
	return result

}

func (s Session) VerifySession(user string, key string) bool {

	e := ""
	db.Raw("select expired from sessions where user = ? and session = ?", user, key).Scan(&e)

	if e == "" {
		return false
	}

	date, _ := time.Parse(utils.TimeFormat, e)

	return time.Now().Before(date)

}
