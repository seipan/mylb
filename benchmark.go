package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/seipan/mylb/serverpool"
)

func responseTime(path string) (time.Duration, error) {
	start := time.Now()
	_, err := http.Get(path)
	if err != nil {
		return 0, err
	}
	return time.Since(start), nil
}

func benchCheck(ctx context.Context, sv serverpool.ServerPool) {
	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			bcd := sv.GetBackends()
			for _, b := range bcd {
				b := b
				rest, err := responseTime(b.GetURL())
				if err != nil {
					log.Fatal(err.Error())
				}
				b.SetResponseTime(rest)
			}
		case <-ctx.Done():
			return
		}

	}

}
