package controllers

import (
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
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
	request := r.Form.Get("request")

	switch request {
	case "block":
		if Relationship.GetRelationship(id, guest) == "be blocked" {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		Relationship.Block(id, guest)
		w.WriteHeader(http.StatusOK)
	case "unblock":
		if Relationship.GetRelationship(id, guest) != "blocked" {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		Relationship.CancelStatus(id, guest)
		w.WriteHeader(http.StatusOK)
	case "unfriend":
		Relationship.CancelStatus(id, guest)
		w.WriteHeader(http.StatusOK)
	}

}
