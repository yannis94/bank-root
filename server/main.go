package main

import (
	"fmt"

	"github.com/yannis94/bank-root/internal/api"
)

func main() {
    fmt.Println("-----------------------------------")
    fmt.Println("------------ Bank Root ------------")
    fmt.Println("-------------- Server -------------")
    fmt.Println("-----------------------------------")
    fmt.Println()

    server := api.NewApiServer(":3001")
    server.Start()
}
