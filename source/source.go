package source

import (
	"context"
	connector "github.com/ahamidi/conduit-connector-redis/client"
	"strconv"
	"time"

	sdk "github.com/conduitio/conduit-connector-sdk"
	redis "github.com/go-redis/redis/v7"
)

// Source connector
type Source struct {
	sdk.UnimplementedSource
	config Config
	client redis.UniversalClient
	pos    sdk.Position
}

func NewSource() sdk.Source {
	return &Source{}
}

// Configure parses and stores the configurations
// returns an error in case of invalid config
func (s *Source) Configure(ctx context.Context, cfg map[string]string) error {
	config2, err := Parse(cfg)
	if err != nil {
		return err
	}

	s.config = config2

	return nil
}

// Open prepare the plugin to start sending records from the given position
func (s *Source) Open(ctx context.Context, rp sdk.Position) error {
	c, err := connector.New(s.config.URL)
	if err != nil {
		return err
	}

	s.client = c
	if rp != nil {
		s.pos = rp
	}
	return nil
}

// Read gets the next object from Redis
func (s *Source) Read(ctx context.Context) (sdk.Record, error) {
	var i int
	var err error
	if s.pos != nil {
		i, err = strconv.Atoi(string(s.pos))
		if err != nil {
			return sdk.Record{}, err
		}

	}
	v, err := s.client.LIndex(s.config.Key, int64(i)).Result()
	if err != nil {
		return sdk.Record{}, err
	}
	pos := sdk.Position(strconv.Itoa(i + 1))
	r := sdk.Record{
		Position:  pos,
		CreatedAt: time.Now(),
		Key:       sdk.RawData(strconv.Itoa(i)),
		Payload:   sdk.RawData(v),
	}

	return r, err
}

// Teardown is called when the connector is stopped
func (s *Source) Teardown(ctx context.Context) error {
	if s.client != nil {
		return s.client.Close()
	}
	return nil
}

// Ack ...
func (s *Source) Ack(ctx context.Context, position sdk.Position) error {
	return nil // no ack needed
}
