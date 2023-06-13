package redis

import (
	"context"

	redisDb "github.com/redis/go-redis/v9"
)

var _ Client = (*client)(nil)

type Client interface {
	Close() error
	Redis() Redis
}

type client struct {
	redis Redis
}

func NewClient(_ context.Context, opts *redisDb.Options) (*client, error) {
	rdb := redisDb.NewClient(opts)

	return &client{
		redis: &redis{
			rdb: rdb,
		},
	}, nil
}

func (c *client) Redis() Redis {
	return c.redis
}

func (c *client) Close() error {
	if c.redis != nil {
		return c.redis.Close()
	}
	return nil
}
