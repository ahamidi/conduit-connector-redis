package destination

import (
	"context"

	"github.com/go-redis/redis/v7"

	"github.com/ahamidi/conduit-connector-redis/client"
	"github.com/ahamidi/conduit-connector-redis/config"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

// Destination client
type Destination struct {
	sdk.UnimplementedDestination
	config config.Config
	client redis.UniversalClient
}

func NewDestination() sdk.Destination {
	return &Destination{}
}

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {
	config2, err := config.Parse(cfg)
	if err != nil {
		return err
	}

	d.config = config2

	return nil
}

func (d *Destination) Open(ctx context.Context) error {
	c, err := client.New(d.config.URL)
	if err != nil {
		return err
	}
	d.client = c
	return nil
}

func (d *Destination) Write(ctx context.Context, r sdk.Record) error {
	key := string(r.Key.Bytes())
	status := d.client.Set(key, r.Payload.Bytes(), 0)
	return status.Err()
}

func (d *Destination) Teardown(ctx context.Context) error {
	if d.client != nil {
		return d.client.Close()
	}
	return nil
}
