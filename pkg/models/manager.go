package models

import (
	"sync"
)

type Manager struct {
	clients map[*Client]bool
	ids     map[string]*Client
	sync.RWMutex
}

func NewManager() *Manager {

	return &Manager{
		clients: make(map[*Client]bool),
		ids:     make(map[string]*Client),
	}

}

func (m *Manager) AddClient(client *Client) {

	m.Lock()
	defer m.Unlock()

	m.ids[client.ID] = client

	m.clients[client] = true

}

func (m *Manager) RemoveClient(client *Client) {

	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.Disconnect()
		delete(m.clients, client)
		delete(m.ids, client.ID)
	}

}

func (m *Manager) GetClient(id string) (bool, *Client) {

	if _, ok := m.ids[id]; ok {
		return true, m.ids[id]
	}

	return false, nil
}
