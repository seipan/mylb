package backend

import (
	"net/http"
	"sync"
)

type Backend interface {
	SetDead(bool)
	GetIsDead() bool
	GetConnections() int
	GetURL() string
	Serve(http.ResponseWriter, *http.Request)
}

type backend struct {
	URL         string `json:"url"`
	IsDead      bool
	mu          sync.RWMutex
	connections int
}

func (backend *backend) SetDead(b bool) {
	backend.mu.Lock()
	backend.IsDead = b
	backend.mu.Unlock()
}

func (backend *backend) GetIsDead() bool {
	backend.mu.RLock()
	isAlive := backend.IsDead
	backend.mu.RUnlock()
	return isAlive
}

func (backend *backend) GetConnections() int {
	return backend.connections
}

func (backend *backend) GetURL() string {
	return backend.URL
}

func NewBackend(url string) backend {
	return backend{
		URL:    url,
		IsDead: false,
	}
}

func NewDefaultBackend() []backend {
	return []backend{
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
