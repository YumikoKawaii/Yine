package routes

import (
	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/gorilla/mux"
)

var ConversationRoutes = func(router *mux.Router) {

	router.HandleFunc("/conversation/", controllers.ChangeNickname).Methods("PUT")

	router.HandleFunc("/conversation/c/", controllers.NewChat).Methods("POST")
	//router.HandleFunc("/conversation/c/", controllers.GetChatId).Methods("GET")
	//router.HandleFunc("/conversation/c/", controllers.DeleteChat).Methods("DELETE")

	router.HandleFunc("/conversation/g/", controllers.NewGroup).Methods("POST")
	router.HandleFunc("/conversation/g/am/", controllers.AddMemeber).Methods("POST")
	router.HandleFunc("/conversation/g/cmr/", controllers.ChangeMemberRole).Methods("PUT")
	router.HandleFunc("/conversation/g/cn/", controllers.ChangeGroupName).Methods("PUT")
	router.HandleFunc("/conversation/g/ca/", controllers.ChangeGroupAvatar).Methods("PUT")
	router.HandleFunc("/conversation/g/dm/", controllers.DeleteMember).Methods("DELETE")
	router.HandleFunc("/conversation/g/dg/", controllers.DeleteGroup).Methods("DELETE")

}
