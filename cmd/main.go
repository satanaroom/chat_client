package main

import (
	"context"
	"log"

	"github.com/satanaroom/chat_client/internal/app"
	root "github.com/satanaroom/chat_client/internal/command/command_v1"
)

func main() {
	ctx := context.Background()

	if err := app.InitApp(ctx); err != nil {
		log.Fatalf("init app: %s", err.Error())
	}

	root.Execute()
}
