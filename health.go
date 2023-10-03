package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"time"
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

func healthCheck(ctx context.Context, sv ServerPool) {
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
				msg := "ok"
				if !isAlive {
					msg = "dead"
				}
				log.Printf("%v checked %v by healthcheck", b.GetURL(), msg)
			}
		case <-ctx.Done():
			return
		}

	}

}
