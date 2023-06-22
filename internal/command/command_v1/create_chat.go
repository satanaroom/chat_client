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

func (c *ChatClient) AddCreateChatCmd() {
	create := &cobra.Command{
		Use:   "create",
		Short: "Создание комнаты чата",
		Run:   c.CreateChatHandler,
	}

	create.Flags().StringP("usernames", "u", "", "Участники чата, разделённые запятой")
	if err := create.MarkFlagRequired("usernames"); err != nil {
		logger.Fatalf("failed to mark usernames flag required: %s", err.Error())
	}
	root.AddCommand(create)
}

func (c *ChatClient) CreateChatHandler(cmd *cobra.Command, _ []string) {
	usernamesFlag, err := cmd.Flags().GetString("usernames")
	if err != nil {
		if _, err = io.WriteString(os.Stdout, color.RedString("Необходимо указать участников чата для создания комнаты.\n")); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	usernames := strings.Split(usernamesFlag, ",")
	chatId, err := c.clientService.CreateChat(root.Context(), usernames)
	if err != nil {
		if _, err = io.WriteString(os.Stdout, color.RedString("Возникла ошибка при создании чата.\n")); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	if _, err = io.WriteString(os.Stdout, color.GreenString(fmt.Sprintf("Чат [%s] успешно создан.\n", chatId))); err != nil {
		logger.Errorf("failed to write to stdout: %s", err.Error())
	}
}
