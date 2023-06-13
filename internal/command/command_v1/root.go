package command_v1

import (
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitRoot() {
	rootCmd := &cobra.Command{
		Use:     "stringer",
		Version: "0.0.1",
		Short:   "stringer - a simple CLI to transform and inspect strings",
		Long: `stringer is a super fancy CLI (kidding)
    
One can use stringer to modify or inspect strings straight from the terminal`,
		Run: c.Root,
	}
	c.root = rootCmd
}

func (c *ChatClient) Root(_ *cobra.Command, _ []string) {
}
