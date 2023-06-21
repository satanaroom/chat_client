package client

import (
	"context"
	"time"

	authClient "github.com/satanaroom/chat_client/internal/client/grpc/auth"
	chatClient "github.com/satanaroom/chat_client/internal/client/grpc/chat_server"
	redisClient "github.com/satanaroom/chat_client/internal/client/redis"
	"github.com/satanaroom/chat_client/internal/config"
	"github.com/satanaroom/chat_client/internal/model"
)

var _ Service = (*service)(nil)

type Service interface {
	Login(ctx context.Context, info *model.UserInfo) error
	CreateChat(ctx context.Context, usernames []string) (string, error)
	ConnectChat(ctx context.Context, chatId string) error
	RefreshTokens(ctx context.Context, refreshTokenPeriod, accessTokenPeriod time.Duration)
}

type service struct {
	authConfig config.AuthClientConfig

	authClient  authClient.Client
	chatClient  chatClient.Client
	redisClient redisClient.Client
}

func NewService(authClient authClient.Client, chatClient chatClient.Client,
	redisClient redisClient.Client, authConfig config.AuthClientConfig) *service {
	return &service{
		authConfig:  authConfig,
		authClient:  authClient,
		chatClient:  chatClient,
		redisClient: redisClient,
	}
}
