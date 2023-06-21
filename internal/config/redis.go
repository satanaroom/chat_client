package config

import (
	"github.com/satanaroom/chat_client/pkg/env"
)

var _ RedisConfig = (*redisConfig)(nil)

const redisHostEnvName = "REDIS_HOST"

type RedisConfig interface {
	Host() string
}

type redisConfig struct {
	host string
}

func NewRedisConfig() (*redisConfig, error) {
	var host string
	env.ToString(&host, redisHostEnvName, "localhost:6369")

	return &redisConfig{
		host: host,
	}, nil
}

func (c *redisConfig) Host() string {
	return c.host
}
