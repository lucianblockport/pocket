package config

import (
	"encoding/json"
	"os"
)

// GenesisConfig defines the structure of the genesis configuration.
// This is used to unmarshal the genesis.json file.
type GenesisConfig struct {
	GenesisTime string `json:"genesis_time"`
	ExtraData   string `json:"extra_data"`
}

// LoadGenesisConfig loads the genesis configuration from the given path.
func LoadGenesisConfig(path string) (*GenesisConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config GenesisConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
