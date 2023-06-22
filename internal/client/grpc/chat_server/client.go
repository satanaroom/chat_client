package chat_server

import (
	"context"

	converter "github.com/satanaroom/chat_client/internal/converter/chat_server"
	"github.com/satanaroom/chat_client/internal/model"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

var _ Client = (*client)(nil)

type Client interface {
	CreateChat(ctx context.Context, usernames []string) (string, error)
	ConnectChat(ctx context.Context, chatId string, username string) (chatV1.ChatV1_ConnectChatClient, error)
	SendMessage(ctx context.Context, chatId string, message *model.Message) error
}

type client struct {
	chatClient chatV1.ChatV1Client
}

func NewClient(cl chatV1.ChatV1Client) *client {
	return &client{
		chatClient: cl,
	}
}

func (c *client) CreateChat(ctx context.Context, usernames []string) (string, error) {
	resp, err := c.chatClient.CreateChat(ctx, converter.ToCreateChatRequest(usernames))
	if err != nil {
		return "", err
	}

	return resp.GetChatId(), nil
}

func (c *client) ConnectChat(ctx context.Context, chatId string, username string) (chatV1.ChatV1_ConnectChatClient, error) {
	resp, err := c.chatClient.ConnectChat(ctx, converter.ToConnectChatRequest(chatId, username))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) SendMessage(ctx context.Context, chatId string, message *model.Message) error {
	if _, err := c.chatClient.SendMessage(ctx, converter.ToSendMessageRequest(chatId, message)); err != nil {
		return err
	}

	return nil
}
