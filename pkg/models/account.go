package models

import (
	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Account{})
}

func (a *Account) CreateAccount() *Account {
	db.Create(&a)
	return a
}

func IsExist(Email string) bool {

	r := Account{}
	db.Raw("select id from accounts where email = ?", Email).Scan(&r)
	return r != Account{}

}

func VerifyAccount(Email string, Password string) bool {

	r := Account{}
	db.Raw("select email, password from accounts where email = ?", Email).Scan(&r)
	return Password == r.Password

}

func DeleteAccount(Email string) {

	db.Exec("delete from accounts where id = ?", utils.Hashing(Email))

}

func UpdateAccount(Email string, NewPassword string) {
	db.Exec("set sql_safe_updates = 0")
	db.Exec("update accounts set password = ? where email = ?", NewPassword, Email)
	db.Exec("set sql_safe_updates = 1")
}
