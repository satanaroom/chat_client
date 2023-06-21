package interceptor

import (
	"context"
	"encoding/json"

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
	loggedUsername, err := i.redisClient.Get(ctx, model.LoggedUsername)
	if err != nil {
		return err
	}

	tokens, err := i.redisClient.Get(ctx, loggedUsername)
	if err != nil {
		return err
	}

	var userTokens model.UserTokens
	if err = json.Unmarshal([]byte(tokens), &userTokens); err != nil {
		return err
	}

	md := metadata.New(map[string]string{"Authorization": "Bearer " + userTokens.AccessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
