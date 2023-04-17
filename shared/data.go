package shared

import (
	"fmt"
	"sync"
)

// This generic function loops through the slice and compares each element in it with the element passed to the second parameter.
// Removes the first occurrence found and returns a new slice without the element.
func removeFromArray[T comparable](slice []T, element T) []T {
	sliceWithoutElement := make([]T, 0)
	alreadyFound := false
	for _, value := range slice {
		if value == element {
			if alreadyFound {
				sliceWithoutElement = append(sliceWithoutElement, value)
				continue
			}
			alreadyFound = true
		} else {
			sliceWithoutElement = append(sliceWithoutElement, value)
		}
	}
	return sliceWithoutElement
}

func (c *Client) display(key int) {
	fmt.Printf("Amount of currently active channels for client %d: %d\n\n", key, len(c.Channels))
}

type Event struct {
	Name string
	Data string
}

type Client struct {
	Channels []*chan Event
}

func NewConnectedClients() ConnectedClients {
	return ConnectedClients{
		clients: make(map[int]*Client),
	}
}

type ConnectedClients struct {
	clients map[int]*Client
	mutex   sync.RWMutex
}

func (cc *ConnectedClients) IsConnected(id int) bool {
	cc.mutex.RLock()
	defer cc.mutex.RUnlock()

	client, ok := cc.clients[id]

	if !ok {
		return false
	}

	return len(client.Channels) > 0
}

func (cc *ConnectedClients) ConnectClient(id int, channel *chan Event) {
	if cc.IsConnected(id) {
		cc.mutex.Lock()
		defer cc.mutex.Unlock()

		client := cc.clients[id]
		client.Channels = append(client.Channels, channel)

		// NOTE: Debug purposes
		fmt.Printf("\n\nAmount of different clients connected: %d\n", len(cc.clients))
		for key, value := range cc.clients {
			value.display(key)
		}
		// NOTE: Debug purposes

		return
	}

	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	client := Client{Channels: make([]*chan Event, 0)}
	client.Channels = append(client.Channels, channel)

	cc.clients[id] = &client

	// NOTE: Debug purposes
	fmt.Printf("\n\nAmount of different clients connected: %d\n", len(cc.clients))
	for key, value := range cc.clients {
		value.display(key)
	}
	// NOTE: Debug purposes
}

func (cc *ConnectedClients) DisconnectClient(id int, channel *chan Event) {
	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	client, ok := cc.clients[id]

	if ok {
		// for index, ch := range client.Channels {
		// 	if ch == channel {
		// 		client.Channels = append(client.Channels[:index], client.Channels[index+1:]...)
		// 		break
		// 	}
		// }

		client.Channels = removeFromArray(client.Channels, channel)

		if len(client.Channels) == 0 {
			delete(cc.clients, id)
		}

		// NOTE: Debug purposes
		fmt.Printf("\n\nAmount of different clients connected: %d\n", len(cc.clients))
		for key, value := range cc.clients {
			value.display(key)
		}
		// NOTE: Debug purposes
	}
}

func (cc *ConnectedClients) BroadCastEvent(id int, data Event) {
	if !cc.IsConnected(id) {
		return
	}

	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	client := cc.clients[id]

	for _, ch := range client.Channels {
		(*ch) <- data
	}
}
