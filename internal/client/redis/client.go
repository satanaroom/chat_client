package redis

import (
	"context"
	"time"

	redisDb "github.com/redis/go-redis/v9"
)

type Client interface {
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Ping(ctx context.Context) (string, error)
	Close() error
}

type client struct {
	rdb *redisDb.Client
}

func NewClient(_ context.Context, opts *redisDb.Options) (*client, error) {
	return &client{
		rdb: redisDb.NewClient(opts),
	}, nil
}

func (r *client) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	status := r.rdb.Set(ctx, key, value, expiration)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (r *client) Get(ctx context.Context, key string) (string, error) {
	status := r.rdb.Get(ctx, key)
	return status.Result()
}

func (r *client) Close() error {
	return r.rdb.Close()
}

func (r *client) Ping(ctx context.Context) (string, error) {
	return r.rdb.Ping(ctx).Result()
}
