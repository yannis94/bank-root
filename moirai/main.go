package main

import (
	"fmt"
	"log"
)

func main() {
    fmt.Println("MOIRAI")
    oracle := StartOracle()

    oracle.api.GetTransferDemand()

    for oracle.transferQueue.Length > 0 {
        transfer, err := oracle.transferQueue.Dequeue()
        
        if err != nil {
            log.Fatal(err)
        }

        oracle.transferValidation(*transfer)
    }
}
