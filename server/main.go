package main

import (
	"fmt"
	"log"

	"github.com/yannis94/bank-root/internal/api"
	"github.com/yannis94/bank-root/internal/auth"
	"github.com/yannis94/bank-root/internal/config"
	"github.com/yannis94/bank-root/internal/repository"
)

func main() {
    fmt.Println("------------ Bank Root ------------")
    fmt.Println("-------------- Server -------------")
    fmt.Println()
    
    repo, err := repository.NewPostgres(config.DB_USER, config.DB_PASS, config.DB_PORT, config.DB_NAME)

    if err != nil {
        log.Fatal(err)
    }

    authService := auth.NewAuthService(config.JWT_SECRET)

    server := api.NewApiServer(":3001", repo, authService)
    server.Start()
}
