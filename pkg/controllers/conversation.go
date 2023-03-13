package controllers

import (
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var (
	Conversation models.Conversation
	Group        models.Group
	Setting      models.Setting
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

	coid := r.Form.Get("coid")
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

	gid := r.Form.Get("gid")
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

	if !Conversation.IsAdmin(gid, id) {
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

	gid := r.Form.Get("gid")
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

	if !Conversation.IsAdmin(gid, id) {
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

	gid := r.Form.Get("gid")
	if !Group.IsGroup(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Group"))
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

	gid := r.FormValue("gid")
	if !Conversation.IsConversationExist(gid) {
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

	gid := r.Form.Get("gid")
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

	if !Conversation.IsAdmin(gid, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	Conversation.RemoveUser(gid, guest)
	w.WriteHeader(http.StatusOK)

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

	//Temporary unavailable

}
