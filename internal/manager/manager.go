package manager

import (
	"sync"

	"github.com/bytemeprod/websockets-go-chat/internal/types"
)

type Manager struct {
	clients map[types.Client]struct{}

	sync.Mutex
}

func (m *Manager) AddClient(client types.Client) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.clients[client] = struct{}{}
}

func (m *Manager) RemoveClient(client types.Client) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	delete(m.clients, client)
}
