package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/satanaroom/auth/pkg/logger"
	converter "github.com/satanaroom/chat_client/internal/converter/chat_server"
	"github.com/satanaroom/chat_client/internal/model"
)

func (s *service) ConnectChat(ctx context.Context, chatId string) error {
	loggedUsername, err := s.redisClient.Get(ctx, model.LoggedUsername)
	if err != nil {
		return fmt.Errorf("redis.Client: %w", err)
	}

	s.RefreshTokens(ctx, s.authConfig.RefreshTokenExpiration(), s.authConfig.AccessTokenExpiration())

	logger.Infof("username %s connected to chat [%s]", loggedUsername, chatId)

	stream, err := s.chatClient.ConnectChat(ctx, chatId, loggedUsername)
	if err != nil {
		return fmt.Errorf("chatClient.ConnectChat: %w", err)
	}

	go func() {
		for {
			resp, errRecv := stream.Recv()
			if errRecv == io.EOF {
				logger.Info("stream closed")
				return
			}
			if errRecv != nil {
				logger.Errorf("failed to receive message from stream: %s", errRecv.Error())
				return
			}

			recievedMessage := converter.ToMessage(resp.GetMessage())
			log.Print(color.BlueString("[%v] - [from: %s]: %s\n", recievedMessage.SentAt, recievedMessage.From, recievedMessage.Text))
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		var lines strings.Builder

		for {
			scanner.Scan()
			line := scanner.Text()
			if len(line) == 0 {
				break
			}

			lines.WriteString(line)
			lines.WriteString("\n")
		}

		if err = scanner.Err(); err != nil {
			logger.Errorf("failed to scan message: %s", err.Error())
		}
		if err = s.chatClient.SendMessage(ctx, chatId, &model.Message{
			From:   loggedUsername,
			To:     "all users",
			Text:   lines.String(),
			SentAt: time.Now(),
		}); err != nil {
			logger.Errorf("failed to send message: %s", err.Error())
			return err
		}
	}
}
