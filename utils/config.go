package utils

import (
	"encoding/json"
	"os"
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
