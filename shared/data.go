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

func (c *ConnectedClients) IsConnected(id int) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	for _, client := range c.clients {
		if client.id == id {
			return true
		}
	}
	return false
}

func (c *ConnectedClients) AddClient(id int, apiKey string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clients = append(c.clients, Client{id: id, apiKey: apiKey})
}

func (c *ConnectedClients) DisconnectClient(id int) {
	removeItem := func(array []Client, elementIdx int) []Client {
		return append(array[:elementIdx], array[elementIdx+1:]...)
	}

	findIndex := func(array []Client, element int) int {
		for idx, client := range array {
			if client.id == element {
				return idx
			}
		}
		return -1
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	desiredIndex := findIndex(c.clients, id)
	if desiredIndex < 0 {
		return
	}

	c.clients = removeItem(c.clients, desiredIndex)

}
