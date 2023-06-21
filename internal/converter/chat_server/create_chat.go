package chat_server

import (
	"github.com/satanaroom/chat_client/internal/model"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateChatRequest(usernames []string) *chatV1.CreateChatRequest {
	return &chatV1.CreateChatRequest{
		Usernames: usernames,
	}
}

func ToConnectChatRequest(chatId, username string) *chatV1.ConnectChatRequest {
	return &chatV1.ConnectChatRequest{
		ChatId:   chatId,
		Username: username,
	}
}

func ToSendMessageRequest(chatId string, message *model.Message) *chatV1.SendMessageRequest {
	return &chatV1.SendMessageRequest{
		ChatId: chatId,
		Message: &chatV1.Message{
			Text:   message.Text,
			From:   message.From,
			To:     message.To,
			SentAt: timestamppb.New(message.SentAt),
		},
	}
}
