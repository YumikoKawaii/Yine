package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var RelationshipRoutes = func(router *mux.Router) {

	router.HandleFunc("/relationship/", controllers.SentRequest).Methods("POST")
	router.HandleFunc("/relationship/a", controllers.AcceptRequest).Methods("POST")
	router.HandleFunc("/relationship/", controllers.ModifyRelationship).Methods("PUT")
}
