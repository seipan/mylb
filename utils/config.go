package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/seipan/mylb/backend"
	"github.com/seipan/mylb/lc"
	"github.com/seipan/mylb/lr"
	"github.com/seipan/mylb/serverpool"
)

type Config struct {
	Type  string   `json:"type"`
	Ports []string `json:"ports"`
}

func GetConfig() (Config, error) {
	var cfg Config
	data, err := os.ReadFile("./config.json")
	if err != nil {
		return cfg, err
	}
	json.Unmarshal(data, &cfg)
	return cfg, nil
}

func GetPoolType(backends []backend.Backend) (serverpool.ServerPool, error) {
	cfg, err := GetConfig()
	if err != nil {
		return nil, err
	}
	switch cfg.Type {
	case "lc":
		return lc.NewlcserverPool(backends), nil
	case "lr":
		return lr.NewlrserverPool(backends), nil
	default:
		return nil, errors.New("invalid server pool type")
	}

}
