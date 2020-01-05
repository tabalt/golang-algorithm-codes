package chapter_2

import (
	"errors"
)

//Loop queue

type LoopQueue struct {
	storage []int

	front int
	rear  int
}

func NewLoopQueue(n int) *LoopQueue {
	return &LoopQueue{
		storage: make([]int, n),
	}
}

func (lq *LoopQueue) Enqueue(val int) error {
	if (lq.rear+1)%len(lq.storage) == lq.front {
		return errors.New("full queue")
	}

	lq.storage[lq.rear] = val
	lq.rear = (lq.rear + 1) % len(lq.storage)

	return nil
}

func (lq *LoopQueue) Dequeue() (int, error) {
	if lq.rear == lq.front {
		return 0, errors.New("empty queue")
	}

	val := lq.storage[lq.front]
	lq.storage[lq.front] = 0
	lq.front = (lq.front + 1) % len(lq.storage)

	return val, nil
}
