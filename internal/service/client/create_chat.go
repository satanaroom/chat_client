package client

import (
	"context"

	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) CreateChat(ctx context.Context, usernames []string) (string, error) {
	chatId, err := s.chatClient.CreateChat(ctx, usernames)
	if err != nil {
		logger.Errorf("chatClient.CreateChat: %s", err.Error())
		return "", err
	}

	return chatId, nil
}
