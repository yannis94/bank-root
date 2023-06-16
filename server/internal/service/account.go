package service

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
    ClientId int `json:"client_id"`
    Deposit int `json:"deposit"`
}

type Account struct {
    Id int `json:"id"`
    ClientId int `json:"client_id"`
    Number string `json:"number"`
    Balance int64 `json:"balance"`
    CreatedAt time.Time `json:"created_at"`
}

func NewAccount(client_id int, deposit int) *Account {
    accountNumber := uuid.New()

    return &Account{
        ClientId: client_id,
        Number: accountNumber.String(),
        Balance: int64(deposit),
        CreatedAt: time.Now().UTC(),
    }
}
