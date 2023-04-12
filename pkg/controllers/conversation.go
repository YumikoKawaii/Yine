package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"github.com/gorilla/mux"
)

var (
	Conversation models.Conversation
	Group        models.Group
)

func ChangeNickname(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	coid := vars["coid"]
	nickname := r.Form.Get("nickname")

	if (Account.IsIdExist(coid) && Conversation.IsConversationBetween(id, coid)) || (Group.IsGroup(coid) && Conversation.IsMember(coid, id)) {
		Conversation.ChangeNickname(coid, id, nickname)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {

	//Temporary unavailable

}

func NewGroup(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gid := utils.Hashing(id + utils.RandomStringRunes(10))
	name := r.Form.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	Group.NewGroup(gid, name)
	Conversation.NewConnect(gid, id, utils.Admin)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(gid))

}

func AddMemeber(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	gid := vars["gid"]
	if !Group.IsGroup(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Group"))
		return
	}

	guest := r.Form.Get("guest")
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID"))
		return
	}

	if !Conversation.IsAdmin(gid, id) || !security.IsAccessable(id, guest) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	Conversation.NewConnect(gid, guest, utils.Member)
	w.WriteHeader(http.StatusOK)

}
func ChangeMemberRole(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	gid := vars["gid"]
	if !Group.IsGroup(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Group"))
		return
	}

	guest := r.Form.Get("guest")
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID"))
		return
	}

	role := r.Form.Get("role")

	if !Conversation.IsAdmin(gid, id) || !Conversation.IsMember(gid, guest) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	Conversation.ChangeRole(gid, guest, role)
	w.WriteHeader(http.StatusOK)

}

func ChangeGroupName(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	gid := vars["gid"]
	if !Group.IsGroup(gid) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	name := r.Form.Get("name")

	if !Conversation.IsMember(gid, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	Group.ChangeName(gid, name)
	w.WriteHeader(http.StatusOK)

}

func ChangeGroupAvatar(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	r.ParseMultipartForm(0)

	vars := mux.Vars(r)
	gid := vars["gid"]
	if !Group.IsGroup(gid) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !Conversation.IsMember(gid, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	url, status := utils.UploadImageToCloudiary(r, "avatar")

	w.WriteHeader(status)

	if status != http.StatusOK {
		return
	}

	Group.ChangeAvatar(gid, url)

}

func DeleteMember(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	gid := vars["gid"]
	if !Group.IsGroup(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Group"))
		return
	}

	guest := r.Form.Get("guest")
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID"))
		return
	}

	if !Conversation.IsAdmin(gid, id) || !Conversation.IsMember(gid, guest) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	Conversation.RemoveUser(gid, guest)
	w.WriteHeader(http.StatusOK)

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

	//Temporary unavailable

}

func FetchAllConversation(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	data := Conversation.GetAllConversation(id)
	res, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func FetchGroupChat(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	vars := mux.Vars(r)
	coid := vars["coid"]

	if !Conversation.IsMember(coid, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	g_data := Group.GetGroup(coid)
	members := Conversation.GetGroupMember(coid)

	data := struct {
		GroupData models.Group          `json:"gdata"`
		Members   []models.Conversation `json:"members"`
	}{GroupData: g_data, Members: members}

	res, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func FetchPersonalChat(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	vars := mux.Vars(r)
	coid := vars["coid"]

	if !Conversation.IsMember(coid, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user := Conversation.GetPartner(coid, id)
	partner := Conversation.GetPartner(id, coid)

	data := struct {
		User    models.Conversation `json:"user"`
		Partner models.Conversation `json:"partner"`
	}{User: user, Partner: partner}

	res, _ := json.Marshal(data)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func FetchBasicInfoConversation(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	vars := mux.Vars(r)
	coid := vars["coid"]

	if !Conversation.IsMember(coid, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if Group.IsGroup(coid) {

		g_data := Group.GetGroup(coid)

		lastest := Message.LastestGroupMessage(coid)

		lastestMessage := "You've received a message!"
		if lastest.Sender == id {
			lastestMessage = "You've sent a message!"
		}

		data := struct {
			Type    string    `json:"type"`
			Name    string    `json:"name"`
			Avatar  string    `json:"avatar"`
			Lastest string    `json:"lastest"`
			Recent  time.Time `json:"recent"`
		}{Type: "group", Name: g_data.Name, Avatar: g_data.Avatar, Lastest: lastestMessage, Recent: lastest.Time}

		res, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	} else {

		partner := Conversation.GetPartner(coid, id)
		profile := Profile.GetUserInfo(coid)
		lastest := Message.LastestPersonalMessage(id, coid)

		name := partner.Nickname
		if name == "" {
			name = profile.Username
		}

		lastestMessage := "You've received a message!"
		if lastest.Sender == id {
			lastestMessage = "You've sent a message!"
		}

		data := struct {
			Type    string    `json:"type"`
			Name    string    `json:"name"`
			Avatar  string    `json:"avatar"`
			Lastest string    `json:"lastest"`
			Recent  time.Time `json:"recent"`
		}{Type: "personal", Name: name, Avatar: profile.Avatar, Lastest: lastestMessage, Recent: lastest.Time}

		res, _ := json.Marshal(data)

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}
