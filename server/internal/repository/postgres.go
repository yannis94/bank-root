package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/yannis94/bank-root/internal/service"
)

type Postgres struct {
    db *sql.DB
}

func NewPostgres(user, pwd, port, db_name string) (*Postgres, error) {
    connStr := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", user, pwd, port, db_name)

    db, err := sql.Open("postgres", connStr)

    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Postgres{ db: db }, nil
}

func (pg *Postgres) CreateClient(client *service.Client) (*service.Client, error) {
    query := `INSERT INTO client (first_name, last_name, email, password, created_at)
    VALUES ($1, $2, $3, $4, $5) RETURNING id;`

    rows, err := pg.db.Query(query, client.FirstName, client.LastName, client.Email, client.Password, client.CreatedAt)

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        rows.Scan(&client.Id)
    }

    return client, nil
}

func (pg *Postgres) GetClientById(id int) (*service.Client, error) {
    client := &service.Client{}
    query := "SELECT * FROM client WHERE id = $1;"

    rows, err := pg.db.Query(query, id)

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        rows.Scan(&client.Id, &client.FirstName, &client.LastName, &client.Email, &client.Password, &client.CreatedAt)
    }

    return client, nil
}

func (pg *Postgres) GetClientByEmail(email string) (*service.Client, error) {
    query := "SELECT * FROM client WHERE email = $1;"

    rows, err := pg.db.Query(query, email)

    if err != nil {
        return nil, err
    }

    client := &service.Client{}

    for rows.Next() {
        rows.Scan(&client.Id, &client.FirstName, &client.LastName, &client.Email, &client.Password, &client.CreatedAt)
    }

    return client, err
}

func (pg *Postgres) DeleteClient(id int) (*service.Client, error) {
    query := "DELETE FROM client WHERE id = $1;"

    rows, err := pg.db.Query(query, id)

    client := &service.Client{}

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        rows.Scan(&client.Id, &client.FirstName, &client.LastName, &client.Email, &client.Password, &client.CreatedAt)
    }

    return client, nil
}

func (pg *Postgres) CreateAccount(account *service.Account) (*service.Account, error) {
    query := `
        INSERT INTO account (client_id, number, balance, created_at) 
        VALUES ($1, $2, $3, $4) RETURNING id;
    `

    rows, err := pg.db.Query(query, account.ClientId, account.Number, account.Balance, account.CreatedAt)

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        rows.Scan(&account.Id)
    }

    return account, nil
}

func (pg *Postgres) UpdateAccountAmount(uuid string, amount int) error {
    query := "UPDATE account SET amount = $2 WHERE number = $1;"

    _, err := pg.db.Query(query, uuid, amount)

    return err
}

func (pg *Postgres) GetAccountById(id int) (*service.Account, error) {
    var account *service.Account
    query := "SELECT * FROM account WHERE id = $1;"
    rows, err := pg.db.Query(query, id)

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        account, err = accountRowsScanner(rows)

        if err != nil {
            return nil, err
        }
    }
    log.Printf("Account %d : %+v", id, account)

    return account, nil
}

func (pg *Postgres) GetAccountNumber(client_id int) (string, error) {
    query := "SELECT number FROM account WHERE client_id = $1;"

    rows, err := pg.db.Query(query, client_id)

    if err != nil {
        return "", err
    }

    var account_number string

    for rows.Next() {
        rows.Scan(&account_number)
    }

    return account_number, nil
}

func (pg *Postgres) CreateClosedAccount(account_num string) error {
    query := "INSERT INTO closed_account (number, created_at) VALUES ($1, $2);"
    
    _, err := pg.db.Query(query, account_num, time.Now().UTC())

    return err
}

func (pg *Postgres) CreateTransferDemand(demand *service.TransferDemand) error {
    query := `
    INSERT INTO transfer (closed, from_account, to_account, message, amount, accepted, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7);
    `
    _, err := pg.db.Query(query, demand.Closed, demand.FromAccount, demand.ToAccount, demand.Message, demand.Amount, demand.Accepted, demand.CreateAt)

    return err
}

func (pg *Postgres) GetTransferDemandById(id int) (*service.TransferDemand, error) {
    query := "SELECT * FROM transfer WHERE id = $1;"

    rows, err := pg.db.Query(query, id)

    if err != nil {
        return nil, err
    }

    transferDemand := &service.TransferDemand{}

    for rows.Next() {
        rows.Scan(&transferDemand.Id, &transferDemand.Closed, &transferDemand.FromAccount, &transferDemand.ToAccount, &transferDemand.Message, &transferDemand.Amount, &transferDemand.Accepted, &transferDemand.CreateAt)
    }

    return transferDemand, nil
}

func (pg *Postgres) GetAcceptedTransferDemands() ([]*service.TransferDemand, error) {
    query := "SELECT * FROM transfer WHERE closed = $1 AND accepted = $2;"

    rows, err := pg.db.Query(query, false, true)

    if err != nil {
        return nil, err
    }

    var transferDemands []*service.TransferDemand

    for rows.Next() {
        demand := &service.TransferDemand{}
        rows.Scan(&demand.Id, &demand.Closed, &demand.FromAccount, &demand.ToAccount, &demand.Message, &demand.Amount, &demand.Accepted, &demand.CreateAt)

        transferDemands = append(transferDemands, demand)
    }

    return transferDemands, nil
}

func (pg *Postgres) UpdateTransferDemand(demand *service.TransferDemand) error {
    query := "UPDATE transfer SET closed = $1, accepted = $2 WHERE id=$3;"

    _, err := pg.db.Query(query, demand.Closed, demand.Accepted, demand.Id)

    return err
}

func (pg *Postgres) GetAccountTransfer(accountId string) ([]*service.TransferDemand, error) {
    var demands []*service.TransferDemand

    query := "SELECT * FROM transfer WHERE from_account = $1 OR to_account = $1;"

    rows, err := pg.db.Query(query, accountId)

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        demand := &service.TransferDemand{}
        rows.Scan(&demand.Id, &demand.Closed, &demand.FromAccount, &demand.ToAccount, &demand.Message, &demand.Amount, &demand.Accepted, &demand.CreateAt)

        demands = append(demands, demand)
    }

    return demands, nil
}

func (pg *Postgres) CreateSession(session *service.Session) error {
    query := "INSERT INTO session (token_id, refresh_token, expires_at) VALUES ($1, $2, $3);"

    _, err := pg.db.Query(query, session.TokenId, session.Refresh_token, session.ExpiresAt)

    return err
}

func (pg *Postgres) GetSessionFromTokenId(id string) (*service.Session, error) {
    query := "SELECT * FROM session WHERE token_id = $1;"

    rows, err := pg.db.Query(query, id)

    if err != nil {
        return nil, err
    }

    session := &service.Session{}

    for rows.Next() {
        rows.Scan(&session.TokenId, &session.Refresh_token, &session.ExpiresAt)
    }

    return session, nil
}

func (pg *Postgres) DeleteSessionFromTokenId(id string) error {
    query := "DELETE FROM session WHERE token_id = $1;"

    _, err := pg.db.Query(query, id)

    return err
}

func accountRowsScanner(rows *sql.Rows) (*service.Account, error) {
    account := &service.Account{}

    err := rows.Scan(&account.Id, &account.ClientId, &account.Number, &account.Balance, &account.CreatedAt)
    
    if err != nil {
        return nil, err
    }

    return account, nil
}
