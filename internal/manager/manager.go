package manager

import (
	"context"
	"sync"

	"github.com/bytemeprod/websockets-go-chat/internal/redisstore"
	"github.com/bytemeprod/websockets-go-chat/internal/types"
)

type Manager struct {
	Clients map[types.Client]struct{}
	storage *redisstore.RedisClient
	sync.Mutex
}

func NewManager(storage *redisstore.RedisClient) *Manager {
	return &Manager{
		Clients: make(map[types.Client]struct{}),
		storage: storage,
	}
}

func (m *Manager) AddClient(client types.Client) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Clients[client] = struct{}{}
}

func (m *Manager) RemoveClient(client types.Client) {
	m.Mutex.Lock()
	delete(m.Clients, client)
	m.Mutex.Unlock()

	m.storage.RemoveClient(context.Background(), client.GetUsername())
}
