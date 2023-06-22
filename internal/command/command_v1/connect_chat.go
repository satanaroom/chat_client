package command_v1

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/satanaroom/auth/pkg/logger"
	"github.com/spf13/cobra"
)

func (c *ChatClient) AddConnectChatCmd() {
	connect := &cobra.Command{
		Use:   "connect",
		Short: "Подключение к чату",
		Long:  "Подключение к существующему чату по его айди",
		Run:   c.ConnectChatHandler,
	}

	connect.Flags().StringP("chat-id", "c", "", "Айди существующего чата")
	if err := connect.MarkFlagRequired("chat-id"); err != nil {
		logger.Fatalf("failed to mark chat-id flag required: %s", err.Error())
	}

	root.AddCommand(connect)
}

func (c *ChatClient) ConnectChatHandler(cmd *cobra.Command, _ []string) {
	chatId, err := cmd.Flags().GetString("chat-id")
	if err != nil {
		if _, err = io.WriteString(os.Stdout, color.RedString("Необходимо указать айди существующего чата для подключения к нему.\n")); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	if err = c.clientService.ConnectChat(root.Context(), chatId); err != nil {
		if _, err = io.WriteString(os.Stdout, color.RedString("Возникла ошибка при подключении к чату.\n")); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	if _, err = io.WriteString(os.Stdout, color.GreenString(fmt.Sprintf("Чат [%s] завершён.\n", chatId))); err != nil {
		logger.Errorf("failed to write to stdout: %s", err.Error())
	}
}
