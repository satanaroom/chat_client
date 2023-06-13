package command_v1

import (
	"github.com/satanaroom/chat_client/pkg/logger"
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitCreateChat() {
	createChat := &cobra.Command{
		Use:   "create",
		Short: "Create chat room",
		Args:  cobra.ExactArgs(1),
		Run:   c.CreateChat,
	}

	c.root.AddCommand(createChat)
}

func (c *ChatClient) CreateChat(_ *cobra.Command, args []string) {
	if len(args) < 1 {
		logger.Errorf("no args")
		return
	}
	if err := c.clientService.CreateChat(c.root.Context(), args[0]); err != nil {
		logger.Errorf("create chat: %s", err.Error())
	}
}
