package main

import (
	"net/http"

	"github.com/seipan/mylb/serverpool"
)

type LBHandler interface {
	Serve(http.ResponseWriter, *http.Request)
}

type lbHandler struct {
	serverPool serverpool.ServerPool
}

func (lb *lbHandler) Serve(w http.ResponseWriter, r *http.Request) {
	peer := lb.serverPool.GetNextValidPeer()
	if peer != nil {
		peer.Serve(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}

func NewLBHandler(serverPool serverpool.ServerPool) LBHandler {
	return &lbHandler{serverPool}
}
