package config

import (
	"fmt"
)

const (
	URL = "url"
)

// Config represents configuration needed for S3
type Config struct {
	URL string
}

// Parse attempts to parse plugins.Config into a Config struct
func Parse(cfg map[string]string) (Config, error) {
	url, ok := cfg[URL]

	if !ok {
		return Config{}, requiredConfigErr(URL)
	}

	config := Config{
		URL: url,
	}

	return config, nil
}

func requiredConfigErr(name string) error {
	return fmt.Errorf("%q config value must be set", name)
}
