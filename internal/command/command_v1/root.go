package command_v1

import (
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitRoot() {
	rootCmd := &cobra.Command{
		Use:     "Chat client",
		Version: "0.0.1",
		Short:   "Chat client - a simple CLI to chat in terminal",
		Run:     c.Root,
	}
	c.root = rootCmd
}

func (c *ChatClient) Root(_ *cobra.Command, _ []string) {
}
