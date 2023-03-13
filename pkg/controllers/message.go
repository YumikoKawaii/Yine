package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}

	Client   models.Client
	manager  models.Manager
	Response models.Response
)

func checkOrigin(r *http.Request) bool {

	//return r.Header.Get("Origin") == "http://localhost:9010"
	return true

}

func Init() {

	manager = *models.NewManager()

}

func ConnectToSocket(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client := models.NewClient(id, conn, &manager)
	client.Scan = true

	manager.AddClient(client)

	go client.ReadMessage()
	go client.WriteMessage()

}

func SendMessage(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	coid := r.URL.Query().Get("coid")
	message := r.Form.Get("message")
	res, _ := json.Marshal(Response.NewResponse("message", id, "send", coid, message))

	if Account.IsIdExist(coid) {

		if !security.IsAccessable(id, coid) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if !Conversation.IsConversationBetween(id, coid) {
			Conversation.NewConnect(id, coid, utils.Member)
			Conversation.NewConnect(coid, id, utils.Member)
		}

		found, client := manager.GetClient(coid)
		if found {
			client.GotMessage(res)
		}
		w.WriteHeader(http.StatusOK)
		return
	} else if Group.IsGroup(coid) {

		if !Conversation.IsMember(coid, id) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		receivers := Conversation.GetReceivers(coid, id)
		for _, s := range receivers {

			found, client := manager.GetClient(s)
			if found {
				client.GotMessage(res)
			}

		}
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
