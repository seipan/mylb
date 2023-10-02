package lc

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/seipan/mylb/backend"
)

// lbHandler is a handler for loadbalancing
func LcHandler(w http.ResponseWriter, r *http.Request) {
	idx := 0
	bcds := backend.NewDefaultBackend()
	lc := NewlcserverPool(bcds)
	maxLen := len(bcds)
	// Round Robin
	lc.mu.Lock()
	currentBackend := lc.Backends[idx%maxLen]
	if currentBackend.GetIsDead() {
		idx++
	}
	targetURL, err := url.Parse(lc.Backends[idx%maxLen].URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	idx++
	lc.mu.Unlock()
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		// NOTE: It is better to implement retry.
		log.Printf("%v is dead.", targetURL)
		currentBackend.SetDead(true)
		LcHandler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}
