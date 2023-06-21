package client

import (
	"context"
	"encoding/json"

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

	userTokens := model.UserTokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	tokensValue, err := json.Marshal(userTokens)
	if err != nil {
		logger.Errorf("marshal tokens: %s", err.Error())
		return err
	}

	if err = s.redisClient.Set(ctx, info.Username, string(tokensValue), 0); err != nil {
		logger.Errorf("set tokens: %s", err.Error())
		return err
	}

	if err = s.redisClient.Set(ctx, model.LoggedUsername, info.Username, 0); err != nil {
		logger.Errorf("set logged username: %s", err.Error())
		return err
	}

	return nil
}
