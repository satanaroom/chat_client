package auth

import (
	"context"

	authV1 "github.com/satanaroom/auth/pkg/auth_v1"
	converter "github.com/satanaroom/chat_client/internal/converter/auth"
)

var _ Client = (*client)(nil)

type Client interface {
	GetRefreshToken(ctx context.Context, username, password string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	UpdateRefreshToken(ctx context.Context, oldToken string) (string, error)
}

type client struct {
	authClient authV1.AuthV1Client
}

func NewClient(cl authV1.AuthV1Client) *client {
	return &client{
		authClient: cl,
	}
}

func (c *client) GetRefreshToken(ctx context.Context, username, password string) (string, error) {
	resp, err := c.authClient.GetRefreshToken(ctx, converter.ToRefreshRequest(username, password))
	if err != nil {
		return "", err
	}

	return resp.GetRefreshToken(), nil
}

func (c *client) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	resp, err := c.authClient.GetAccessToken(ctx, converter.ToAccessRequest(refreshToken))
	if err != nil {
		return "", err
	}

	return resp.GetAccessToken(), nil
}

func (c *client) UpdateRefreshToken(ctx context.Context, oldToken string) (string, error) {
	resp, err := c.authClient.UpdateRefreshToken(ctx, converter.ToUpdateRefreshRequest(oldToken))
	if err != nil {
		return "", err
	}

	return resp.GetNewToken(), nil
}
