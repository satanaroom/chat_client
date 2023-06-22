package main

import (
	"context"

	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_client/internal/app"
	root "github.com/satanaroom/chat_client/internal/command/command_v1"
)

func main() {
	ctx := context.Background()

	if err := app.InitApp(ctx); err != nil {
		logger.Fatalf("init app: %s", err.Error())
	}

	root.Execute()
}
