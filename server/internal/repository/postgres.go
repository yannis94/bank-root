package repository

import (
	"database/sql"
	"fmt"

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

func (pg *Postgres) CreateAccount(account *service.Account) error {
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
