package command_v1

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/satanaroom/chat_client/pkg/logger"
	"github.com/spf13/cobra"
)

func (c *ChatClient) InitCreateChat() {
	createChat := &cobra.Command{
		Use:   "create",
		Short: "Create chat room.",
		Long:  "Chat room can be created by providing usernames divided by comma.",
		Args:  cobra.ExactArgs(1),
		Run:   c.CreateChat,
	}

	c.root.AddCommand(createChat)
}

func (c *ChatClient) CreateChat(_ *cobra.Command, args []string) {
	if len(args) < 1 {
		out := color.RedString("You should pass usernames to create a chat.\n")
		if _, err := io.WriteString(os.Stdout, out); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	loggedUsername, err := c.clientService.GetLoggedUsername(c.root.Context())
	if err != nil || loggedUsername == "" {
		out := color.RedString("You are not registered in the service.\n")
		if _, err = io.WriteString(os.Stdout, out); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	usernames := strings.Split(args[0], ",")
	if len(usernames) == 0 {
		out := color.RedString("No usernames in command or there aren't valid.\n")
		if _, err = io.WriteString(os.Stdout, out); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	chatId, err := c.clientService.CreateChat(c.root.Context(), loggedUsername, usernames)
	if err != nil {
		out := color.RedString("Chat could not be created.\n")
		if _, err = io.WriteString(os.Stdout, out); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	userMsg := color.GreenString(fmt.Sprintf("Chat %d created successfully\n", chatId))
	if _, err = io.WriteString(os.Stdout, userMsg); err != nil {
		logger.Errorf("failed to write to stdout: %s", err.Error())
	}
}
