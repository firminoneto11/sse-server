package shared

import "sync"

type Client struct {
	id     int
	apiKey string
}

type ConnectedClients struct {
	clients []Client
	mutex   sync.RWMutex
}

func (c *ConnectedClients) GetConnectedClients() []Client {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.clients
}

func (c *ConnectedClients) AddClient(id int, apiKey string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clients = append(c.clients, Client{id: id, apiKey: apiKey})
}
