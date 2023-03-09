package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var Profile models.Profile

func UpdateAvatar(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	url, status := utils.UploadImageToCloudiary(r, "avatar")

	w.WriteHeader(status)

	if status != http.StatusOK {
		return
	}

	Profile.UpdateField(id, "avatar", url)

}

func UpdateRegularInfo(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	field := r.Form.Get("field")
	value := r.Form.Get("value")

	if field == "" || value == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if !Profile.UpdateField(id, field, value) {
		w.WriteHeader(http.StatusTooEarly)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetProfile(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	guest := r.URL.Query().Get("id")

	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	profile := Profile.GetUserInfo(guest)

	if id != guest && !security.IsAccessable(id, guest) {

		profile.Address = ""
		profile.Birthday = ""
		profile.Hobbies = ""

	}

	res, _ := json.Marshal(profile)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
