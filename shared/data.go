package shared

import (
	"sync"
)

type Event struct {
	Name string
	Data string
}

type Client struct {
	Channels []*chan Event
}

// This function can be used to craft the new ConnectedClients instances
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

		return
	}

	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	client := Client{Channels: make([]*chan Event, 0)}
	client.Channels = append(client.Channels, channel)

	cc.clients[id] = &client
}

func (cc *ConnectedClients) DisconnectClient(id int, channel *chan Event) {
	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	client, ok := cc.clients[id]

	if ok {
		for index, ch := range client.Channels {
			if ch == channel {
				client.Channels = append(client.Channels[:index], client.Channels[index+1:]...)
				break
			}
		}

		if len(client.Channels) == 0 {
			delete(cc.clients, id)
		}
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
