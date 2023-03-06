package controllers

import (
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
)

var Relationship models.Relationship

func SentRequest(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get("id")
	session := r.Header.Get("session")

	if !models.VerifySession(id, session) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	guest := r.Form.Get("guest")

	models.ProcessRequest(id, guest)
	w.WriteHeader(http.StatusOK)

}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get("id")
	session := r.Header.Get("session")

	if !models.VerifySession(id, session) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	guest := r.Form.Get("guest")

	models.AcceptRequest(id, guest)

}

func ModifyRelationship(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get("id")
	session := r.Header.Get("session")

	if !models.VerifySession(id, session) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	guest := r.Form.Get("guest")
	request := r.Form.Get("request")

	switch request {
	case "block":
		if models.GetRelationship(id, guest) == "be blocked" {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		models.Block(id, guest)
		w.WriteHeader(http.StatusOK)
	case "unblock":
		if models.GetRelationship(id, guest) != "blocked" {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		models.CancelStatus(id, guest)
		w.WriteHeader(http.StatusOK)
	case "unfriend":
		models.CancelStatus(id, guest)
		w.WriteHeader(http.StatusOK)
	}

}
