package controllers

import (
	"encoding/json"
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

	pw := models.VerifyPassword(id, password)
	ex := models.VerifySession(id, session)

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
		rawUserInfo.Username, _ = utils.EncryptData(rawUserInfo.Username)
	}

	if rawUserInfo.Birthday != "" {
		rawUserInfo.Birthday, _ = utils.EncryptData(rawUserInfo.Birthday)
	}

	if rawUserInfo.Address != "" {
		rawUserInfo.Address, _ = utils.EncryptData(rawUserInfo.Address)
	}

	if rawUserInfo.Gender != "" {
		rawUserInfo.Gender, _ = utils.EncryptData(rawUserInfo.Gender)
	}

	if rawUserInfo.Hobbies != "" {
		rawUserInfo.Hobbies, _ = utils.EncryptData(rawUserInfo.Hobbies)
	}

	models.UpdateUserInfo(id, rawUserInfo)

	utils.ResponseWriter(w, "Content-Type", "application-json", http.StatusOK, "Update Successfully")

}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {

	session := r.Header.Get("session")
	id := r.URL.Query().Get("id")

	if models.VerifySession(id, session) {

		data := models.GetUserInfo(id)
		res, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	} else {
		utils.ResponseWriter(w, "Content-Type", "application-json", http.StatusOK, "Invalid Request")
	}

}
