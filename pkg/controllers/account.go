package controllers

import (
	"net/http"
	"net/mail"
	"unicode"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var Account models.Account

func verifyPassword(p string) bool {

	length := len(p) >= 8
	number := false
	upper := false
	special := false

	for _, c := range p {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}

	}

	return length && number && upper && special
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	newAccountInfo := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	utils.ParseBody(r, newAccountInfo)

	if _, err := mail.ParseAddress(newAccountInfo.Email); err != nil {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Invalid Email")
		return
	}

	if !verifyPassword(newAccountInfo.Password) {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Invalid Password")
		return
	}

	newAccount := &models.Account{}
	newAccount.Email = utils.Hashing(newAccountInfo.Email)

	if !models.IsExist(newAccount.Email) {

		newAccount.Password = utils.Hashing(newAccountInfo.Password)
		newAccount.ID = utils.Hashing(newAccount.Email)
		newAccount.CreateAccount()
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusOK, "Created successfully")
		return

	} else {

		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Email existed")
		return
	}

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	accountInfo := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	utils.ParseBody(r, accountInfo)

	accountInfo.Email = utils.Hashing(accountInfo.Email)
	accountInfo.Password = utils.Hashing(accountInfo.Password)

	if models.IsExist(accountInfo.Email) {
		if models.VerifyAccount(accountInfo.Email, accountInfo.Password) {
			models.DeleteAccount(accountInfo.Email)
			utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusOK, "Deleted")
		} else {
			utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Wrong Password")
		}
	} else {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Account is not exist")
	}

}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {

	accountInfo := &struct {
		Email       string `json:"email"`
		OldPassword string `json:"oldpassword"`
		NewPassword string `json:"newpassword"`
	}{}

	utils.ParseBody(r, accountInfo)

	accountInfo.Email = utils.Hashing(accountInfo.Email)
	accountInfo.OldPassword = utils.Hashing(accountInfo.OldPassword)

	if models.IsExist(accountInfo.Email) {
		if models.VerifyAccount(accountInfo.Email, accountInfo.OldPassword) {

			if !verifyPassword(accountInfo.NewPassword) {
				utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "New Password is invalid")
			} else {
				models.UpdateAccount(accountInfo.Email, utils.Hashing(accountInfo.NewPassword))
				utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusOK, "Password updated")
			}

		} else {
			utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Wrong Password")
		}
	} else {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Account is not exist")
	}

}
