package main

import (
	"log"

	"github.com/yannis94/bank-root/core"
	"github.com/yannis94/bank-root/queue"
)

type oracle struct {
    api *core.BankAPI
    transferQueue *queue.Queue
    errorQueue *queue.Queue
}

func StartOracle() *oracle {
    return &oracle{
        api: core.NewBankAPI(),
        transferQueue: queue.BuildQueue(),
        errorQueue: queue.BuildQueue(),
    }
}

func (oracle *oracle) getWaitingTransfer() {
    demands, err := oracle.api.GetTransferDemand()

    if err != nil {
        log.Fatalf("Cannot get transfer demands. Error: %s", err.Error())
    }

    for _, demand := range demands {
        oracle.transferQueue.Enqueue(&demand)
    }
}

func (oracle *oracle) isAccountBalanceEnough(amount, balance int) bool {
    return balance - amount >= 0
}

func (oracle *oracle) transferValidation(transfer *core.TransferDemand) error {
    emitterAccountNumber := transfer.FromAccount
    receiverAccountNumber := transfer.ToAccount

    emitter, err := oracle.api.GetAccount(emitterAccountNumber)

    if err != nil {
        return err
    }

    if !oracle.isAccountBalanceEnough(transfer.Amount, int(emitter.Balance)) {
        transfer.Accepted = false
        return oracle.api.SendTransferValidation(*transfer)
    }

    receiver, err := oracle.api.GetAccount(receiverAccountNumber)

    receiver.Balance += int64(transfer.Amount)

    err = oracle.api.UpdateAccount(*receiver)

    if err != nil {
        oracle.errorQueue.Enqueue(transfer)
        return err
    }

    emitter.Balance -= int64(transfer.Amount)
    err = oracle.api.UpdateAccount(*emitter)

    if err != nil {
        oracle.errorQueue.Enqueue(transfer)
        return err
    }
    
    transfer.Accepted = true
    return oracle.api.SendTransferValidation(*transfer)
}
