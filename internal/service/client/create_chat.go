package client

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) CreateChat(ctx context.Context, username string) error {
	token, err := s.accessRepository.GetAccessToken(ctx, username)
	if err != nil {
		return fmt.Errorf("accessRepository.GetAccessToken: %w", err)
	}

	logger.Infof("user %s access token: %s", username, token)
	return nil
}
