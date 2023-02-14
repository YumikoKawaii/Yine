package controllers

import (
	"net/http"
	"net/mail"
	"unicode"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"github.com/gorilla/mux"
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
	newAccount.Email = newAccountInfo.Email

	if !models.ValidEmail(newAccount.Email) {

		newAccount.Password = utils.Hashing(newAccountInfo.Password)
		newAccount.ID = utils.Hashing(newAccount.Email + utils.RandomStringRunes(10))
		models.CreateSession(newAccount.ID)
		newAccount.CreateAccount()
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusOK, newAccount.ID)
		models.CreateEmptyRecord(newAccount.ID)
		return

	} else {

		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Email existed")
		return

	}

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	password := r.Header.Get("password")

	if models.IsExist(id) {
		if models.VerifyPassword(id, password) {
			models.DeleteAccount(id)
			utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusOK, "Deleted")
		} else {
			utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Wrong Password")
		}

	} else {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Account is not exist")
	}

}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	password := r.Header.Get("password")

	n := &struct {
		Password string `json:"nPassword"`
	}{}

	utils.ParseBody(r, n)

	if models.VerifyPassword(id, password) {
		if verifyPassword(n.Password) {
			models.UpdateAccount(id, n.Password)
			utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusOK, "Succesfully")
		} else {
			utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "New Password is not valid")
		}
	} else {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Wrong Password")
	}

}
