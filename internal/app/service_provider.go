package app

import (
	"context"

	redisDb "github.com/redis/go-redis/v9"
	authV1 "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/logger"
	authClient "github.com/satanaroom/chat_client/internal/client/grpc/auth"
	chatClient "github.com/satanaroom/chat_client/internal/client/grpc/chat_server"
	"github.com/satanaroom/chat_client/internal/client/redis"
	"github.com/satanaroom/chat_client/internal/closer"
	"github.com/satanaroom/chat_client/internal/config"
	"github.com/satanaroom/chat_client/internal/interceptor"
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

	tlsAuthCredentials       credentials.TransportCredentials
	tlsChatServerCredentials credentials.TransportCredentials

	redisClient redis.Client
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ClientService(ctx context.Context) clientService.Service {
	if s.clientService == nil {
		s.clientService = clientService.NewService(s.AuthClient(ctx), s.ChatClient(ctx), s.RedisClient(ctx), s.AuthClientConfig())
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
		opts := grpc.WithTransportCredentials(s.TLSAuthCredentials(ctx))

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
		authInterceptor := interceptor.NewAuthInterceptor(s.RedisClient(ctx))

		conn, err := grpc.DialContext(
			ctx,
			s.ChatClientConfig().Host(),
			grpc.WithUnaryInterceptor(authInterceptor.Unary),
			grpc.WithTransportCredentials(s.TLSChatServerCredentials(ctx)))
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

func (s *serviceProvider) TLSAuthCredentials(_ context.Context) credentials.TransportCredentials {
	if s.tlsAuthCredentials == nil {
		creds, err := credentials.NewClientTLSFromFile(s.TLSConfig().AuthCertFile(), "")
		if err != nil {
			logger.Fatalf("new access tls from file: %s", err.Error())
		}

		s.tlsAuthCredentials = creds
	}

	return s.tlsAuthCredentials
}

func (s *serviceProvider) TLSChatServerCredentials(_ context.Context) credentials.TransportCredentials {
	if s.tlsChatServerCredentials == nil {
		creds, err := credentials.NewClientTLSFromFile(s.TLSConfig().ChatServerCertFile(), "")
		if err != nil {
			logger.Fatalf("new access tls from file: %s", err.Error())
		}

		s.tlsChatServerCredentials = creds
	}

	return s.tlsChatServerCredentials
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

		if _, err = client.Ping(ctx); err != nil {
			logger.Fatalf("failed to ping redis: %s", err.Error())

		}
		closer.Add(client.Close)

		s.redisClient = client
	}
	return s.redisClient
}
