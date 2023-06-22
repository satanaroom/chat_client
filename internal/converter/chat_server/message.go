package chat_server

import (
	"github.com/satanaroom/chat_client/internal/model"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func ToMessage(message *chatV1.Message) *model.Message {
	return &model.Message{
		Text:   message.GetText(),
		From:   message.GetFrom(),
		To:     message.GetTo(),
		SentAt: message.GetSentAt().AsTime(),
	}
}
