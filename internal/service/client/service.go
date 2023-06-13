package client

import (
	"context"

	authClient "github.com/satanaroom/chat_client/internal/clients/grpc/auth"
	chatClient "github.com/satanaroom/chat_client/internal/clients/grpc/chat_server"
	"github.com/satanaroom/chat_client/internal/model"
	accessRepository "github.com/satanaroom/chat_client/internal/repository/access"
)

var _ Service = (*service)(nil)

type Service interface {
	Login(ctx context.Context, credentials *model.UserCredentials) error
	CreateChat(ctx context.Context, username string) error
	ConnectChat(ctx context.Context) error
}

type service struct {
	authClient       authClient.Client
	chatClient       chatClient.Client
	accessRepository accessRepository.Repository
}

func NewService(auth authClient.Client, chat chatClient.Client, accessRepository accessRepository.Repository) *service {
	return &service{
		authClient:       auth,
		chatClient:       chat,
		accessRepository: accessRepository,
	}
}
