package shared

import (
	"sync"
)

// This function can be used to create the new ConnectedClients instance
func NewConnectedClients() ConnectedClients {
	return ConnectedClients{
		clients: make(map[int]*Client),
	}
}

type Client struct {
	id     int
	apiKey string
	events chan string
}

type ConnectedClients struct {
	clients map[int]*Client
	mutex   sync.RWMutex
}

func (c *ConnectedClients) IsConnected(id int) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, ok := c.clients[id]
	return ok
}

func (c *ConnectedClients) ConnectClient(id int, apiKey string) {
	if c.IsConnected(id) {
		c.DisconnectClient(id)
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clients[id] = &Client{id: id, apiKey: apiKey, events: make(chan string, 10)}
}

func (c *ConnectedClients) DisconnectClient(id int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	client, ok := c.clients[id]
	if ok {
		close(client.events)
		delete(c.clients, id)
	}
}

// func (c *ConnectedClients) SendEvent(id int, data string) {
// 	if !c.IsConnected(id) {
// 		return
// 	}
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()
// 	client := c.clients[id]
// 	client.events <- data
// }

func (c *ConnectedClients) GetClientChannel(id int) chan string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	client, ok := c.clients[id]
	if !ok {
		return nil
	}
	return client.events
}
