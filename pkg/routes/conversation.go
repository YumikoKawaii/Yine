package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var ConversationRoutes = func(router *mux.Router) {

	router.HandleFunc("/cvs", controllers.FetchAllConversation).Methods("GET")
	router.HandleFunc("/cvs/p/d/{coid}", controllers.FetchPersonalChat).Methods("GET")
	router.HandleFunc("/cvs/g/d/{coid}", controllers.FetchGroupChat).Methods("GET")
	router.HandleFunc("/cvs/{coid}", controllers.FetchBasicInfoConversation).Methods("GET")
	router.HandleFunc("/cvs/{coid}", controllers.ChangeNickname).Methods("PUT")

	router.HandleFunc("/cvs/g", controllers.NewGroup).Methods("POST")
	router.HandleFunc("/cvs/g/am/{gid}", controllers.AddMemeber).Methods("POST")
	router.HandleFunc("/cvs/g/cmr/{gid}", controllers.ChangeMemberRole).Methods("PUT")
	router.HandleFunc("/cvs/g/cn/{gid}", controllers.ChangeGroupName).Methods("PUT")
	router.HandleFunc("/cvs/g/ca/{gid}", controllers.ChangeGroupAvatar).Methods("PUT")
	router.HandleFunc("/cvs/g/dm/{gid}", controllers.DeleteMember).Methods("DELETE")
	router.HandleFunc("/cvs/g/dg/{gid}", controllers.DeleteGroup).Methods("DELETE")

}
