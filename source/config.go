package source

import (
	"fmt"
	"github.com/ahamidi/conduit-connector-redis/config"
)

const (
	// operating mode. Only `list` is currently supported
	Mode = "mode"
	Key  = "key"

	// operating modes
	ModeList = "list"
)

// Config represents source configuration with Redis configurations
type Config struct {
	config.Config
	Mode string
	Key  string
}

// Parse attempts to parse the configurations into a Config struct that Source could utilize
func Parse(cfg map[string]string) (Config, error) {
	common, err := config.Parse(cfg)
	if err != nil {
		return Config{}, err
	}

	if mode, exists := cfg[Mode]; !exists || mode == "" {
		return Config{}, fmt.Errorf("%q config value is required", Mode)
	}
	mode := cfg[Mode]

	var k string
	var exists bool
	if mode == ModeList {
		if k, exists = cfg[Key]; !exists {
			return Config{}, fmt.Errorf("%q config value is required", Key)
		}
	}

	sourceConfig := Config{
		Config: common,
		Mode:   mode,
		Key:    k,
	}

	return sourceConfig, nil
}
