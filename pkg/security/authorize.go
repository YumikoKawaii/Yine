package security

import (
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
)

var Session models.Session

func Authorize(w http.ResponseWriter, r *http.Request) string {

	id := r.Header.Get("id")
	session := r.Header.Get("session")

	if !Session.VerifySession(id, session) {
		w.WriteHeader(http.StatusUnauthorized)
		return ""
	}

	return id

}
