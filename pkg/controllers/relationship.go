package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"github.com/gorilla/mux"
)

var Relationship models.Relationship

func GetRelationship(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	result := Relationship.GetRelationship(id)
	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ModifyRelationship(w http.ResponseWriter, r *http.Request) {

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

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	action := r.Form.Get("action")

	switch action {
	case "sent", "accept", "reject", "unfriend", "block":
		if rlts := Relationship.GetRelationshipBetween(id, guest); rlts == utils.Block || rlts == utils.BeBlocked {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		switch action {
		case "sent":
			Relationship.SentRequest(id, guest)
		case "accept":
			Relationship.AcceptRequest(id, guest)
		case "reject", "unfriend":
			Relationship.CancelStatus(id, guest)
		case "block":
			Relationship.Block(id, guest)
		}
	case "unblock":
		if rlts := Relationship.GetRelationshipBetween(id, guest); rlts != utils.Block {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		Relationship.CancelStatus(id, guest)
	}

	w.WriteHeader(http.StatusOK)

}
