package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers/message"
	"github.com/gorilla/mux"
)

var MessageRoutes = func(router *mux.Router) {

	router.HandleFunc("/c", message.ConnectToSocket).Methods("GET")
	router.HandleFunc("/m/{coid}", message.FetchMessage).Methods("GET")
	router.HandleFunc("/m/{coid}", message.SendMessage).Methods("POST")

}
