package chat_server

import chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"

func ToCreateChatRequest(usernames []string) *chatV1.CreateChatRequest {
	return &chatV1.CreateChatRequest{
		Usernames: usernames,
	}
}
