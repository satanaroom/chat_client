package config

import (
	"github.com/satanaroom/chat_client/pkg/env"
)

var _ ChatClientConfig = (*chatClientConfig)(nil)

const chatHostEnvName = "CHAT_HOST"

type ChatClientConfig interface {
	Host() string
}

type chatClientConfig struct {
	host string
}

func NewChatClientConfig() (*chatClientConfig, error) {
	var host string
	env.ToString(&host, chatHostEnvName, "localhost:50052")

	return &chatClientConfig{
		host: host,
	}, nil
}

func (c *chatClientConfig) Host() string {
	return c.host
}
