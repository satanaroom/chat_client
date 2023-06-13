package chat_server

import (
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

var _ Client = (*client)(nil)

type Client interface {
}

type client struct {
	chatClient chatV1.ChatV1Client
}

func NewClient(cl chatV1.ChatV1Client) *client {
	return &client{
		chatClient: cl,
	}
}
