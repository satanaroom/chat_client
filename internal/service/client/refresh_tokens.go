package client

import (
	"context"
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
				refreshToken, err := s.redisClient.Get(ctx, model.RefreshToken)
				if err != nil {
					logger.Errorf("failed to get refresh token: %s", err.Error())
					continue
				}

				newRefreshToken, err := s.authClient.UpdateRefreshToken(ctx, refreshToken)
				if err != nil {
					logger.Errorf("failed to get new refresh token: %s", err.Error())
					continue
				}

				if err = s.redisClient.Set(ctx, model.RefreshToken, newRefreshToken, 0); err != nil {
					logger.Errorf("failed to set new tokens: %s", err.Error())
					continue
				}

				logger.Info("refresh token has been updated")
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
				refreshToken, err := s.redisClient.Get(ctx, model.RefreshToken)
				if err != nil {
					logger.Errorf("failed to get refresh token: %s", err.Error())
					continue
				}

				newAccessToken, err := s.authClient.GetAccessToken(ctx, refreshToken)
				if err != nil {
					logger.Errorf("failed to get new access token: %s", err.Error())
					continue
				}

				if err = s.redisClient.Set(ctx, model.AccessToken, newAccessToken, 0); err != nil {
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
