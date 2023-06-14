package access

import (
	"context"

	"github.com/satanaroom/chat_client/internal/client/redis"
	"github.com/satanaroom/chat_client/internal/model"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	SaveAccessToken(ctx context.Context, tokenInfo *model.TokenInfo) error
	GetAccessToken(ctx context.Context, username string) (string, error)
	SetLoggedUsername(ctx context.Context, username string) error
	GetLoggedUsername(ctx context.Context) (string, error)
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

func (r *repository) GetAccessToken(ctx context.Context, username string) (string, error) {
	return r.redisClient.Redis().GetToken(ctx, username)
}

func (r *repository) SetLoggedUsername(ctx context.Context, username string) error {
	return r.redisClient.Redis().SetLoggedUsername(ctx, username)
}

func (r *repository) GetLoggedUsername(ctx context.Context) (string, error) {
	return r.redisClient.Redis().GetLoggedUsername(ctx)
}
