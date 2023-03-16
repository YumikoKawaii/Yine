package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var RelationshipRoutes = func(router *mux.Router) {

	router.HandleFunc("/rlts", controllers.GetRelationship).Methods("GET")
	router.HandleFunc("/rlts/{id}", controllers.ModifyRelationship).Methods("PUT")

}
