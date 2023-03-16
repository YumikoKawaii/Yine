package models

import (
	"time"
)

type Profile struct {
	ID         string    `json:"id" gorm:"primarykey"`
	LastUpdate time.Time `json:"last_update"`
	Avatar     string    `json:"avatar"`
	Username   string    `json:"username"`
	Birthday   string    `json:"birthday"`
	Address    string    `json:"address"`
	Gender     string    `json:"gender"`
	Hobbies    string    `json:"hobbies"`
}

func init() {
	db.AutoMigrate(&Profile{})
}

func (p Profile) CreateEmptyRecord(id string) {

	db.Create(Profile{
		ID:         id,
		LastUpdate: time.Now().AddDate(0, 0, -61),
	})

}

func (p Profile) verifyUpdateTime(id string) bool {
	var lastUpdate time.Time

	db.Raw("select last_update from profiles where id = ?", id).Scan(&lastUpdate)

	var nextUpdate time.Time = lastUpdate.AddDate(0, 0, 60)

	return time.Now().After(nextUpdate)
}

func (p Profile) UpdateField(id string, field string, value string) bool {

	if field != "username" || (field == "username" && p.verifyUpdateTime(id)) {

		db.Exec("update profiles set "+field+" = ? where id = ?", value, id)

		if field == "username" {
			db.Exec("update profiles set last_update = ? where id = ?", time.Now(), id)
		}

		return true

	}

	return false
}

func (p Profile) GetUserInfo(id string) Profile {

	data := Profile{}

	db.Raw("select * from profiles where id = ?", id).Scan(&data)

	return data

}
