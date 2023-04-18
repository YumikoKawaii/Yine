package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers/profile"
	"github.com/gorilla/mux"
)

var ProfileRoutes = func(router *mux.Router) {

	router.HandleFunc("/p/{id}", profile.GetProfile).Methods("GET")
	router.HandleFunc("/p/a/{id}", profile.GetProfileAvatarAndUsername).Methods("GET")
	router.HandleFunc("/p/a", profile.UpdateAvatar).Methods("PUT")
	router.HandleFunc("/p", profile.UpdateRegularInfo).Methods("PUT")

}
