package model

import "time"

const LoggedUsername = "logged_username"

type UserInfo struct {
	Username string
	Password string
}

type TokenInfo struct {
	Username   string
	Token      string
	Expiration time.Duration
}

type UserTokens struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

type Message struct {
	Text   string
	From   string
	To     string
	SentAt time.Time
}
