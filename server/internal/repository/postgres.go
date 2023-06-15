package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/yannis94/bank-root/internal/config"
	"github.com/yannis94/bank-root/internal/service"
)

type Postgres struct {
    db *sql.DB
}

func NewPostgres() (*Postgres, error) {
    connStr := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", config.DB_USER, config.DB_PASS, config.DB_PORT, config.DB_NAME)

    db, err := sql.Open("postgres", connStr)

    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Postgres{ db: db }, nil
}

func (pg *Postgres) Init() error {
    return pg.createAccoutTable()
}

func (pg *Postgres) createAccoutTable() error {
    query := `CREATE TABLE IF NOT EXISTS account (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(50),
        last_name VARCHAR(120),
        number UUID,
        balance INTEGER, 
        created_at TIMESTAMP
    );
    `
    _, err := pg.db.Exec(query)

    return err
}

func (pg *Postgres) CreateAccount(account *service.Account) error {
    query := `
        INSERT INTO account (first_name, last_name, number, balance, created_at) 
        VALUES ($1, $2, $3, $4, $5);
    `

    resp, err := pg.db.Query(query, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt)

    if err != nil {
        return err
    }

    log.Println(resp)

    return nil
}

func (pg *Postgres) DeleteAccount(id int) error {
    return nil
}

func (pg *Postgres) UpdateAccount(account *service.Account) error {
    return nil
}

func (pg *Postgres) GetAccountById(id int) (*service.Account, error) {
    return nil, nil
}
