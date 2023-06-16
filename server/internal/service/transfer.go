package service

import "time"

type Transfer struct {
    Id int `json:"id"`
    DemandId int `json:"demand_id"`
    Done bool `json:"done"`
    CreateAt time.Time `json:"created_at"`
}

type TransferRequest struct {
    DemandId int `json:"demand_id"`
    Done bool `json:"done"`
}

type TransferDemandRequest struct {
    FromAccount string `json:"from_account"`
    ToAccount string `json:"to_account"`
    Amount int `json:"amount"`
    Message string `json:"message"`
}

type TransferDemandUpdateRequest struct {
    TransferDemandId int `json:"transfer_demand_id"`
    Acceped bool `json:"accepted"`
}

type TransferDemand struct {
    Id int `json:"id"`
    Closed bool `json:"closed"`
    FromAccount string `json:"from_account"`
    ToAccount string `json:"to_account"`
    Message string `json:"message"`
    Amount int `json:"amount"`
    Accepted bool `json:"accepted"`
    CreateAt time.Time `json:"created_at"`
}

func NewTransferDemand(amount int, from, to, message string) *TransferDemand {
    return &TransferDemand{
        FromAccount: from,
        ToAccount: to,
        Message: message,
        Amount: amount,
        CreateAt: time.Now().UTC(),
    }
}

func NewTransfer(demand_id int, done bool) *Transfer {
    return &Transfer{
        DemandId: demand_id,
        Done: done,
        CreateAt: time.Now().UTC(),
    }
}
