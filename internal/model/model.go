package model

import "time"

const (
	LoggedUsername = "logged_username"
	AccessToken    = "access_token"
	RefreshToken   = "refresh_token"
)

type UserInfo struct {
	Username string
	Password string
}

type TokenInfo struct {
	Username   string
	Token      string
	Expiration time.Duration
}

type Message struct {
	Text   string
	From   string
	To     string
	SentAt time.Time
}
