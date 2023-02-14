package models

import (
	"github.com/YumikoKawaii/Yine/pkg/config"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

type Account struct {
	gorm.Model
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func ValidEmail(email string) bool {
	r := Account{}
	db.Raw("select * from accounts where email = ?", email).Scan(&r)
	return r != Account{}
}

func IsExist(id string) bool {

	r := Account{}
	db.Raw("select * from accounts where id = ?", id).Scan(&r)
	return r != Account{}

}

func VerifyAccount(Email string, Password string) bool {

	r := Account{}
	db.Raw("select email, password from accounts where email = ?", Email).Scan(&r)
	return Password == r.Password

}

func DeleteAccount(id string) {

	db.Exec("delete from accounts where id = ?", id)

}

func UpdateAccount(id string, new_password string) {
	db.Exec("set sql_safe_updates = 0")
	db.Exec("update accounts set password = ? where id = ?", utils.Hashing(new_password), id)
	db.Exec("set sql_safe_updates = 1")
}

func VerifyPassword(Id string, password string) bool {

	r := &struct {
		Password string `json:"password"`
	}{}
	db.Raw("select password from accounts where id = ?", Id).Scan(&r)

	return utils.Hashing(password) == r.Password

}
