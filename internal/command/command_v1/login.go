package command_v1

import (
	"io"
	"os"

	"github.com/fatih/color"
	converter "github.com/satanaroom/chat_client/internal/converter/auth"
	"github.com/satanaroom/chat_client/pkg/logger"
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitLogin() {
	login := &cobra.Command{
		Use:   "login",
		Short: "Login to chat server via username and password.",
		Args:  cobra.ExactArgs(2),
		Run:   c.Login,
	}

	c.root.AddCommand(login)
}

func (c *ChatClient) Login(_ *cobra.Command, args []string) {
	if len(args) < 2 {
		out := color.RedString("You should pass username and password to login.\n")
		if _, err := io.WriteString(os.Stdout, out); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}
	username := args[0]
	password := args[1]
	if err := c.clientService.Login(c.root.Context(), converter.ToLoginService(username, password)); err != nil {
		out := color.RedString("No such username or password isn't valid.\n")
		if _, err = io.WriteString(os.Stdout, out); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	if err := c.clientService.SetLoggedUsername(c.root.Context(), username); err != nil {
		out := color.RedString("Error while setting logged username.\n")
		if _, err = io.WriteString(os.Stdout, out); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	out := color.GreenString("You logged in successfully\n")
	if _, err := io.WriteString(os.Stdout, out); err != nil {
		logger.Errorf("failed to write to stdout: %s", err.Error())
	}
}
