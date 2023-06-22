package app

import (
	"context"
	"fmt"

	"github.com/satanaroom/chat_client/internal/closer"
	commandV1 "github.com/satanaroom/chat_client/internal/command/command_v1"
	"github.com/satanaroom/chat_client/internal/config"
)

func InitApp(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	if err := config.Init(ctx); err != nil {
		return fmt.Errorf("init config: %w", err)
	}

	chatClient := commandV1.NewChatClient(newServiceProvider().ClientService(ctx))
	chatClient.AddLoginCmd()
	chatClient.AddCreateChatCmd()
	chatClient.AddConnectChatCmd()

	return nil
}
