package config

import (
	"encoding/json"
	"os"
)

type VtxConfig struct {
	PluginId    string `json:"pluginId"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Country     string `json:"country"`
	HasFrontend bool   `json:"hasFrontend"`
	Backend     struct {
		Project string `json:"project"`
	} `json:"backend"`
	Frontend struct {
		Root string `json:"root"`
	} `json:"frontend"`
}

func LoadConfig(path string) (*VtxConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg VtxConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
