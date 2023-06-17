package repository

import "github.com/yannis94/bank-root/internal/service"

type Storage interface {
    CreateClient(client *service.Client) (*service.Client, error)
    DeleteClient(id int) (*service.Client, error)
    GetClientById(id int) (*service.Client, error)
    GetClientByEmail(email string) (*service.Client, error)
    CreateAccount(account *service.Account) (*service.Account, error)
    GetAccountNumber(client_id int) (string, error) 
    UpdateAccountAmount(uuid string, amount int) error
    GetAccountById(id int) (*service.Account, error)
    CreateClosedAccount(account_num string) error
    CreateTransferDemand(demand *service.TransferDemand) error
    GetTransferDemandById(id int) (*service.TransferDemand, error)
    GetAcceptedTransferDemands() ([]*service.TransferDemand, error)
    UpdateTransferDemand(demand *service.TransferDemand) error
    CreateTransfer(transfer *service.Transfer) error
}
