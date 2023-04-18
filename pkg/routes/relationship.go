package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers/relationship"
	"github.com/gorilla/mux"
)

var RelationshipRoutes = func(router *mux.Router) {

	router.HandleFunc("/rlts", relationship.GetRelationship).Methods("GET")
	router.HandleFunc("/rlts/{id}", relationship.ModifyRelationship).Methods("PUT")

}
