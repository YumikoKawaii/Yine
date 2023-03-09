package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var AccountRoutes = func(router *mux.Router) {
	router.HandleFunc("/account/", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/account/", controllers.Login).Methods("GET")
	router.HandleFunc("/account/email", controllers.ChangeEmail).Methods("PUT")
	router.HandleFunc("/account/password", controllers.ChangePassword).Methods("PUT")
	router.HandleFunc("/account/id", controllers.ChangeId).Methods("PUT")
	router.HandleFunc("/account/", controllers.DeleteAccount).Methods("DELETE")
}
