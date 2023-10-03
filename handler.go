package main

import "net/http"

type LBHandler interface {
	Serve(http.ResponseWriter, *http.Request)
}

type lbHandler struct {
	serverPool ServerPool
}

func (lb *lbHandler) Serve(w http.ResponseWriter, r *http.Request) {
	peer := lb.serverPool.GetNextValidPeer()
	if peer != nil {
		peer.Serve(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
