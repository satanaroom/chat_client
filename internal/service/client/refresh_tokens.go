package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_client/internal/model"
)

func (s *service) RefreshTokens(ctx context.Context, refreshTokenPeriod, accessTokenPeriod time.Duration) {
	go func() {
		t := time.NewTicker(refreshTokenPeriod)

		for {
			select {
			case <-t.C:
				loggedUsername, err := s.redisClient.Get(ctx, model.LoggedUsername)
				if err != nil {
					logger.Errorf("failed to get logged username: %s", err.Error())
					continue
				}

				tokens, err := s.redisClient.Get(ctx, loggedUsername)
				if err != nil {
					logger.Errorf("failed to get logged username tokens: %s", err.Error())
					continue
				}

				var userTokens model.UserTokens
				if err = json.Unmarshal([]byte(tokens), &userTokens); err != nil {
					logger.Errorf("unmarshal tokens: %s", err.Error())
					continue
				}

				newRefreshToken, err := s.authClient.UpdateRefreshToken(ctx, userTokens.RefreshToken)
				if err != nil {
					logger.Errorf("failed to get new refresh token: %s", err.Error())
					continue
				}

				userTokens.RefreshToken = newRefreshToken

				newTokens, err := json.Marshal(userTokens)
				if err != nil {
					logger.Errorf("marshal refreshed tokens: %s", err.Error())
					continue
				}

				if err = s.redisClient.Set(ctx, loggedUsername, string(newTokens), 0); err != nil {
					logger.Errorf("failed to set new tokens: %s", err.Error())
					continue
				}

				logger.Info("access token has been updated")
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		t := time.NewTicker(accessTokenPeriod)

		for {
			select {
			case <-t.C:
				loggedUsername, err := s.redisClient.Get(ctx, model.LoggedUsername)
				if err != nil {
					logger.Errorf("failed to get logged username: %s", err.Error())
					continue
				}

				tokens, err := s.redisClient.Get(ctx, loggedUsername)
				if err != nil {
					logger.Errorf("failed to get logged username tokens: %s", err.Error())
					continue
				}

				var userTokens model.UserTokens
				if err = json.Unmarshal([]byte(tokens), &userTokens); err != nil {
					logger.Errorf("unmarshal tokens: %s", err.Error())
					continue
				}

				newAccessToken, err := s.authClient.GetAccessToken(ctx, userTokens.RefreshToken)
				if err != nil {
					logger.Errorf("failed to get new access token: %s", err.Error())
					continue
				}

				userTokens.AccessToken = newAccessToken

				newTokens, err := json.Marshal(userTokens)
				if err != nil {
					logger.Errorf("marshal refreshed tokens: %s", err.Error())
					continue
				}

				if err = s.redisClient.Set(ctx, loggedUsername, string(newTokens), 0); err != nil {
					logger.Errorf("failed to set new tokens: %s", err.Error())
					continue
				}

				logger.Info("access token has been updated")
			case <-ctx.Done():
				return
			}
		}
	}()
}
