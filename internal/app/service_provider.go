package app

import (
	"context"

	redisDb "github.com/redis/go-redis/v9"
	authV1 "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_client/internal/client/redis"
	authClient "github.com/satanaroom/chat_client/internal/clients/grpc/auth"
	chatClient "github.com/satanaroom/chat_client/internal/clients/grpc/chat_server"
	"github.com/satanaroom/chat_client/internal/closer"
	commandV1 "github.com/satanaroom/chat_client/internal/command/command_v1"
	"github.com/satanaroom/chat_client/internal/config"
	accessRepository "github.com/satanaroom/chat_client/internal/repository/access"
	clientService "github.com/satanaroom/chat_client/internal/service/client"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type serviceProvider struct {
	authConfig  config.AuthClientConfig
	chatConfig  config.ChatClientConfig
	tlsConfig   config.TLSConfig
	redisConfig config.RedisConfig

	authClient    authClient.Client
	chatClient    chatClient.Client
	clientService clientService.Service

	tlsCredentials credentials.TransportCredentials

	redisClient      redis.Client
	accessRepository accessRepository.Repository

	client *commandV1.ChatClient
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ClientService(ctx context.Context) clientService.Service {
	if s.clientService == nil {
		s.clientService = clientService.NewService(s.AuthClient(ctx), s.ChatClient(ctx), s.AccessRepository(ctx))
	}

	return s.clientService
}

func (s *serviceProvider) AuthClientConfig() config.AuthClientConfig {
	if s.authConfig == nil {
		cfg, err := config.NewAuthClientConfig()
		if err != nil {
			logger.Fatalf("failed to get access access config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) ChatClientConfig() config.ChatClientConfig {
	if s.chatConfig == nil {
		cfg, err := config.NewChatClientConfig()
		if err != nil {
			logger.Fatalf("failed to get chat access config: %s", err.Error())
		}

		s.chatConfig = cfg
	}

	return s.chatConfig
}

func (s *serviceProvider) AuthClient(ctx context.Context) authClient.Client {
	if s.authClient == nil {
		opts := grpc.WithTransportCredentials(s.TLSCredentials(ctx))

		conn, err := grpc.DialContext(ctx, s.AuthClientConfig().Host(), opts)
		if err != nil {
			logger.Fatalf("failed to connect %s: %s", s.AuthClientConfig().Host(), err.Error())
		}
		closer.Add(conn.Close)

		client := authV1.NewAuthV1Client(conn)
		s.authClient = authClient.NewClient(client)
	}

	return s.authClient
}

func (s *serviceProvider) ChatClient(ctx context.Context) chatClient.Client {
	if s.chatClient == nil {
		opts := grpc.WithTransportCredentials(s.TLSCredentials(ctx))

		conn, err := grpc.DialContext(ctx, s.ChatClientConfig().Host(), opts)
		if err != nil {
			logger.Fatalf("failed to connect %s: %s", s.ChatClientConfig().Host(), err.Error())
		}
		closer.Add(conn.Close)

		client := chatV1.NewChatV1Client(conn)
		s.chatClient = chatClient.NewClient(client)
	}

	return s.chatClient
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			logger.Fatalf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) TLSConfig() config.TLSConfig {
	if s.tlsConfig == nil {
		cfg, err := config.NewTLSConfig()
		if err != nil {
			logger.Fatalf("failed to get TLS config: %s", err.Error())
		}

		s.tlsConfig = cfg
	}

	return s.tlsConfig
}

func (s *serviceProvider) TLSCredentials(_ context.Context) credentials.TransportCredentials {
	if s.tlsCredentials == nil {
		creds, err := credentials.NewClientTLSFromFile(s.TLSConfig().CertFile(), "")
		if err != nil {
			logger.Fatalf("new access tls from file: %s", err.Error())
		}

		s.tlsCredentials = creds
	}

	return s.tlsCredentials
}

func (s *serviceProvider) RedisClient(ctx context.Context) redis.Client {
	if s.redisClient == nil {
		opts := redisDb.Options{
			Addr: s.RedisConfig().Host(),
		}

		client, err := redis.NewClient(ctx, &opts)
		if err != nil {
			logger.Fatalf("failed to initialize redis access: %s", err.Error())
		}
		if _, err := client.Redis().Ping(ctx); err != nil {
			logger.Fatalf("failed to ping redis: %s", err.Error())

		}

		closer.Add(client.Close)
		s.redisClient = client
	}
	return s.redisClient
}

func (s *serviceProvider) AccessRepository(ctx context.Context) accessRepository.Repository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.RedisClient(ctx))
	}

	return s.accessRepository
}
