package service

type Account struct {
    Id int `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Number int64 `json:"number"`
    Balance int64 `json:"balance"`
}

func NewAccount(firstname, lastname string) *Account {
    return &Account{
        FirstName: firstname,
        LastName: lastname,
    }
}
