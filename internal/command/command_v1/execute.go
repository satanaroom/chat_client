package command_v1

import "fmt"

func (c *ChatClient) Execute() error {
	if err := c.root.Execute(); err != nil {
		return fmt.Errorf("execute root: %w", err)
	}
	return nil
}
