package config

import (
	"github.com/satanaroom/chat_client/pkg/env"
)

var _ TLSConfig = (*tlsConfig)(nil)

const (
	tlsAuthCertFileEnvName       = "TLS_AUTH_CERT_FILE"
	tlsChatServerCertFileEnvName = "TLS_CHAT_SERVER_CERT_FILE"
)

type TLSConfig interface {
	AuthCertFile() string
	ChatServerCertFile() string
}

type tlsConfig struct {
	authCertFile       string
	chatServerCertFile string
}

func NewTLSConfig() (*tlsConfig, error) {
	var authCertFile, chatServerCertFile string

	env.ToString(&authCertFile, tlsAuthCertFileEnvName, "../auth/service.pem")
	env.ToString(&chatServerCertFile, tlsChatServerCertFileEnvName, "../chat_server/service.pem")

	return &tlsConfig{
		authCertFile:       authCertFile,
		chatServerCertFile: chatServerCertFile,
	}, nil
}

func (c *tlsConfig) AuthCertFile() string {
	return c.authCertFile
}

func (c *tlsConfig) ChatServerCertFile() string {
	return c.chatServerCertFile
}
