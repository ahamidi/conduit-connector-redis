package main

import (
	redis "github.com/ahamidi/conduit-connector-redis"
	"github.com/ahamidi/conduit-connector-redis/destination"
	"github.com/ahamidi/conduit-connector-redis/source"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(redis.Specification, source.NewSource, destination.NewDestination)
}
