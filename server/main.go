package main

import (
	"fmt"
	"log"

	"github.com/yannis94/bank-root/internal/api"
	"github.com/yannis94/bank-root/internal/repository"
)

func main() {
    fmt.Println("-----------------------------------")
    fmt.Println("------------ Bank Root ------------")
    fmt.Println("-------------- Server -------------")
    fmt.Println()
    
    repo, err := repository.NewPostgres()
    err = repo.Init()

    if err != nil {
        log.Fatal(err)
    }


    server := api.NewApiServer(":3001", repo)
    server.Start()
}
