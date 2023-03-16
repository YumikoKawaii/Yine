package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YumikoKawaii/Yine/pkg/models"
	"github.com/YumikoKawaii/Yine/pkg/security"
	"github.com/YumikoKawaii/Yine/pkg/utils"
	"github.com/gorilla/mux"
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
	Message  models.Message
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

	vars := mux.Vars(r)
	coid := vars["coid"]

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
		Message.NewMessage(id, coid, "text", message)
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
		Message.NewMessage(id, coid, "text", message)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func getAllMessageData(id string, coid string) []models.Message {

	if Group.IsGroup(coid) {

		return Message.FetchAllGroupConversation(coid)

	} else {

		return Message.FetchAllPersonalConversation(id, coid)

	}

}

func getMessageData(id string, coid string, mid uint32) []models.Message {

	if Group.IsGroup(coid) {

		return Message.FetchGroupConversation(coid, mid)

	} else {

		return Message.FetchPersonalConversation(id, coid, mid)

	}

}

func FetchMessage(w http.ResponseWriter, r *http.Request) {

	id := security.Authorize(w, r)
	if id == "" {
		return
	}

	vars := mux.Vars(r)
	coid := vars["coid"]

	if !Conversation.IsConversationExist(coid) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m_id := r.Form.Get("m_id")

	if m_id == "" {

		data := getAllMessageData(id, coid)
		res, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	} else {

		mid, err := strconv.Atoi(m_id)
		if err != nil || mid < 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data := getMessageData(id, coid, uint32(mid))
		res, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	}

}
