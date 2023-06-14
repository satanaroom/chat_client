package model

import "time"

const RegisteredUsername = "registered_username"

type UserCredentials struct {
	Username string
	Password string
}

type TokenInfo struct {
	Username   string
	Token      string
	Expiration time.Duration
}
