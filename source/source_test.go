package source

import (
	"context"
	"github.com/ahamidi/conduit-connector-redis/config"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"github.com/go-redis/redis/v7"
	"testing"
)

const LocalTestURL = "redis://localhost:6379"

func TestSource_Lifecycle(t *testing.T) {

	// Seed local redis
	err := seedRedisList(LocalTestURL)
	if err != nil {
		t.Fatalf("unable to seed test Redis database; error: %s", err.Error())
	}

	cfg := map[string]string{
		config.URL: LocalTestURL,
		Mode:       "list",
		Key:        "demo",
	}

	ctx := context.Background()
	source := &Source{}
	err = source.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = source.Open(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, err := source.Read(ctx)
	if err != nil && err.Error() != sdk.ErrBackoffRetry.Error() {
		t.Fatalf("expected a BackoffRetry error, got: %v", err)
	}

	if r.Payload == nil {
		t.Fatalf("expected payload, but got: %+v", r.Payload)

	}

	err = source.Teardown(ctx)
	if err != nil {
		t.Fatalf("expected a no error, got: %v", err)
	}
}

func seedRedisList(url string) error {
	rurl, err := redis.ParseURL(LocalTestURL)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: rurl.Addr,
	})

	for i := 0; i < 10; i++ {
		err := rdb.LPush("demo", i, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
