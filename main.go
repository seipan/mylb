package main

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/seipan/mylb/backend"
	"github.com/seipan/mylb/lc"
	"github.com/seipan/mylb/utils"
	"go.uber.org/zap"
)

// Serve serves a loadbalancer.
func main() {
	backends := backend.NewDefaultBackend()
	for _, b := range backends {
		url, err := url.Parse(b.GetURL())
		if err != nil {
			utils.Error("usrl parse err",
				zap.String("error", err.Error()),
				zap.String("url", b.GetURL()),
			)
		}
		b.SetReverProxy(httputil.NewSingleHostReverseProxy(url))
	}
	serverPool := lc.NewlcserverPool(backends)
	lbHandler := NewLBHandler(serverPool)

	go healthCheck(context.Background(), nil)

	s := http.Server{
		Addr:    ":" + "8080",
		Handler: http.HandlerFunc(lbHandler.Serve),
	}
	if err := s.ListenAndServe(); err != nil {
		utils.Error("listen and serve err",
			zap.String("error", err.Error()),
		)
	}
}
