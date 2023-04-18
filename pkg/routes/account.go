package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers/account"
	"github.com/gorilla/mux"
)

var AccountRoutes = func(router *mux.Router) {
	router.HandleFunc("/a", account.CreateAccount).Methods("POST")
	router.HandleFunc("/a", account.Login).Methods("GET")
	router.HandleFunc("/a/email", account.ChangeEmail).Methods("PUT")
	router.HandleFunc("/a/password", account.ChangePassword).Methods("PUT")
	router.HandleFunc("/a/id", account.ChangeId).Methods("PUT")
	router.HandleFunc("/a", account.DeleteAccount).Methods("DELETE")
}
