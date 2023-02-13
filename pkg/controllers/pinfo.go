package controllers

import (
	_ "fmt"
	"net/http"
	_ "net/mail"
	_ "unicode"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var UserInfo models.PersonalInfo

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {

	password := r.Header.Get("password")
	session := r.Header.Get("session")
	id := r.URL.Query().Get("id")

	pw, ex := models.ValidSession(id, password, session)

	if !ex {
		utils.ResponseWriter(w, "Content-Type", "application-json", http.StatusInternalServerError, "Invalid Request")
		return
	} else if !pw {
		utils.ResponseWriter(w, "Content-Type", "application-json", http.StatusInternalServerError, "Wrong Password")
		return
	}

	rawUserInfo := UserInfo
	utils.ParseBody(r, &rawUserInfo)

	if rawUserInfo.Username != "" {
		rawUserInfo.Username = utils.Hashing(rawUserInfo.Username)
	}

	if rawUserInfo.Birthday != "" {
		rawUserInfo.Birthday = utils.Hashing(rawUserInfo.Birthday)
	}

	if rawUserInfo.Address != "" {
		rawUserInfo.Address = utils.Hashing(rawUserInfo.Address)
	}

	if rawUserInfo.Gender != "" {
		rawUserInfo.Gender = utils.Hashing(rawUserInfo.Gender)
	}

	if rawUserInfo.Hobbies != "" {
		rawUserInfo.Hobbies = utils.Hashing(rawUserInfo.Hobbies)
	}

	models.UpdateUserInfo(id, rawUserInfo)

	utils.ResponseWriter(w, "Content-Type", "application-json", http.StatusOK, "Update Successfully")

}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {

	//Waiting for encrypt function

}
