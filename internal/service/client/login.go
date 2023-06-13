package client

import (
	"context"
	"fmt"
	"time"

	"github.com/satanaroom/chat_client/internal/model"
)

func (s *service) Login(ctx context.Context, credentials *model.UserCredentials) error {
	refreshToken, err := s.authClient.GetRefreshToken(ctx, credentials.Username, credentials.Password)
	if err != nil {
		return fmt.Errorf("authClient.GetRefreshToken: %w", err)
	}
	accessToken, err := s.authClient.GetAccessToken(ctx, refreshToken)
	if err != nil {
		return fmt.Errorf("authClient.GetRefreshToken: %w", err)
	}

	if err = s.accessRepository.SaveAccessToken(ctx, &model.TokenInfo{
		Username:   credentials.Username,
		Token:      accessToken,
		Expiration: 5 * time.Minute,
	}); err != nil {
		return fmt.Errorf("accessRepository.SaveAccessToken: %w", err)
	}
	return nil
}
