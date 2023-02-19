package models

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/config"
)

type PersonalInfo struct {
	ID        string `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"Username"`
	Birthday  string `json:"Birthday"`
	Address   string `json:"Address"`
	Gender    string `json:"Gender"`
	Hobbies   string `json:"Hobbies"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&PersonalInfo{})
}

func CreateEmptyRecord(ID string) {

	newRecord := &PersonalInfo{}
	newRecord.ID = ID
	newRecord.CreatedAt = time.Now()
	newRecord.UpdatedAt = time.Now()
	db.Create(newRecord)

}

func UpdateUserInfo(Id string, newInfo PersonalInfo) {

	data := PersonalInfo{}
	db.Raw("select id, username, birthday, address, gender, hobbies from personal_infos where id = ?", Id).Scan(&data)

	if newInfo.Username != "" {
		data.Username = newInfo.Username
	}

	if newInfo.Birthday != "" {
		data.Birthday = newInfo.Birthday
	}

	if newInfo.Address != "" {
		data.Address = newInfo.Address
	}

	if newInfo.Gender != "" {
		data.Gender = newInfo.Gender
	}

	if newInfo.Hobbies != "" {
		data.Hobbies = newInfo.Hobbies
	}

	data.UpdatedAt = time.Now()

	db.Exec("set sql_safe_updates = 0")
	db.Exec("update personal_infos set updated_at = ?, username = ?, birthday = ?, address = ?, gender = ?, hobbies = ? where id = ?", data.UpdatedAt, data.Username, data.Birthday, data.Address, data.Gender, data.Hobbies, Id)
	db.Exec("set sql_safe_updates = 1")
}

func GetUserInfo(Id string) PersonalInfo {

	data := PersonalInfo{}

	db.Raw("select username, birthday, address, gender, hobbies where id = ?", Id).Scan(&data)

	return data

}
