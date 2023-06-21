package command_v1

import (
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitRoot() {
	c.root = &cobra.Command{
		Use:   "chat-client",
		Short: "Клиент для многопользовательского чат-сервера",
	}
}
