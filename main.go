package main

import (
	"context"
	"log"
	"net/http"

	"github.com/seipan/mylb/lc"
)

// Serve serves a loadbalancer.
func main() {
	go healthCheck(context.Background(), nil)

	s := http.Server{
		Addr:    ":" + "8080",
		Handler: http.HandlerFunc(lc.LcHandler),
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
