package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAccountRoutes = func(router *mux.Router) {
	router.HandleFunc("/account/", controllers.CreateAccount).Methods("POST")
	//router.HandleFunc("/account/", controllers.UserLogin).Methods("GET")
	router.HandleFunc("/account/{id}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/{id}", controllers.DeleteAccount).Methods("DELETE")
}
