package main

import (
	"context"

	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_client/internal/app"
)

func main() {
	ctx := context.Background()

	chatClientApp, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatalf("failed to initialize app: %s", err.Error())
	}

	if err = chatClientApp.Run(); err != nil {
		logger.Fatalf("failed to run app: %s", err.Error())
	}
}
