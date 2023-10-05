package lr

import (
	"sync"

	"github.com/seipan/mylb/backend"
	"github.com/seipan/mylb/serverpool"
)

type lrserverPool struct {
	Backends []backend.Backend `json:"backends"`
	mu       sync.RWMutex
}

func (s *lrserverPool) GetNextValidPeer() backend.Backend {
	var leastConnectedPeer backend.Backend
	for _, b := range s.Backends {
		if !b.GetIsDead() {
			leastConnectedPeer = b
			break
		}
	}

	for _, b := range s.Backends {
		if b.GetIsDead() {
			continue
		}
		if leastConnectedPeer.GetResponseTime() > b.GetResponseTime() {
			leastConnectedPeer = b
		}
	}
	return leastConnectedPeer
}

func (s *lrserverPool) GetBackends() []backend.Backend {
	return s.Backends
}

func (s *lrserverPool) AddBackend(b backend.Backend) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Backends = append(s.Backends, b)
}

func (s *lrserverPool) GetServerPoolSize() int {
	return len(s.Backends)
}

func NewlrserverPool(backends []backend.Backend) serverpool.ServerPool {
	return &lrserverPool{
		Backends: backends,
	}
}
