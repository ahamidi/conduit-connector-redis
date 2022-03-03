package destination

import (
	"context"
	"github.com/ahamidi/conduit-connector-redis/config"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"github.com/go-redis/redis/v7"
	"testing"
	"time"
)

const (
	LocalTestURL = "redis://localhost:6379"
	TestKey      = "conduit-test"
	TestValue    = "redis is the best"
)

func TestDestination_Lifecycle(t *testing.T) {
	cfg := map[string]string{
		config.URL: LocalTestURL,
	}

	ctx := context.Background()
	dest := &Destination{}
	err := dest.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = dest.Open(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	r := sdk.Record{
		Position:  nil,
		Metadata:  nil,
		CreatedAt: time.Time{},
		Key:       sdk.RawData(TestKey),
		Payload:   sdk.RawData(TestValue),
	}

	err = dest.Write(ctx, r)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	val, err := readTestKey()
	if err != nil {
		t.Fatalf("expected a no error, got: %v", err)
	}

	if val != TestValue {
		t.Fatalf("expected value %s, got: %s", TestValue, val)
	}

	err = dest.Teardown(ctx)
	if err != nil {
		t.Fatalf("expected a no error, got: %v", err)
	}
}

func readTestKey() (string, error) {
	rurl, err := redis.ParseURL(LocalTestURL)
	if err != nil {
		return "", err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: rurl.Addr,
	})

	res := rdb.Get(TestKey)
	if res.Err() != nil {
		return "", err
	}
	return res.Val(), err
}
