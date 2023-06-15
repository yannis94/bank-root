package service

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
}

type Account struct {
    Id int `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Number string `json:"number"`
    Balance int64 `json:"balance"`
    CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstname, lastname string) *Account {
    accountNumber := uuid.New()

    return &Account{
        FirstName: firstname,
        LastName: lastname,
        Number: accountNumber.String(),
        Balance: 0,
        CreatedAt: time.Now().UTC(),
    }
}
