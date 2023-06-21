package core

import "time"

type TransferDemand struct {
    Id int `json:"id"`
    Closed  bool `json:"closed"`
    FromAccount string `json:"from_account"`
    ToAccount string `json:"to_account"`
    Amount int `json:"amount"`
    Accepted bool `json:"accepted"`
}

type Account struct {
    Id int `json:"id"`
    ClientId int `json:"client_id"`
    Number string `json:"number"`
    Balance int64 `json:"balance"`
    CreatedAt time.Time `json:"created_at"`
}
