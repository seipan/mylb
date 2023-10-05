package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"time"

	"github.com/seipan/mylb/serverpool"
)

func isAlive(ctx context.Context, url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute*1)
	if err != nil {
		log.Printf("Unreachable to %v, error %v:", url.Host, err)
		return false
	}
	defer conn.Close()
	return true
}

func healthCheck(ctx context.Context, sv serverpool.ServerPool) {
	t := time.NewTicker(time.Minute * 1)
	for {
		select {
		case <-t.C:
			bcd := sv.GetBackends()
			for _, b := range bcd {
				b := b
				pingURL, err := url.Parse(b.GetURL())
				if err != nil {
					log.Fatal(err.Error())
				}
				isAlive := isAlive(ctx, pingURL)
				b.SetDead(!isAlive)
			}
		case <-ctx.Done():
			return
		}

	}

}
