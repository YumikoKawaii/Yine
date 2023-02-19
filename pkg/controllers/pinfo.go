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

	userInfo := UserInfo
	utils.ParseBody(r, &userInfo)

	models.UpdateUserInfo(id, userInfo)

	utils.ResponseWriter(w, "Content-Type", "application-json", http.StatusOK, "Update Successfully")

}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	data := models.GetUserInfo(id)
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
