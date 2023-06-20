package queue

import (
	"errors"

	"github.com/yannis94/bank-root/core"
)

type node struct {
    Data *core.TransferDemand
    next *node
}

type Queue struct {
    Length int
    head *node
    tail *node
}

func BuildQueue() *Queue {
    return &Queue{
        Length: 0,
    }
}

func (queue *Queue) Enqueue(data *core.TransferDemand) {
    newNode := &node{ Data: data }

    if queue.Length == 0 {
        queue.head = newNode
    } else {
        queue.tail.next = newNode
    }

    queue.tail = newNode
    queue.Length++
}

func (queue *Queue) Dequeue() (*core.TransferDemand, error) {
    if queue.Length == 0 {
        return nil, errors.New("Queue is empty, cannot dequeue.")
    }

    tmp := queue.head
    queue.head = tmp.next
    tmp.next = nil

    queue.Length-- 

    return tmp.Data, nil
}

func (queue *Queue) Peek() core.TransferDemand {
    return *queue.head.Data
}
