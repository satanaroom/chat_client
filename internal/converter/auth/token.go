package auth

import authV1 "github.com/satanaroom/auth/pkg/auth_v1"

func ToRefreshRequest(username, password string) *authV1.GetRefreshTokenRequest {
	return &authV1.GetRefreshTokenRequest{
		Username: username,
		Password: password,
	}
}

func ToAccessRequest(refreshToken string) *authV1.GetAccessTokenRequest {
	return &authV1.GetAccessTokenRequest{
		RefreshToken: refreshToken,
	}
}

func ToUpdateRefreshRequest(oldToken string) *authV1.UpdateRefreshTokenRequest {
	return &authV1.UpdateRefreshTokenRequest{
		OldToken: oldToken,
	}
}
