package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yannis94/bank-root/internal/config"
)

type Postgres struct {
    db *sql.DB
}

func NewPostgres() (*Postgres, error) {
    connStr := fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", config.DB_USER, config.DB_NAME)
    connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"

}
