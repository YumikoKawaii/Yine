package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var ProfileRoutes = func(router *mux.Router) {

	router.HandleFunc("/profile/", controllers.GetUserInfo).Methods("GET")
	router.HandleFunc("/profile/avatar/", controllers.UpdateAvatar).Methods("PUT")
	router.HandleFunc("/profile/", controllers.UpdateRegularInfo).Methods("PUT")

}
