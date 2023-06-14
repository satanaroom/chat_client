package redis

import (
	"context"

	redisDb "github.com/redis/go-redis/v9"
	"github.com/satanaroom/chat_client/internal/model"
)

type Tokener interface {
	SetToken(ctx context.Context, tokenInfo *model.TokenInfo) error
	GetToken(ctx context.Context, key string) (string, error)
}

type Register interface {
	SetLoggedUsername(ctx context.Context, username string) error
	GetLoggedUsername(ctx context.Context) (string, error)
}

type Pinger interface {
	Ping(ctx context.Context) (string, error)
}

type Redis interface {
	Tokener
	Register
	Pinger
	Close() error
}

type redis struct {
	rdb *redisDb.Client
}

func (r *redis) SetToken(ctx context.Context, tokenInfo *model.TokenInfo) error {
	status := r.rdb.Set(ctx, tokenInfo.Username, tokenInfo.Token, tokenInfo.Expiration)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (r *redis) GetToken(ctx context.Context, key string) (string, error) {
	status := r.rdb.Get(ctx, key)
	return status.Result()
}

func (r *redis) SetLoggedUsername(ctx context.Context, username string) error {
	status := r.rdb.Set(ctx, model.RegisteredUsername, username, 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
func (r *redis) GetLoggedUsername(ctx context.Context) (string, error) {
	status := r.rdb.Get(ctx, model.RegisteredUsername)
	return status.Result()
}

func (r *redis) Close() error {
	return r.rdb.Close()
}

func (r *redis) Ping(ctx context.Context) (string, error) {
	return r.rdb.Ping(ctx).Result()
}
