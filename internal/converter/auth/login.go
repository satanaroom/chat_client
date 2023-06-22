package auth

import "github.com/satanaroom/chat_client/internal/model"

func ToLoginService(username, password string) *model.UserInfo {
	return &model.UserInfo{
		Username: username,
		Password: password,
	}
}
