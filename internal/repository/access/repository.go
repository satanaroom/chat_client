package access

import (
	"context"

	"github.com/satanaroom/chat_client/internal/client/redis"
	"github.com/satanaroom/chat_client/internal/model"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	SaveAccessToken(ctx context.Context, tokenInfo *model.TokenInfo) error
	GetAccessToken(ctx context.Context, key string) (string, error)
}

type repository struct {
	redisClient redis.Client
}

func NewRepository(redisClient redis.Client) *repository {
	return &repository{
		redisClient: redisClient,
	}
}

func (r *repository) SaveAccessToken(ctx context.Context, tokenInfo *model.TokenInfo) error {
	return r.redisClient.Redis().SetToken(ctx, tokenInfo)
}
func (r *repository) GetAccessToken(ctx context.Context, key string) (string, error) {
	return r.redisClient.Redis().GetToken(ctx, key)
}
