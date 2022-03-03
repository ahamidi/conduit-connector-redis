package redis

import (
	"github.com/ahamidi/conduit-connector-redis/config"
	"github.com/ahamidi/conduit-connector-redis/source"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Spec struct{}

// Specification returns the Plugin's Specification.
func Specification() sdk.Specification {
	return sdk.Specification{
		Name:    "redis",
		Summary: "A Neo4j source and destination plugin for Conduit, written in Go.",
		Version: "v0.0.1",
		Author:  "Ali Hamidi",
		SourceParams: map[string]sdk.Parameter{
			config.URL: {
				Default:     "",
				Required:    true,
				Description: "The URL of the redis server.",
			},
			source.Mode: {
				Default:     "list",
				Required:    false,
				Description: "The operating mode to use. Valid options are: \"list\"",
			},
			source.Key: {
				Default:     "",
				Required:    true,
				Description: "The Key for the Redis list to read.",
			},
		},
	}
}
