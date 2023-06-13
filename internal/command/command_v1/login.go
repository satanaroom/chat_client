package command_v1

import (
	"io"
	"os"

	converter "github.com/satanaroom/chat_client/internal/converter/auth"
	"github.com/satanaroom/chat_client/pkg/logger"
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitLogin() {
	login := &cobra.Command{
		Use:   "login",
		Short: "Login to chat server via username and password",
		Args:  cobra.ExactArgs(2),
		Run:   c.Login,
	}

	c.root.AddCommand(login)
}

func (c *ChatClient) Login(_ *cobra.Command, args []string) {
	if len(args) < 2 {
		logger.Errorf("no args")
		return
	}
	if err := c.clientService.Login(c.root.Context(), converter.ToLoginService(args[0], args[1])); err != nil {
		logger.Errorf("login: %s", err.Error())
	}

	if _, err := io.WriteString(os.Stdout, "You logged in successfully\n"); err != nil {
		logger.Errorf("failed to write to stdout: %s", err.Error())
	}
}
