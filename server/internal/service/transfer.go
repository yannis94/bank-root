package service

type TransferRequest struct {
    FromAccount string `json:"from_account"`
    ToAccount string `json:"to_account"`
    Amount int `json:"amount"`
}
