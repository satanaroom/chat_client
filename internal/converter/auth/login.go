package auth

import "github.com/satanaroom/chat_client/internal/model"

func ToLoginService(username, password string) *model.UserCredentials {
	return &model.UserCredentials{
		Username: username,
		Password: password,
	}
}
