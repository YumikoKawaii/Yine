package controllers

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var Account models.Account
var Session models.Session

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if _, err := mail.ParseAddress(email); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Email"))
		return
	}

	if security.VerifyPassword(password) {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Password"))
		return
	}

	if !Account.IsEmailExist(email) {

		id := utils.Hashing(email + utils.RandomStringRunes(10))

		Session.CreateSession(id)
		Profile.CreateEmptyRecord(id)

		resInfo := struct {
			ID      string `json:"id"`
			Session string `json:"session"`
		}{ID: id, Session: Session.GetSession(id)}

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

	//Wait until complete other features!

}

func ChangeEmail(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
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

	Account.UpdateEmail(id, new_email)
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

	if Account.VerifyPassword(id, old_password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !security.VerifyPassword(new_password) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	Account.UpdatePassword(id, new_password)
	w.WriteHeader(http.StatusOK)

}

func ChangeId(w http.ResponseWriter, r *http.Request) {
	//Wait until complete other features!
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	//Wait until complete other features!

}
