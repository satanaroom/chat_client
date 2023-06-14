package client

import (
	"context"
	"time"

	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_client/internal/model"
)

func (s *service) Login(ctx context.Context, credentials *model.UserCredentials) error {
	refreshToken, err := s.authClient.GetRefreshToken(ctx, credentials.Username, credentials.Password)
	if err != nil {
		logger.Errorf("authClient.GetRefreshToken: %s", err.Error())
		return err
	}
	accessToken, err := s.authClient.GetAccessToken(ctx, refreshToken)
	if err != nil {
		logger.Errorf("authClient.GetAccessToken: %s", err.Error())
		return err
	}

	if err = s.accessRepository.SaveAccessToken(ctx, &model.TokenInfo{
		Username:   credentials.Username,
		Token:      accessToken,
		Expiration: 5 * time.Minute,
	}); err != nil {
		logger.Errorf("accessRepository.SaveAccessToken: %s", err.Error())
		return err
	}
	return nil
}

func (s *service) SetLoggedUsername(ctx context.Context, username string) error {
	if err := s.accessRepository.SetLoggedUsername(ctx, username); err != nil {
		logger.Errorf("accessRepository.SetLoggedUsername: %s", err.Error())
		return err
	}

	return nil
}

func (s *service) GetLoggedUsername(ctx context.Context) (string, error) {
	username, err := s.accessRepository.GetLoggedUsername(ctx)
	if err != nil {
		logger.Errorf("accessRepository.GetLoggedUsername: %s", err.Error())
		return "", err
	}

	return username, nil
}
