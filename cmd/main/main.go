package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterAccountRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Println("Listening on 9010!")
}
