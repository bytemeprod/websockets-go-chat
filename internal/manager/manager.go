package manager

import (
	"sync"

	"github.com/bytemeprod/websockets-go-chat/internal/types"
)

type Manager struct {
	Clients map[types.Client]struct{}

	sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Clients: make(map[types.Client]struct{}),
	}
}

func (m *Manager) AddClient(client types.Client) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Clients[client] = struct{}{}
}

func (m *Manager) RemoveClient(client types.Client) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	delete(m.Clients, client)
}
