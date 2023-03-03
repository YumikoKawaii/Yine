package models

import (
	"time"

	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

type Account struct {
	ID        string `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `json:"email"`
	Password  string `json:"password"`
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

func IsEmailExist(email string) bool {
	r := Account{}
	db.Raw("select * from accounts where email = ?", email).Scan(&r)
	return r != Account{}
}

func IsIdExist(id string) bool {

	r := Account{}
	db.Raw("select * from accounts where id = ?", id).Scan(&r)
	return r != Account{}

}

func VerifyAccount(email string, password string) bool {

	var p string = ""
	db.Raw("select password from accounts where email = ?", email).Scan(&p)
	return p == utils.Hashing(password)

}

func UpdateEmail(id string, new_email string) {

	db.Exec("set sql_safe_updates = 0")
	db.Exec("update accounts set email = ? where id = ?", new_email, id)
	db.Exec("update accounts set updated_at = ? where id = ?", time.Now(), id)
	db.Exec("set sql_safe_updates = 1")

}

func UpdatePassword(id string, new_password string) {

	db.Exec("set sql_safe_updates = 0")
	db.Exec("update accounts set password = ? where id = ?", utils.Hashing(new_password), id)
	db.Exec("update accounts set updated_at = ? where id = ?", time.Now(), id)
	db.Exec("set sql_safe_updates = 1")

}

func VerifyPassword(Id string, password string) bool {

	var p string = ""
	db.Raw("select password from accounts where id = ?", Id).Scan(&p)

	return p == utils.Hashing(password)

}

func DeleteAccount(id string) {

	//Wait until complete other features!

}
