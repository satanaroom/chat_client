package command_v1

import (
	"github.com/satanaroom/auth/pkg/logger"
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitConnectChat() {
	connectChat := &cobra.Command{
		Use:     "login",
		Aliases: []string{"l"},
		Short:   "Inspects a string",
		Args:    cobra.ExactArgs(1),
		Run:     c.ConnectChat,
	}

	c.root.AddCommand(connectChat)
}

func (c *ChatClient) ConnectChat(_ *cobra.Command, _ []string) {
	if err := c.clientService.ConnectChat(c.root.Context()); err != nil {
		logger.Errorf("connect: %s", err.Error())
	}
}
