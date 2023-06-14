package client

import (
	"context"

	"github.com/satanaroom/auth/pkg/logger"
	"google.golang.org/grpc/metadata"
)

func (s *service) CreateChat(ctx context.Context, loggedUsername string, usernames []string) (int64, error) {
	token, err := s.accessRepository.GetAccessToken(ctx, loggedUsername)
	if err != nil {
		logger.Errorf("accessRepository.GetAccessToken: %s", err.Error())
		return 0, err
	}

	md := metadata.New(map[string]string{"Authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	chatId, err := s.chatClient.CreateChat(ctx, usernames)
	if err != nil {
		logger.Errorf("chatClient.CreateChat: %s", err.Error())

		return 0, err
	}

	return chatId, nil
}
