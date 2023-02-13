package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var PersonalInfomationRoutes = func(router *mux.Router) {

	router.HandleFunc("/pinfo/", controllers.UpdateUserInfo).Methods("PUT")
	router.HandleFunc("/pinfo/{id}/", controllers.GetUserInfo).Methods("PUT")

}
