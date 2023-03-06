package controllers

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"time"
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

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if _, err := mail.ParseAddress(email); err != nil {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusNotAcceptable, "Email")
		return
	}

	if !verifyPassword(password) {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusNotAcceptable, "Password")
		return
	}

	if !models.IsEmailExist(email) {

		newAccount := &models.Account{

			ID:        utils.Hashing(email + utils.RandomStringRunes(10)),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     email,
			Password:  utils.Hashing(password),
		}

		models.CreateSession(newAccount.ID)
		newAccount.CreateAccount()
		models.CreateEmptyRecord(newAccount.ID)

		resInfo := struct {
			ID      string `json:"id"`
			Session string `json:"session"`
		}{ID: newAccount.ID, Session: models.GetSession(newAccount.ID)}

		res, _ := json.Marshal(resInfo)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

		return

	} else {

		w.WriteHeader(http.StatusConflict)
		return

	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Unidentified")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if !models.IsEmailExist(email) {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Account is not exist")
		return
	}

	if !models.VerifyAccount(email, password) {
		utils.ResponseWriter(w, "Content-Type", "application/json", http.StatusInternalServerError, "Wrong password")
		return
	}

	//Wait until complete other features!

}

func ChangeEmail(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get("id")
	session := r.Header.Get("session")

	if !models.VerifySession(id, session) {

		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	err := r.ParseForm()

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	new_email := r.Form.Get("email")

	if _, err := mail.ParseAddress(new_email); err != nil {

		w.WriteHeader(http.StatusNotAcceptable)
		return

	}

	models.UpdateEmail(id, new_email)
	w.WriteHeader(http.StatusOK)

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get("id")

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	old_password := r.Form.Get("old_password")
	new_password := r.Form.Get("new_password")

	if !models.VerifyPassword(id, old_password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !verifyPassword(new_password) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	models.UpdatePassword(id, new_password)
	w.WriteHeader(http.StatusOK)

}

func ChangeId(w http.ResponseWriter, r *http.Request) {
	//Wait until complete other features!
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	//Wait until complete other features!

}
