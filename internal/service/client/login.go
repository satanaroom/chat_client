package client

import (
	"context"

	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_client/internal/model"
)

func (s *service) Login(ctx context.Context, info *model.UserInfo) error {
	refreshToken, err := s.authClient.GetRefreshToken(ctx, info.Username, info.Password)
	if err != nil {
		logger.Errorf("authClient.GetRefreshToken: %s", err.Error())
		return err
	}

	accessToken, err := s.authClient.GetAccessToken(ctx, refreshToken)
	if err != nil {
		logger.Errorf("authClient.GetAccessToken: %s", err.Error())
		return err
	}

	if err = s.redisClient.Set(ctx, model.RefreshToken, refreshToken, 0); err != nil {
		logger.Errorf("set refresh token: %s", err.Error())
		return err
	}

	if err = s.redisClient.Set(ctx, model.AccessToken, accessToken, 0); err != nil {
		logger.Errorf("set refresh token: %s", err.Error())
		return err
	}

	if err = s.redisClient.Set(ctx, model.LoggedUsername, info.Username, 0); err != nil {
		logger.Errorf("set logged username: %s", err.Error())
		return err
	}

	return nil
}
