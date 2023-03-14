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

func (a Account) CreateAccount(id string, email string, password string) {

	newRecord := Account{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     email,
		Password:  password,
	}

	db.Create(&newRecord)

}

func (a Account) IsEmailExist(email string) bool {
	result := ""
	db.Raw("select email from accounts where email = ?", email).Scan(&result)
	return result != ""
}

func (a Account) IsIdExist(id string) bool {
	result := ""
	db.Raw("select * from accounts where id = ?", id).Scan(&result)
	return result != ""
}

func (a Account) VerifyAccount(email string, password string) bool {

	p := ""
	db.Raw("select password from accounts where email = ?", email).Scan(&p)
	return p == utils.Hashing(password)

}

func (a Account) UpdateEmail(id string, new_email string) {

	db.Exec("update accounts set email = ? where id = ?", new_email, id)

}

func (a Account) UpdatePassword(id string, new_password string) {

	db.Exec("update accounts set password = ? where id = ?", utils.Hashing(new_password), id)

}

func (a Account) VerifyPassword(Id string, password string) bool {

	p := ""
	db.Raw("select password from accounts where id = ?", Id).Scan(&p)
	return p == utils.Hashing(password)

}

func (a Account) GetID(email string) string {

	result := ""
	db.Raw("select id from accounts where email = ?", email).Scan(&result)
	return result
}

func (a Account) DeleteAccount(id string) {

	//Wait until complete other features!

}
