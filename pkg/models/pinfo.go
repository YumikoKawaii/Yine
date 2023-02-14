package models

import (
	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

type PersonalInfo struct {
	gorm.Model
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Birthday string `json:"Birthday"`
	Address  string `json:"Address"`
	Gender   string `json:"Gender"`
	Hobbies  string `json:"Hobbies"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&PersonalInfo{})
}

func CreateEmptyRecord(ID string) {

	newRecord := &PersonalInfo{}
	newRecord.ID = ID
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

	db.Exec("set sql_safe_updates = 0")
	db.Exec("update personal_infos set username = ?, birthday = ?, address = ?, gender = ?, hobbies = ? where id = ?", data.Username, data.Birthday, data.Address, data.Gender, data.Hobbies, Id)
	db.Exec("set sql_safe_updates = 1")
}

func GetUserInfo(Id string) PersonalInfo {

	data := PersonalInfo{}

	db.Raw("select username, birthday, address, gender, hobbies where id = ?", Id).Scan(&data)

	data.Username, _ = utils.DecryptData(data.Username)
	data.Birthday, _ = utils.DecryptData(data.Birthday)
	data.Address, _ = utils.DecryptData(data.Address)
	data.Gender, _ = utils.DecryptData(data.Gender)
	data.Hobbies, _ = utils.DecryptData(data.Hobbies)

	return data

}
