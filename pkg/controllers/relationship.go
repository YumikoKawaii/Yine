package controllers

import (
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
)

var Relationship models.Relationship

func SentRequest(w http.ResponseWriter, r *http.Request) {

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

	Relationship.ProcessRequest(id, guest)
	w.WriteHeader(http.StatusOK)

}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {

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

	Relationship.AcceptRequest(id, guest)
}

func RejectRequest(w http.ResponseWriter, r *http.Request) {

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

	Relationship.CancelStatus(id, guest)

}

func ModifyRelationship(w http.ResponseWriter, r *http.Request) {

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

	request := r.Form.Get("request")

	relationship := Relationship.GetRelationship(id, guest)

	switch request {
	case "block":
		if relationship == utils.BeBlocked {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		Relationship.Block(id, guest)
		w.WriteHeader(http.StatusOK)
	case "unblock":
		if relationship != utils.Block {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		Relationship.CancelStatus(id, guest)
		w.WriteHeader(http.StatusOK)
	case "unfriend":
		if relationship == utils.Friend {
			Relationship.CancelStatus(id, guest)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotAcceptable)

	}

}
