package lc

import (
	"github.com/seipan/mylb/backend"
	"github.com/seipan/mylb/proxy"
)

type lcserverPool struct {
	Proxy    proxy.Proxy       `json:"proxy"`
	Backends []backend.Backend `json:"backends"`
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
