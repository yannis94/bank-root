package service

import "time"

type Session struct {
    TokenId string
    Refresh_token string
    ExpiresAt int64
}

func NewSession(id, token string) *Session {
    exp := time.Now().Add(60 * time.Minute).Unix()
    return &Session{ TokenId: id, Refresh_token: token, ExpiresAt: exp }
}
