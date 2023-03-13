package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var MessageRoutes = func(router *mux.Router) {

	router.HandleFunc("/c/", controllers.ConnectToSocket).Methods("GET")
	router.HandleFunc("/m/", controllers.SendMessage).Methods("POST")

}
