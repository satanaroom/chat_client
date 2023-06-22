package config

import (
	"time"

	"github.com/satanaroom/chat_client/pkg/env"
)

var _ AuthClientConfig = (*authClientConfig)(nil)

const (
	authHostEnvName = "AUTH_HOST"

	refreshTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION_MINUTES"
	accessTokenExpirationEnvName  = "ACCESS_TOKEN_EXPIRATION_MINUTES"
)

type AuthClientConfig interface {
	Host() string
	RefreshTokenExpiration() time.Duration
	AccessTokenExpiration() time.Duration
}

type authClientConfig struct {
	host                   string
	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration
}

func NewAuthClientConfig() (*authClientConfig, error) {
	var (
		host                   string
		refreshTokenExpiration int
		accessTokenExpiration  int
	)
	env.ToString(&host, authHostEnvName, "localhost:50051")
	env.ToInt(&refreshTokenExpiration, refreshTokenExpirationEnvName, 60)
	env.ToInt(&accessTokenExpiration, accessTokenExpirationEnvName, 5)

	return &authClientConfig{
		host:                   host,
		refreshTokenExpiration: time.Minute * time.Duration(refreshTokenExpiration),
		accessTokenExpiration:  time.Minute * time.Duration(accessTokenExpiration),
	}, nil
}

func (c *authClientConfig) Host() string {
	return c.host
}

func (c *authClientConfig) RefreshTokenExpiration() time.Duration {
	return c.refreshTokenExpiration
}
func (c *authClientConfig) AccessTokenExpiration() time.Duration {
	return c.accessTokenExpiration
}
