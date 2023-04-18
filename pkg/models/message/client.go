package message

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Scan       bool
	connection *websocket.Conn
	manager    *Manager
	egress     chan []byte
}

var (
	pongWait = 10 * time.Second

	pingInterval = (pongWait * 9) / 10
)

func NewClient(id string, conn *websocket.Conn, manager *Manager) *Client {

	return &Client{
		ID:         id,
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte),
	}

}

func (c *Client) Disconnect() {

	c.connection.Close()

}

func (c *Client) ReadMessage() {

	defer func() {
		c.manager.RemoveClient(c)
	}()

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return
	}

	c.connection.SetPongHandler(c.pongHandler)

	for {

		_, payLoad, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(err)
			}
			return
		}

		c.egress <- payLoad

	}

}

func (c *Client) WriteMessage() {

	defer func() {
		c.manager.RemoveClient(c)
	}()

	ticker := time.NewTicker(pingInterval)

	for {

		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					fmt.Println(err)
				}
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println(err)

			}
		case <-ticker.C:

			if err := c.connection.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				return
			}

		}

	}

}

func (c *Client) pongHandler(pongMsg string) error {

	return c.connection.SetReadDeadline(time.Now().Add(pongWait))

}

func (c *Client) GotMessage(message []byte) {

	c.egress <- message
}
