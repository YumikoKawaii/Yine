package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/controllers"
	"github.com/YumikoKawaii/Yine/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.AccountRoutes(r)
	routes.ProfileRoutes(r)
	routes.RelationshipRoutes(r)
	routes.ConversationRoutes(r)
	routes.MessageRoutes(r)
	controllers.Init()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Println("Listening on 9010!")
}
