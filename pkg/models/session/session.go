package session

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Session struct {
	ID      string `json:"id"`
	Session string `json:"session"`
	Expired string `json:"expired"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (s Session) CreateSession(id string) {

	db.Create(Session{
		ID:      id,
		Session: utils.Hashing(id + string(rune(time.Now().UnixNano()))),
		Expired: time.Now().AddDate(0, 0, 10).Format(utils.TimeFormat),
	})

}

func (s Session) GetSession(id string) string {

	result := ""
	db.Raw("select session from sessions where id = ?", id).Scan(&result)
	return result

}

func (s Session) VerifySession(id string, key string) bool {

	e := ""
	db.Raw("select expired from sessions where id = ? and session = ?", id, key).Scan(&e)

	if e == "" {
		return false
	}

	date, _ := time.Parse(utils.TimeFormat, e)

	return time.Now().Before(date)

}
