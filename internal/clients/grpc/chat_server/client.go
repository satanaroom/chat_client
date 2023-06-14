package chat_server

import (
	"context"
	"fmt"

	converter "github.com/satanaroom/chat_client/internal/converter/chat_server"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

var _ Client = (*client)(nil)

type Client interface {
	CreateChat(ctx context.Context, usernames []string) (int64, error)
}

type client struct {
	chatClient chatV1.ChatV1Client
}

func NewClient(cl chatV1.ChatV1Client) *client {
	return &client{
		chatClient: cl,
	}
}

func (c *client) CreateChat(ctx context.Context, usernames []string) (int64, error) {
	resp, err := c.chatClient.CreateChat(ctx, converter.ToCreateChatRequest(usernames))
	if err != nil {
		return 0, fmt.Errorf("chatClient.CreateChat: %w", err)
	}

	return resp.GetChatId(), nil
}
