package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers/role"
	"github.com/gorilla/mux"
)

var ConversationRoutes = func(router *mux.Router) {

	router.HandleFunc("/cvs", role.FetchAllConversation).Methods("GET")
	router.HandleFunc("/cvs/p/d/{coid}", role.FetchPersonalChat).Methods("GET")
	router.HandleFunc("/cvs/g/d/{coid}", role.FetchGroupChat).Methods("GET")
	router.HandleFunc("/cvs/{coid}", role.FetchBasicInfoConversation).Methods("GET")
	router.HandleFunc("/cvs/{coid}", role.ChangeNickname).Methods("PUT")

	router.HandleFunc("/cvs/g", role.NewGroup).Methods("POST")
	router.HandleFunc("/cvs/g/am/{gid}", role.AddMemeber).Methods("POST")
	router.HandleFunc("/cvs/g/cmr/{gid}", role.ChangeMemberRole).Methods("PUT")
	router.HandleFunc("/cvs/g/cn/{gid}", role.ChangeGroupName).Methods("PUT")
	router.HandleFunc("/cvs/g/ca/{gid}", role.ChangeGroupAvatar).Methods("PUT")
	router.HandleFunc("/cvs/g/dm/{gid}", role.DeleteMember).Methods("DELETE")
	router.HandleFunc("/cvs/g/dg/{gid}", role.DeleteGroup).Methods("DELETE")

}
