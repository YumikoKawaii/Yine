package profile

import (
	"encoding/json"
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models/account"
	"github.com/YumikoKawaii/Yine/pkg/models/profile"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"github.com/gorilla/mux"
)

var (
	Account account.Account
	Profile profile.Profile
)

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
	w.Write([]byte(url))

}

func UpdateRegularInfo(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	if err := r.ParseForm(); err != nil {
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

	vars := mux.Vars(r)
	guest := vars["id"]
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

func GetProfileAvatarAndUsername(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	vars := mux.Vars(r)
	guest := vars["id"]
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	username, avatar := Profile.GetUserAvatarAndUsername(guest)

	res_info := struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	}{Username: username, Avatar: avatar}

	res, _ := json.Marshal(res_info)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
