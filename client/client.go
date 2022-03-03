package client

import "github.com/go-redis/redis/v7"

func New(url string) (redis.UniversalClient, error) {
	rurl, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{rurl.Addr},
		DB:       rurl.DB,
		Username: rurl.Username,
		Password: rurl.Password,
	})

	return c, nil
}
