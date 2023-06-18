package main

import "errors"

type node struct {
    Data *transferDemand
    prev *node
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

func (queue *Queue) Enqueue(data transferDemand) error {
    newNode := &node{ Data: &data }

    queue.Length = queue.Length + 1

    if queue.Length == 0 {
        queue.tail = newNode
        queue.head = newNode
        return nil
    }

    queue.tail.prev = newNode
    queue.tail = newNode

    return nil
}

func (queue *Queue) Dequeue() (*transferDemand, error) {
    if queue.Length == 0 {
        return nil, errors.New("Queue is empty, cannot dequeue.")
    }

    tmp := queue.head
    queue.head = tmp.prev
    tmp.prev = nil

    queue.Length = queue.Length - 1

    return tmp.Data, nil
}

func (queue *Queue) Peek() *transferDemand {
    return queue.head.Data
}
