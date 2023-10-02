package backend

import "sync"

type Backend struct {
	URL         string `json:"url"`
	IsDead      bool
	mu          sync.RWMutex
	connections int
}

func (backend *Backend) SetDead(b bool) {
	backend.mu.Lock()
	backend.IsDead = b
	backend.mu.Unlock()
}

func (backend *Backend) GetIsDead() bool {
	backend.mu.RLock()
	isAlive := backend.IsDead
	backend.mu.RUnlock()
	return isAlive
}

func (backend *Backend) GetConnections() int {
	return backend.connections
}
