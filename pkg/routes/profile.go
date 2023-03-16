package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var ProfileRoutes = func(router *mux.Router) {

	router.HandleFunc("/p/{id}", controllers.GetProfile).Methods("GET")
	router.HandleFunc("/p/a", controllers.UpdateAvatar).Methods("PUT")
	router.HandleFunc("/p", controllers.UpdateRegularInfo).Methods("PUT")

}
