package controllers

import (
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var Conversation models.Conversation
var Chat models.Chat
var Group models.Group
var Setting models.Setting

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
	if !Conversation.IsConversationExist(coid) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if Conversation.GetRole(coid, id) == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	nickname := r.Form.Get("nickname")

	Conversation.ChangeNickname(coid, id, nickname)
	w.WriteHeader(http.StatusOK)

}

func NewChat(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	guest := r.Form.Get("guest")
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	relationship := Relationship.GetRelationship(id, guest)

	if relationship == utils.Friend || (relationship != utils.Block && relationship != utils.BeBlocked && Setting.GetPublic(guest)) {
		cid := utils.Hashing(id + guest + utils.RandomStringRunes(10))
		Chat.NewChat(id, guest, cid)
		Conversation.NewConnect(cid, id, "Member")
		Conversation.NewConnect(cid, guest, "Member")
		w.Write([]byte(cid))
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotAcceptable)

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

	Group.NewGroup(gid, name)
	Conversation.NewConnect(gid, id, "Admin")
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
	if !Conversation.IsConversationExist(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Conversation"))
		return
	}

	guest := r.Form.Get("guest")
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID"))
		return
	}

	if Conversation.GetRole(gid, id) != "Admin" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	Conversation.NewConnect(gid, guest, "Member")
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
	if !Conversation.IsConversationExist(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Conversation"))
		return
	}

	guest := r.Form.Get("guest")
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID"))
		return
	}

	role := r.Form.Get("role")

	if Conversation.GetRole(gid, id) == "Admin" {

		Conversation.ChangeRole(gid, guest, role)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotAcceptable)

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
	if !Conversation.IsConversationExist(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Conversation"))
		return
	}

	name := r.Form.Get("name")

	if Conversation.GetRole(gid, id) == "" {
		w.WriteHeader(http.StatusNotAcceptable)
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
		w.Write([]byte("Conversation"))
		return
	}

	if Conversation.GetRole(gid, id) == "" {
		w.WriteHeader(http.StatusNotAcceptable)
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
	if !Conversation.IsConversationExist(gid) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Conversation"))
		return
	}

	guest := r.Form.Get("guest")
	if !Account.IsIdExist(guest) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID"))
		return
	}

	if Conversation.GetRole(gid, id) != "Admin" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	Conversation.RemoveUser(gid, guest)
	w.WriteHeader(http.StatusOK)

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

	//Temporary unavailable

}
