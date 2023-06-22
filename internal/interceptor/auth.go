package interceptor

import (
	"context"

	"github.com/satanaroom/chat_client/internal/client/redis"
	"github.com/satanaroom/chat_client/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	redisClient redis.Client
}

func NewAuthInterceptor(redisClient redis.Client) *AuthInterceptor {
	return &AuthInterceptor{
		redisClient: redisClient,
	}
}

func (i *AuthInterceptor) Unary(ctx context.Context, method string, req interface{}, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	accessToken, err := i.redisClient.Get(ctx, model.AccessToken)
	if err != nil {
		return err
	}

	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
