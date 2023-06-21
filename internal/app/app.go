package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/satanaroom/chat_client/internal/closer"
	commandV1 "github.com/satanaroom/chat_client/internal/command/command_v1"
	"github.com/satanaroom/chat_client/internal/config"
	"github.com/satanaroom/chat_client/pkg/logger"
)

type App struct {
	serviceProvider *serviceProvider
	chatClient      *commandV1.ChatClient
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, fmt.Errorf("init deps: %w", err)
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.runChatClient(); err != nil {
			logger.Fatalf("failed of usage chat-client: %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		a.initServiceProvider,
		a.initCommands,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return fmt.Errorf("init: %w", err)
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initCommands(ctx context.Context) error {
	a.chatClient = commandV1.NewChatClient(a.serviceProvider.ClientService(ctx))
	a.chatClient.InitRoot()
	a.chatClient.InitLogin()
	a.chatClient.InitCreateChat()
	a.chatClient.InitConnectChat()

	return nil
}

func (a *App) runChatClient() error {
	if err := a.chatClient.Execute(); err != nil {
		return fmt.Errorf("execute: %w", err)
	}

	return nil
}
