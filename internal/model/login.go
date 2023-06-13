package model

import "time"

type UserCredentials struct {
	Username string
	Password string
}

type TokenInfo struct {
	Username   string
	Token      string
	Expiration time.Duration
}
