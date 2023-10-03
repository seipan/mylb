package main

import (
	"context"
	"log"
	"net/http"
)

// Serve serves a loadbalancer.
func main() {
	go healthCheck(context.Background(), nil)

	s := http.Server{
		Addr:    ":" + "8080",
		Handler: http.HandlerFunc(nil),
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
