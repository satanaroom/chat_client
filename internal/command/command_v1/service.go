package command_v1

import (
	"github.com/satanaroom/chat_client/internal/service/client"
)

type ChatClient struct {
	clientService client.Service
}

func NewChatClient(clientService client.Service) *ChatClient {
	return &ChatClient{
		clientService: clientService,
	}
}
