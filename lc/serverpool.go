package lc

import (
	"sync"

	"github.com/seipan/mylb/backend"
)

type lcserverPool struct {
	Backends []backend.Backend `json:"backends"`
	mu       sync.RWMutex
}

func (s *lcserverPool) GetNextValidPeer() backend.Backend {
	var leastConnectedPeer backend.Backend
	for _, b := range s.Backends {
		if !b.GetIsDead() {
			leastConnectedPeer = b
			break
		}
	}

	for _, b := range s.Backends {
		if !b.GetIsDead() {
			continue
		}
		if leastConnectedPeer.GetConnections() > b.GetConnections() {
			leastConnectedPeer = b
		}
	}
	return leastConnectedPeer
}

func (s *lcserverPool) GetBackends() []backend.Backend {
	return s.Backends
}

func (s *lcserverPool) AddBackend(b backend.Backend) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Backends = append(s.Backends, b)
}

func (s *lcserverPool) GetServerPoolSize() int {
	return len(s.Backends)
}

func NewlcserverPool(backends []backend.Backend) lcserverPool {
	return lcserverPool{
		Backends: backends,
	}
}
