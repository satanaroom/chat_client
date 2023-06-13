package command_v1

import (
	"github.com/satanaroom/chat_client/internal/service/client"
	"github.com/spf13/cobra"
)

type ChatClient struct {
	clientService client.Service
	root          *cobra.Command
}

func NewChatClient(clientService client.Service) *ChatClient {
	return &ChatClient{
		clientService: clientService,
	}
}
