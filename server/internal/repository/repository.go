package repository

import "github.com/yannis94/bank-root/internal/service"

type Repository interface {
    CreateAccount(account *service.Account) error
    DeleteAccount(id int) error
    UpdateAccount(account *service.Account) error
    GetAccountById(id int) (*service.Account, error)
}
