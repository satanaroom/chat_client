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
		Short: "Авторизация на сервере чата",
		Long:  "Авторизация на сервере чата при помощи логина и пароля",
		Run:   c.Login,
	}

	login.Flags().StringP("username", "u", "", "Имя пользователя")
	if err := login.MarkFlagRequired("username"); err != nil {
		logger.Fatalf("failed to mark username flag as required: %s", err.Error())
	}

	login.Flags().StringP("password", "p", "", "Пароль пользователя")
	if err := login.MarkFlagRequired("password"); err != nil {
		logger.Fatalf("failed to mark password flag as required: %s", err.Error())
	}

	c.root.AddCommand(login)
}

func (c *ChatClient) Login(cmd *cobra.Command, args []string) {
	username, err := cmd.Flags().GetString("username")
	if err != nil {
		if _, err = io.WriteString(os.Stdout, color.RedString("Необходимо указать логин для авторизации на сервере.\n")); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		if _, err = io.WriteString(os.Stdout, color.RedString("Необходимо указать пароль для авторизации на сервере.\n")); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	if err = c.clientService.Login(c.root.Context(), converter.ToLoginService(username, password)); err != nil {
		if _, err = io.WriteString(os.Stdout, color.RedString("Пользователь не найден.\n")); err != nil {
			logger.Errorf("failed to write to stdout: %s", err.Error())
		}
		return
	}

	if _, err = io.WriteString(os.Stdout, color.GreenString("Вы успешно авторизировались.\n")); err != nil {
		logger.Errorf("failed to write to stdout: %s", err.Error())
	}
}
