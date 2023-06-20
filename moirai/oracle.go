package main

import "github.com/yannis94/bank-root/queue"

type oracle struct {
    transferQueue *queue.Queue
}

func (oracle *oracle) GetWaitingTransfer() {
}

func (oracle *oracle) transferValidation() {
}
