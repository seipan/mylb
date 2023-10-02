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

func NewBackend(url string) Backend {
	return Backend{
		URL:    url,
		IsDead: false,
	}
}

func NewDefaultBackend() []Backend {
	return []Backend{
		{
			URL:    "http://localhost:8081/",
			IsDead: false,
		},
		{
			URL:    "http://localhost:8082/",
			IsDead: false,
		},
		{
			URL:    "http://localhost:8083/",
			IsDead: false,
		},
		{
			URL:    "http://localhost:8086/",
			IsDead: false,
		},
		{
			URL:    "http://localhost:8087/",
			IsDead: false,
		},
		{
			URL:    "http://localhost:8088/",
			IsDead: false,
		},
		{
			URL:    "http://localhost:8089/",
			IsDead: false,
		},
	}
}
