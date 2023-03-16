package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var AccountRoutes = func(router *mux.Router) {
	router.HandleFunc("/a", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/a", controllers.Login).Methods("GET")
	router.HandleFunc("/a/email", controllers.ChangeEmail).Methods("PUT")
	router.HandleFunc("/a/password", controllers.ChangePassword).Methods("PUT")
	router.HandleFunc("/a/id", controllers.ChangeId).Methods("PUT")
	router.HandleFunc("/a", controllers.DeleteAccount).Methods("DELETE")
}
